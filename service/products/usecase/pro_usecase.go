package usecase

import (
	"context"
	"pheet-tutorial-echo/models"
	"pheet-tutorial-echo/service/products"

	"github.com/gofrs/uuid"
)

type productUsecase struct {
	proRepo products.ProductRepository
}

func NewProductUsecase(proRepo products.ProductRepository) products.ProductUsecase {
	return productUsecase{proRepo: proRepo}
}

func (r productUsecase)GetProducts(ctx context.Context)([]*models.Products,error){
	return r.proRepo.FetchAll(ctx)	
}

func (r productUsecase)GetProduct(id int)(*models.Products, error){
	return r.proRepo.FetchById(id)
}

func (r productUsecase)GetUser(username string)(*models.User, error){
	return r.proRepo.FetchUser(username)
}

func (r productUsecase)GetProductByType(coffType string)([]*models.Products, error){
	return r.proRepo.FetchByType(coffType)
}

func (r productUsecase)Create(ctx context.Context, product *models.Products) error{
	return r.proRepo.Create(ctx, product)
}

func (r productUsecase)SignUp(ctx context.Context, user *models.User) error {
	return r.proRepo.SignUp(ctx, user)
}

func(r productUsecase)Update(ctx context.Context, product *models.Products, id *uuid.UUID) error{
	return r.proRepo.Update(ctx, product, id)
}

func(r productUsecase)Delete(id int)error{
	return r.proRepo.Delete(id)
}