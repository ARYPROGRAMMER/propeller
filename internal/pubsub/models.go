package pubsub

// PublishRequest is the publishRequest
type PublishRequest struct {
	Channel string
	Data    []byte
}
