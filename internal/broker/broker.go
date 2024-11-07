package broker

import (
	"context"

	natspkg "github.com/CRED-CLUB/propeller/pkg/nats"
	redispkg "github.com/CRED-CLUB/propeller/pkg/redis"
)

// NewNATSClient returns a NATS client
func NewNATSClient(ctx context.Context, config Config) (*natspkg.Client, error) {
	var err error

	// start embedded nats server and set its url
	if config.Nats.EmbeddedServer {
		// override nats URL
		config.Nats.URL, err = natspkg.NewEmbeddedServer(ctx)
		if err != nil {
			return nil, err
		}
	}
	conn, err := natspkg.NewClient(ctx, config.Nats)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewRedisClient returns a redis client
func NewRedisClient(ctx context.Context, config Config) (*redispkg.Client, error) {
	redisClient := redispkg.NewClient(config.Redis)
	return redisClient, nil
}
