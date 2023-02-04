package repository

import (
	"fmt"
	"github.com/hecomp/cars/internal/models"
	"github.com/hecomp/cars/pkg/utils"
	"sync"
)

type carsDB struct {
	Storage map[string]*models.Car
}

type Repository interface {
	Find(id string) (*models.Car, error)
	List() []*models.Car
	Save(user *models.Car) (*models.Car, error)
	Update(user *models.Car) (*models.Car, error)
}

type repository struct {
	mutex *sync.Mutex
	carsDB
}

func NewRepository() Repository {
	var db carsDB
	db.Storage = make(map[string]*models.Car)
	return &repository{
		carsDB: db,
		mutex:  &sync.Mutex{},
	}
}

func (r repository) Find(id string) (*models.Car, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.Storage[id]; !ok {
		return nil, fmt.Errorf("car not found %v", id)
	}
	return r.Storage[id], nil
}

func (r repository) List() []*models.Car {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	cars := make([]*models.Car, 0, len(r.Storage))
	for _, car := range r.Storage {
		cars = append(cars, car)
	}
	return cars
}

func (r repository) Save(user *models.Car) (*models.Car, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.Storage[user.Id]; !ok {
		user.Id = utils.GenId(9)
		r.Storage[user.Id] = user
	} else {
		return nil, fmt.Errorf("duplicate car %v", user)
	}
	return user, nil
}

func (r repository) Update(user *models.Car) (*models.Car, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.Storage[user.Id]; !ok {
		return nil, fmt.Errorf("car not found %v", user)
	}
	r.Storage[user.Id] = user

	return r.Storage[user.Id], nil
}
