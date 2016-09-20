// Package slack provides a writer that implements the io.Writer interface,
// meaning that it can be used to dump messages to a Slack channel directly
// from Go code.
package slack

import (
	"errors"
	"net/http"
	"net/url"
)

// The connection URL for the API resource.
const api = "https://slack.com/api/chat.postMessage"

// Writer implements the io.Writer interface and stores basic state.
type Writer struct {
	data url.Values
}

// NewWriter returns a new writer configured for use with the Slack API.
func NewWriter(channel, user, token string) (*Writer, error) {
	// Basic validation.
	if channel == "" || user == "" || token == "" {
		return &Writer{}, errors.New("invalid configuration")
	}
	// Pre-populate the static values for the data.
	d := url.Values{}
	d.Add("channel", channel)
	d.Add("username", user)
	d.Add("token", token)
	// Configure the writer with the details.
	w := Writer{
		data: d,
	}
	return &w, nil
}

// Write writes the input byte data to the nominated channel as a string.
func (w *Writer) Write(p []byte) (int, error) {
	w.data.Add("text", string(p))
	_, err := http.PostForm(api, w.data)
	// The errors and return code don't line up but are just
	// so that this can be called as a Writer.
	return len(p), err
}

// Channel is the name of the channel the writer will post to.
func (w *Writer) Channel() string {
	return w.data.Get("channel")
}

// Username is the name that the writer will appear as in the channel.
func (w *Writer) Username() string {
	return w.data.Get("username")
}

// Token is the access token that will be used for the Slack API.
func (w *Writer) Token() string {
	return w.data.Get("token")
}
