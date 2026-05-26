package mocks

import (
	"workhub/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) Create(
	job *models.Job,
) error {

	args := m.Called(job)

	return args.Error(0)
}

func (m *MockJobRepository) FindAll() (
	[]models.Job,
	error,
) {

	args := m.Called()

	return args.Get(0).([]models.Job),
		args.Error(1)
}

func (m *MockJobRepository) FindByID(
	id uint,
) (*models.Job, error) {

	args := m.Called(id)

	return args.Get(0).(*models.Job),
		args.Error(1)
}

func (m *MockJobRepository) Update(
	job *models.Job,
) error {

	args := m.Called(job)

	return args.Error(0)
}

func (m *MockJobRepository) Delete(
	id uint,
) error {

	args := m.Called(id)

	return args.Error(0)
}
