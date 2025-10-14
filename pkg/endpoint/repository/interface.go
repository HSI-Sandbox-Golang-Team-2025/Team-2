package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint"
)

type Repository interface {
	GetEndpoint(ctx context.Context, e *endpoint.Endpoint, condition *FindEndpointCondition) error
	CreateEndpoint(ctx context.Context, e *endpoint.Endpoint) error
}
