package wzlib

const (
	CHANNEL_PUBLIC     = "public"     // Public channel for all clients
	CHANNEL_PRIVATE    = "private"    // Private channel: payload is encrypted
	CHANNEL_CONSOLE    = "remote"     // Remote controller is talking here and controller is listening to it
	CHANNEL_CONTROLLER = "controller" // Controller is listening here
	CHANNEL_CLIENT     = "client"     // Channel for clients responses

)

const (
	CLIENT_STATUS_NEW = iota
	CLIENT_STATUS_ACCEPTED
	CLIENT_STATUS_REJECTED
)
