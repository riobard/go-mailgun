package mailgun

import (
	"net/url"
	"testing"
)

func TestSend(t *testing.T) {
	email := &Email{
		From:    *from,
		To:      []string{*to},
		Subject: "test mail",
		Text:    "this is a test mail",
	}
	email.Headers = url.Values{}

	id, err := mg.Send(*domain, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MessageID: %s", id)
}
