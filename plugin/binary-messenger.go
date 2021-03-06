package plugin

// BinaryMessenger defines a bidirectional binary messenger.
type BinaryMessenger interface {
	// Send sends a binary message to the Flutter application.
	Send(channel string, binaryMessage []byte) (binaryReply []byte, err error)

	// SetChannelHandler registers a handler to be invoked when the Flutter
	// application sends a message to its host platform on given channel.
	//
	// Registration overwrites any previous registration for the same channel
	// name. Use nil as handler to deregister.
	SetChannelHandler(channel string, handler ChannelHandlerFunc)
}

// ResponseSender defines the interface that must be implemented by a messenger
// to handle replies on on message. It is an error to call Send multiple times
// on the same ResponseSender.
type ResponseSender interface {
	Send(binaryReply []byte)
}

// ChannelHandlerFunc describes the function that handles binary messages sent
// on a channel. For each message, ResponseSender.Send must be called once.
type ChannelHandlerFunc func(binaryMessage []byte, r ResponseSender) (err error)
