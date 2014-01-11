package mailgun

import (
	"reflect"
	"testing"
	"time"
)

func TestNewWebhook(t *testing.T) {
	wh := NewWebhook("testWh")
	if reflect.TypeOf(wh) != reflect.TypeOf(&Webhook{}) {
		t.Errorf("We should have a WebHook type, instead we have %v", reflect.TypeOf(wh))
	}
}

func TestEvent(t *testing.T) {
	now := time.Now()
	event := Event{
		name:   "unsubscribe",
		time:   now,
		rcpt:   "postmaster@mbm.mailgun.org",
		msgid:  "82938129382083",
		reason: "leave me alone.",
	}
	if event.Name() != "unsubscribe" {
		t.Errorf("Wrong event name: %v", event.Name())
	}
	if event.Time() != now {
		t.Errorf("Wrong event time: %v", event.Time())
	}
	if event.Rcpt() != "postmaster@mbm.mailgun.org" {
		t.Errorf("Wrong event recipient: %v", event.Rcpt())
	}
	if event.MsgId() != "82938129382083" {
		t.Errorf("Wrong msg id: %v", event.MsgId())
	}
	if event.Reason() != "leave me alone." {
		t.Errorf("Wrong reason: %v", event.Reason())
	}
}
