package slack

import (
	"testing"
)

func TestNewWriter(t *testing.T) {
	samples := []struct {
		channel string
		user    string
		token   string
		err     bool
	}{
		{"", "", "", true},
		{"", "user", "", true},
		{"", "", "token", true},
		{"", "user", "token", true},
		{"channel", "", "", true},
		{"channel", "user", "", true},
		{"channel", "", "token", true},
		{"channel", "user", "token", false},
	}

	for _, s := range samples {
		_, err := NewWriter(s.channel, s.user, s.token)
		if (err != nil) != s.err {
			t.Errorf("NewWriter(channel=%s, user=%s, token=%s): got err!=nil == %t",
				s.channel, s.user, s.token, err != nil)
		}
	}
}

func TestWriterFields(t *testing.T) {
	samples := []struct {
		channel string
		user    string
		token   string
	}{
		{"#channelname", "username", "token"},
	}

	for _, s := range samples {
		w, err := NewWriter(s.channel, s.user, s.token)
		if err != nil {
			t.Errorf("NewWriter failed c=%s u=%s t=%s",
				s.channel, s.user, s.token)
		}
		if c := w.Channel(); c != s.channel {
			t.Errorf(".Channel(): expected %s got %s", s.channel, c)
		}
		if u := w.Username(); u != s.user {
			t.Errorf(".Username(): expected %s got %s", s.user, u)
		}
		if a := w.Token(); a != s.token {
			t.Errorf(".Token(): expected %s got %s", s.token, a)
		}

	}
}
