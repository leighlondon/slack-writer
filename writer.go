package slack

import (
	"net/http"
	"net/url"
)

const api = "https://slack.com/api/chat.postMessage"

type Writer struct {
	Channel string
	User    string
	Token   string
}

func NewWriter(channel, user, token string) *Writer {
	w := Writer{
		Channel: channel,
		User:    user,
		Token:   token,
	}
	return &w
}

func (w *Writer) Write(p []byte) (int, error) {
	d := url.Values{}
	d.Add("channel", w.Channel)
	d.Add("username", w.User)
	d.Add("token", w.Token)
	d.Add("text", string(p))
	_, _ = http.PostForm(api, d)
	// The errors and return code don't line up but are just
	// so that this can be called as a Writer.
	return 0, nil
}
