package products

import (
	"context"
	"pheet-tutorial-echo/models"

	"github.com/gofrs/uuid"
)

type ProductRepository interface {
	FetchAll(ctx context.Context) ([]*models.Products, error)
	FetchById(id int) (*models.Products, error)
	FetchByType(coffType string) ([]*models.Products, error)
	FetchUser(username string) (*models.User, error)
	Create(ctx context.Context, product *models.Products) error
	SignUp(ctx context.Context, user *models.User) error
	Update(ctx context.Context, product *models.Products, id *uuid.UUID) error
	Delete(id int) error
}
