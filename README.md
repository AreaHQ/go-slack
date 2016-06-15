[![Codeship Status for AreaHQ/go-slack](https://codeship.com/projects/7252c9a0-09f0-0134-e361-2adbeb910e90/status?branch=master)](https://codeship.com/projects/155402)

# Go Slack

A simple Golang SDK for Slack.

## Usage

```go
package main

import (
	"log"

	slack "github.com/AreaHQ/go-slack"
)

func main() {
	cnf := &slack.Config{IncomingWebhook: "incoming_webhook"}
	adapter := slack.NewAdapter(cnf)
	err = adapter.SendMessage(
		"#some-channel",
		"some-username",
		"message to send",
		"", // emmoji
	)
}
```
