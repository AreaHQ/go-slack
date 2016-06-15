package slack

// AdapterInterface defines exported methods
type AdapterInterface interface {
	// Exported methods
	SendMessage(incomingWebhook, channel, username, text, emoji string) error
}
