package models

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net"
	"time"
)

// custom structs to allow for mocking
type Dbase interface {
	NewTx() Txn
	Close() error
}

type DB struct {
	*sqlx.DB
}

func (d *DB) NewTx() Txn {
	tx := d.DB.MustBegin()
	return &Tx{tx}
}

func NewDB(addr, username, password, dbName, schemaFile string, timeout int) Dbase {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		glog.Fatalf("Invalid DB addr: %s", addr)
	}
	if host == "" {
		host = "localhost"
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable", host, port, username, password, dbName, timeout)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		glog.Fatalf("Cant open DB: %v", err)
	}
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		glog.Fatalf("Unable to read schema file")
	}
	db.MustExec(string(schema))
	return &DB{db}
}

type Txn interface {
	InQuery(query string, arg ...interface{}) error
	InSelect(query string, to interface{}, arg ...interface{}) error
	UpdateAlert(alert *Alert) error
	NewAlert(alert *Alert) (int64, error)
	GetAlert(query string, args ...interface{}) (*Alert, error)
	SelectAlerts(query string) (Alerts, error)
	SelectRules(query string) (SuppRules, error)
	Rollback() error
	Commit() error
}

type Tx struct {
	*sqlx.Tx
}

func (tx *Tx) InQuery(query string, arg ...interface{}) error {
	query, args, err := sqlx.In(query, arg...)
	if err != nil {
		return err
	}
	query = tx.Rebind(query)
	_, err = tx.Exec(query, args...)
	return err
}

func (tx *Tx) InSelect(query string, to interface{}, arg ...interface{}) error {
	query, args, err := sqlx.In(query, arg...)
	if err != nil {
		return err
	}
	query = tx.Rebind(query)
	return tx.Select(to, query, args...)
}

// WithTx wraps a transaction around a function call.
func WithTx(ctx context.Context, tx Txn, cb func(ctx context.Context, tx Txn) error) error {
	err := cb(ctx, tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

type MyTime struct {
	time.Time
}

func (t MyTime) Value() (driver.Value, error) {
	return driver.Value(t.Unix()), nil
}

func (t *MyTime) Scan(src interface{}) error {
	ns := sql.NullInt64{}
	if err := ns.Scan(src); err != nil {
		return err
	}

	if !ns.Valid {
		return fmt.Errorf("MyTime.Scan: column is not nullable")
	}
	*t = MyTime{time.Unix(ns.Int64, 0)}
	return nil
}
