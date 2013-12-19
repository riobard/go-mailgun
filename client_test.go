package mailgun

import (
	"flag"
	"testing"
)

/*
	key    = flag.String("key", "", "key-3kp1667q2vkpubc51c30ifhzv3ypai08")
	domain = flag.String("domain", "", "mbm.mailgun.org")
	from   = flag.String("from", "", "julien.dasilva+mailgunsender@gmail.com")
	to     = flag.String("to", "", "julien.dasilva+mailgunreceiver@gmail.com")
*/

var (
	key    = flag.String("key", "key-3kp1667q2vkpubc51c30ifhzv3ypai08", "Mailgun key")
	domain = flag.String("domain", "mbm.mailgun.org", "Test domain")
	from   = flag.String("from", "postmaster@mbm.mailgun.org", "Test mail sender address")
	to     = flag.String("to", "postmaster@mbm.mailgun.org", "Test mail recipient address")
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
		cc:        []string{"postmaster@mbm.mailgun.org"},
		bcc:       []string{"postmaster@mbm.mailgun.org"},
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
