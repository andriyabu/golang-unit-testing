package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (Service CategoryService) Get(id string) (*entity.Category, error) {
	category := Service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found")

	} else {
		return category, nil
	}
}
