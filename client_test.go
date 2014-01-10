package mailgun

import (
	"flag"
	"testing"
)

var (
	key    = flag.String("key", "<your key>", "Mailgun key")
	domain = flag.String("domain", "<your domain>", "Test domain")
	from   = flag.String("from", "postmaster@mailgun.org", "Test mail sender address")
	to     = flag.String("to", "postmaster@mailgun.org", "Test mail recipient address")
	c      *Client
)

func init() {
	flag.Parse()
	c = New(*key)
}

func TestWrongResponse(t *testing.T) {

	m := &mail{
		from:      *from,
		to:        []string{*to},
		cc:        []string{"postmaster@mailgun.org"},
		bcc:       []string{"postmaster@mailgun.org"},
		options:   map[string]string{"tag": "awesome"},
		headers:   map[string]string{"julien": "is_awesome"},
		variables: map[string]string{"name": "julien", "twitter": "juliendsv"},
		subject:   "test mail",
		text:      "this is a test mail",
	}

	cw := New("")

	id, err := cw.Send(m)
	if err == nil {
		t.Fatal(id)
	}
	t.Logf("MessageID = %s", id)
}

func TestWrongFrom(t *testing.T) {

	m := &mail{
		from:    "gn",
		to:      []string{*to},
		subject: "test mail",
		text:    "this is a test mail",
	}

	id, err := c.Send(m)
	if err == nil {
		t.Fatal("We should have an error here %v", id)
	}
	t.Logf("MessageID = %s", id)
}
