package tests

import (
	"testing"
	"workhub/internal/models"
	"workhub/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMockFindJobByID(t *testing.T) {
	mockRepo := new(mocks.MockJobRepository)

	expectedJob := &models.Job{
		ID:       1,
		Title:    "Backend Developer",
		Category: "IT",
		Status:   "open",
	}

	mockRepo.On("FindByID", uint(1)).Return(expectedJob, nil)

	job, err := mockRepo.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedJob.Title, job.Title)
	assert.Equal(t, "IT", job.Category)

	mockRepo.AssertExpectations(t)
}

func TestMockDeleteJob(t *testing.T) {
	mockRepo := new(mocks.MockJobRepository)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := mockRepo.Delete(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
