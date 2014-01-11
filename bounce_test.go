package mailgun

import (
	"reflect"
	"testing"
	"time"
)

func TestBounce(t *testing.T) {
	n, res, err := c.Bounces(*domain, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total bounces: %d", n)
	for _, r := range res {
		t.Logf("%+v", r)
	}
}

func TestBounceTime(t *testing.T) {
	bounce := Bounce{
		Code:      550,
		CreatedAt: "Fri, 21 Oct 2011 11:02:55 GMT",
		Error:     "Message was not accepted -- invalid mailbox.  Local mailbox 'baz@example.com is unavailable: user not found",
		Address:   "baz@example.com",
	}
	if reflect.TypeOf(bounce.Time()) != reflect.TypeOf(time.Now()) {
		t.Fatalf("Should retrun a Time %+v", bounce.Time())
	}
}
