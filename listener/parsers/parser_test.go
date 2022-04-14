package parsers

import (
	"testing"

	"github.com/mayuresh82/alert_manager/listener"
	"github.com/stretchr/testify/assert"
)

var testDatas = map[string]struct {
	raw    string
	parsed *listener.WebHookAlertData
}{
	"grafana": {
		raw: `{
      "title": "My First Test Alert",
      "ruleId": 1,
      "ruleName": "Neteng BB Input Errors Test",
      "ruleUrl": "http://url.to.grafana/db/dashboard/my_dashboard?panelId=2",
      "state": "alerting",
      "imageUrl": "http://s3.image.url",
      "message": "A BB Link is experiencing > 1000pps input errors over 15 minutes!",
      "evalMatches": [
        {
          "metric": "input_errors",
          "tags": {"device": "br1-sjc1", "interface": "et-0/0/3:0"},
          "value": 1222
        }
      ]
    }`,
		parsed: &listener.WebHookAlertData{
			Id:      "1",
			Name:    "Neteng BB Input Errors Test",
			Details: "A BB Link is experiencing > 1000pps input errors over 15 minutes!\nMetric: input_errors, Value: 1222\n",
			Device:  "br1-sjc1",
			Entity:  "et-0/0/3:0",
			Status:  listener.Status_ALERTING,
			Source:  "grafana",
		},
	},
	"observium": {
		raw: `{"ALERT_STATE":"ALERTING","ALERT_URL":"http:blah","ALERT_UNIXTIME":1535139630,"ALERT_TIMESTAMP":"2018-08-24 19:40:30 +00:00","ALERT_TIMESTAMP_RFC2822":"Fri, 24 Aug 2018 19:40:30 +0000","ALERT_TIMESTAMP_RFC3339":"2018-08-24T19:40:30+00:00","ALERT_ID":"106678","ALERT_MESSAGE":"Neteng BGP Down","CONDITIONS":"bgpPeerState notequals established (idle)","METRICS":"bgpPeerState = idle","DURATION":"2m 11s (2018-08-24 19:38:19)","ENTITY_LINK":"<a href=\"http:\/\/blah.local\/device\/device=529\/tab=routing\/proto=bgp\/\" class=\"entity-popup \" data-eid=\"4278\" data-etype=\"bgp_peer\">AS65101 10.130.249.121<\/a>","ENTITY_NAME":"AS65101 10.130.249.121","ENTITY_ID":"4278","ENTITY_TYPE":"bgp_peer","ENTITY_DESCRIPTION":"","DEVICE_HOSTNAME":"bs1-ash1","DEVICE_SYSNAME":"bs1-ash1","DEVICE_ID":"627","DEVICE_LINK":"<a href=\"http:\/\/rp-lutil9.roblox.local\/device\/device=627\/\" class=\"entity-popup \" data-eid=\"627\" data-etype=\"device\">bs1-ash1<\/a>","DEVICE_HARDWARE":"QFX10002-72Q","DEVICE_OS":"Juniper JunOS 17.3R2-S2.1 (Ethernet Switch)","DEVICE_LOCATION":"44274 Round Table Plaza, Building L Suite 112, Ashburn, VA 20147","DEVICE_UPTIME":"56 days, 22h 26m 23s","DEVICE_REBOOTED":"2018-06-28 21:13:44","TITLE":"ALERT: [bs1-ash1] [bgp_peer] [AS65101 10.130.249.121] Neteng BGP Down"}`,
		parsed: &listener.WebHookAlertData{
			Id:      "106678",
			Name:    "Neteng BGP Down",
			Details: "ALERT: [bs1-ash1] [bgp_peer] [AS65101 10.130.249.121] Neteng BGP Down Url: http:blah",
			Device:  "bs1-ash1",
			Entity:  "AS65101 10.130.249.121",
			Status:  listener.Status_ALERTING,
			Source:  "observium",
		},
	},
	"kapacitor": {
		raw: `{"id":"Neteng Transit Util Out","message":"br2-lhr1:et-0/0/9:1","details":"Transit Util exceeds 0.4265067984","time":"2018-11-08T00:20:00Z","duration":0,"level":"WARNING","data":{"series":[{"name":"jnpr_interface_stat","tags":{"device":"br2-lhr1","entity":"et-0/0/9:1"},"columns":["time","sigma","stat"],"values":[["2018-11-08T00:20:00Z",0,0.4265067984]]}]},"previousLevel":"WARNING"}`,
		parsed: &listener.WebHookAlertData{
			Id:      "",
			Name:    "Neteng Transit Util Out",
			Details: "br2-lhr1:et-0/0/9:1\nTransit Util exceeds 0.4265067984",
			Device:  "br2-lhr1",
			Entity:  "et-0/0/9:1",
			Status:  listener.Status_ALERTING,
			Source:  "kapacitor",
		},
	},
	"ns1": {
		raw: `{"job": {"notify_delay": 0,"job_type": "tcp","last_log_removal": 1548139477,"frequency": 20, "rapid_recheck": false, "region_scope": "fixed","id": "5c46bbc1a632f60001fwerwer", "notify_repeat": 0,"notify_regional": false,"regions": ["ams","dal","lga","sin","sjc"],"policy": "all","config": {"host": "1.6.7.35","port": 443},"status": {},"notify_failback": true,"rules": [],"v2": true,"active": true,"name": "monitornames","notes": null},"region": "global","since": 1548139477,"state": "down"}`,
		parsed: &listener.WebHookAlertData{
			Id:      "5c46bbc1a632f60001fwerwer",
			Name:    "Neteng DNS Monitor Down",
			Details: "monitornames",
			Entity:  "monitornames",
			Status:  listener.Status_ALERTING,
			Source:  "ns1",
		},
	},
}

func TestParsing(t *testing.T) {

	for name, data := range testDatas {
		var parser listener.Parser
		switch name {
		case "grafana":
			parser = &GrafanaParser{name: "grafana"}
		case "observium":
			parser = &ObserviumParser{name: "observium"}
		case "kapacitor":
			parser = &KapacitorParser{name: "kapacitor"}
		case "ns1":
			parser = &Ns1Parser{name: "ns1"}
		}
		result, err := parser.Parse([]byte(data.raw))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, result.Id, data.parsed.Id)
		assert.Equal(t, result.Name, data.parsed.Name)
		assert.Equal(t, result.Details, data.parsed.Details)
		assert.Equal(t, result.Device, data.parsed.Device)
		assert.Equal(t, result.Entity, data.parsed.Entity)
		assert.Equal(t, result.Status, data.parsed.Status)
		assert.Equal(t, result.Source, data.parsed.Source)
	}
}

func TestParsingGeneric(t *testing.T) {
	raw := `{"id": 1, "name": "Generic JSON alert", "entity": "ent1", "device": "dev1", "description": "its down", "timestamp": "2018-12-04T14:57:34-06:00", "severity": "info", "status": "alerting"}`
	parser := &GenericParser{name: "generic"}
	result, err := parser.Parse([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result.Id, "1")
	assert.Equal(t, result.Name, "Generic JSON alert")
	assert.Equal(t, result.Entity, "ent1")
	assert.Equal(t, result.Device, "dev1")
	assert.Equal(t, result.Details, "its down")
	assert.Equal(t, result.Status, listener.Status_ALERTING)
	assert.Equal(t, result.Level, "INFO")

	raw = `{"id": 1, "name": "Generic JSON alert", "entity": "ent1"}`
	result, err = parser.Parse([]byte(raw))
	assert.Error(t, err)
}
