// Package slack provides a writer that implements the io.Writer interface,
// meaning that it can be used to dump messages to a Slack channel directly
// from Go code.
package slack

import (
	"net/http"
	"net/url"
)

// The connection URL for the API resource.
const api = "https://slack.com/api/chat.postMessage"

// Writer implements the io.Writer interface and stores basic state.
type Writer struct {
	Channel string
	User    string
	Token   string
}

// NewWriter returns a new writer configured for use with the Slack API.
func NewWriter(channel, user, token string) *Writer {
	w := Writer{
		Channel: channel,
		User:    user,
		Token:   token,
	}
	return &w
}

// Write writes the input byte data to the nominated channel as a string.
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
