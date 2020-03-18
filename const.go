package wzlib_const

const (
	CHANNEL_PUBLIC     = "public"     // Public channel for all clients
	CHANNEL_PRIVATE    = "private"    // Private channel: payload is encrypted
	CHANNEL_CONSOLE    = "remote"     // Remote controller is talking here and controller is listening to it
	CHANNEL_CONTROLLER = "controller" // Controller is listening here
)
