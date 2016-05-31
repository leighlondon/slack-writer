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
