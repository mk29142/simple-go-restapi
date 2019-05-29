package repository

import (
	"context"
	"simple-go-restapi/app/models"
)

type UserRepo interface {
	GetByName(ctx context.Context, name string) (*models.User, error)
	Create(ctx context.Context, p *models.User) error
}
