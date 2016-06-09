# slack-writer

An `io.Writer` to send messages to a Slack channel.

```go
package main

import (
	"fmt"

	"github.com/leighlondon/slack-writer"
)

func main() {
	writer, err := slack.NewWriter("#channel-name", "the-bot-username", "your-token-here")
	if err != nil {
		fmt.Println("an invalid configuration with the writer")
		return
	}
	fmt.Fprintf(writer, "Hello world!")
}
```
