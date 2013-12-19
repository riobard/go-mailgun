package mailgun

import (
	"reflect"
	"testing"
	"time"
)

// Zero/uninitialized date is 1-1-1
func UTCDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func TestStat(t *testing.T) {
	n, res, err := c.Stats(*domain, 10, 0, nil, UTCDate(1, 1, 1))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total stats: %d", n)
	for _, r := range res {
		t.Logf("%+v", r)
	}
}

func TestStatTime(t *testing.T) {
	s := Stat{
		Count:     2,
		CreatedAt: "Tue, 15 Nov 2011 08:25:11 GMT",
		Tags:      map[string]int{"nice": 1},
		Event:     "sent",
	}
	if reflect.TypeOf(s.Time()) != reflect.TypeOf(time.Now()) {
		t.Fatalf("Should retrun a Time %+v", s.Time())
	}
}
