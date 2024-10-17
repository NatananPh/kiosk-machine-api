package repository

type productRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}