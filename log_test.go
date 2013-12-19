package mailgun

import (
	"reflect"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	n, res, err := c.Logs(*domain, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total logs: %d", n)
	for _, r := range res {
		t.Logf("%+v", r)

		created_at := r.Time()
		if reflect.TypeOf(created_at) != reflect.TypeOf(time.Now()) {
			t.Fatalf("Should retrun a Time %+v", r.Time())
		}
	}
}
