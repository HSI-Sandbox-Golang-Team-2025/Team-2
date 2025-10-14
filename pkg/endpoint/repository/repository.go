package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type FindEndpointCondition struct {
	endpoint.Endpoint
}

func (r *repository) GetEndpoint(ctx context.Context, e *endpoint.Endpoint, condition *FindEndpointCondition) error {
	err := r.db.WithContext(ctx).
		Where("path = ?", condition.Path).
		First(&e).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateEndpoint(ctx context.Context, e *endpoint.Endpoint) error {
	findEndpointCondition := FindEndpointCondition{Endpoint: *e}

	if err := r.GetEndpoint(ctx, e, &findEndpointCondition); err == nil {
		return errors.New("Endpoint is already registered")
	}

	if err := r.db.WithContext(ctx).Create(&e).Error; err != nil {
		return err
	}

	return nil
}
