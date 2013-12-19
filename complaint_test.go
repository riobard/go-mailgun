package mailgun

import (
	"reflect"
	"testing"
	"time"
)

func TestComplaint(t *testing.T) {
	n, res, err := c.Complaints(*domain, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total complaints: %d", n)
	for _, r := range res {
		t.Logf("%+v", r)
	}
}

func TestComplaintTime(t *testing.T) {
	c := Complaint{
		Count:     2,
		CreatedAt: "Tue, 15 Nov 2011 08:25:11 GMT",
		Address:   "baz@example.com",
	}
	if reflect.TypeOf(c.Time()) != reflect.TypeOf(time.Now()) {
		t.Fatalf("Should retrun a Time %+v", c.Time())
	}
}
