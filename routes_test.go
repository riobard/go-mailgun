package mailgun

import (
	"testing"
)

func TestRoute(t *testing.T) {
	n, res, err := c.Routes(10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total routes: %d", n)
	for _, r := range res {
		t.Logf("%+v", r)
	}

	r := &Route{
		Priority:    0,
		Description: "test",
		Expression:  "match_recipient('.*@gmail.com')",
		Actions:     []string{"forward(\"alex@mailgun.net\")"},
	}

	r.Id, err = c.Create(r)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Delete(r)
	if err != nil {
		t.Fatal(err)
	}
}
