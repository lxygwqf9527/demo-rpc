package client

import (
	"context"

	"github.com/lxygwqf9527/demo-rpc/grpc/middleware/server"
)

func NewAuthentication(ak, sk string) *Authentication {
	return &Authentication{
		clientId:     ak,
		clientSecret: sk,
	}
}

type Authentication struct {
	clientId     string
	clientSecret string
}

// WithClientCredentials todo
func (a *Authentication) build() map[string]string {
	return map[string]string{
		server.ClientHeaderKey: a.clientId,
		server.ClientSecretKey: a.clientSecret,
	}
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {

	return a.build(), nil

}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
