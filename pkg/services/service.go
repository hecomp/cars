package services

import (
	"github.com/hecomp/cars/internal/models"
	"github.com/hecomp/cars/pkg/repository"
)

type CarsService interface {
	GetCar(id string) (*models.Car, error)
	GetCars() []*models.Car
	Create(user *models.Car) (*models.Car, error)
	Update(user *models.Car) (*models.Car, error)
}

type carsService struct {
	repo repository.Repository
}

func NewCarsService(repo repository.Repository) CarsService {
	return &carsService{
		repo: repo,
	}
}

func (s carsService) GetCar(id string) (*models.Car, error) {
	car, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (s carsService) GetCars() []*models.Car {
	return s.repo.List()
}

func (s carsService) Create(car *models.Car) (*models.Car, error) {
	car, err := s.repo.Save(car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (s carsService) Update(car *models.Car) (*models.Car, error) {
	car, err := s.repo.Update(car)
	if err != nil {
		return nil, err
	}
	return car, nil
}
