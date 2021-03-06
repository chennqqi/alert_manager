package stats

import (
	"context"
	"github.com/golang/glog"
	"github.com/mayuresh82/alert_manager/internal/reporting"
	"runtime"
	"sync"
	"time"
)

type statType int

const measurement = "alert_manager_stats"

type Stat interface {
	Add(Value int64)
	Reset()
	Set(value int64)
}

// counter is always incremented
type Counter struct {
	name  string
	value int64

	sync.Mutex
}

// gauge is any arbitrary value
type Gauge struct {
	name    string
	values  map[time.Time]int64
	lastVal int64

	sync.Mutex
}

var (
	allCounters []*Counter
	allGauges   []*Gauge
)

func NewCounter(name string) *Counter {
	s := &Counter{name: name}
	allCounters = append(allCounters, s)
	return s
}

func NewGauge(name string) *Gauge {
	s := &Gauge{name: name, values: make(map[time.Time]int64)}
	allGauges = append(allGauges, s)
	return s
}

func (c *Counter) Add(value int64) {
	c.Lock()
	defer c.Unlock()
	if value < 0 {
		glog.Errorf("Cannot subtract from counter type")
		return
	}
	c.value += value
}

func (c *Counter) Reset() {
	c.Lock()
	defer c.Unlock()
	c.value = 0
}

func (c *Counter) Set(value int64) {
	glog.Errorf("CAnnot set a counter type")
}

func (c *Counter) toDatapoint() *reporting.Datapoint {
	c.Lock()
	defer c.Unlock()
	return &reporting.Datapoint{
		Measurement: measurement,
		Fields:      map[string]interface{}{c.name: c.value},
		TimeStamp:   time.Now(),
	}
}

func (g *Gauge) Set(value int64) {
	g.Lock()
	defer g.Unlock()
	g.lastVal = value
	g.values[time.Now()] = g.lastVal
}

func (g *Gauge) Add(value int64) {
	g.Lock()
	defer g.Unlock()

	// add to the last datapoint in the series
	newVal := g.lastVal + value
	g.values[time.Now()] = newVal
	g.lastVal = newVal
}

func (g *Gauge) Reset() {
	g.Lock()
	defer g.Unlock()
	g.values = make(map[time.Time]int64)
}

func (g *Gauge) toDatapoint() []*reporting.Datapoint {
	g.Lock()
	defer g.Unlock()
	dp := []*reporting.Datapoint{}
	for ts, v := range g.values {
		dp = append(dp, &reporting.Datapoint{
			Measurement: measurement,
			Fields:      map[string]interface{}{g.name: v},
			TimeStamp:   ts,
		})
	}
	if len(dp) == 0 {
		dp = append(dp, &reporting.Datapoint{
			Measurement: measurement,
			Fields:      map[string]interface{}{g.name: g.lastVal},
			TimeStamp:   time.Now(),
		})
	}
	return dp
}

func StartExport(ctx context.Context, interval time.Duration) {
	if interval == 0 {
		interval = 60 * time.Second
	}
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C:
			for _, c := range allCounters {
				reporting.DataChan <- c.toDatapoint()
			}
			for _, g := range allGauges {
				for _, dp := range g.toDatapoint() {
					reporting.DataChan <- dp
				}
				g.Reset()
			}
			for _, dp := range internalStats() {
				reporting.DataChan <- dp
			}
		case <-ctx.Done():
			return
		}
	}
}

func internalStats() []*reporting.Datapoint {
	return []*reporting.Datapoint{
		&reporting.Datapoint{
			Measurement: measurement,
			Fields:      map[string]interface{}{"internal.num_goroutines": runtime.NumGoroutine()},
			TimeStamp:   time.Now(),
		},
	}
}
