package products

import (
	"context"
	"pheet-tutorial-echo/models"

	"github.com/gofrs/uuid"
)

type ProductUsecase interface {
	GetProducts(ctx context.Context)([]*models.Products,error)
	GetProduct(id int)(*models.Products, error)
	GetProductByType(coffType string)([]*models.Products, error)
	GetUser(username string)(*models.User, error)
	Create(ctx context.Context, product *models.Products) error
	SignUp(ctx context.Context, user *models.User) error
	Update(ctx context.Context, product *models.Products, id *uuid.UUID) error
	Delete(id int)error
}