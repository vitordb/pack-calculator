package application

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"pack-calculator/internal/domain/ports"
)

type mockRepoSuccess struct {
	lastSaved *ports.Result
}

func (m *mockRepoSuccess) SaveResult(result *ports.Result) error {
	m.lastSaved = result
	return nil
}

func (m *mockRepoSuccess) GetResultByID(id string) (*ports.Result, error) {
	return m.lastSaved, nil
}

type mockRepoFailure struct{}

func (m *mockRepoFailure) SaveResult(result *ports.Result) error {
	return errors.New("repo failure")
}

func (m *mockRepoFailure) GetResultByID(id string) (*ports.Result, error) {
	return nil, errors.New("repo failure")
}

func TestCalculatorService_Valid(t *testing.T) {
	repo := &mockRepoSuccess{}
	service := NewCalculatorService(repo)

	target := 500000
	packageSizes := []int{23, 31, 53}

	result, err := service.CalculateAndSave(target, packageSizes)
	assert.NoError(t, err, "expected no error for valid input")
	assert.NotNil(t, result, "expected a result for valid input")

	expectedSolution := map[int]int{
		23: 2,
		31: 7,
		53: 9429,
	}
	assert.Equal(t, target, result.Amount, "target should match input")
	assert.Equal(t, packageSizes, result.PackSizes, "packageSizes should match input")
	assert.Equal(t, expectedSolution, result.Solution, "solution should match expected output")
	assert.Equal(t, result, repo.lastSaved, "repository should have saved the result")
}

func TestCalculatorService_CalculateError(t *testing.T) {
	repo := &mockRepoSuccess{}
	service := NewCalculatorService(repo)

	target := -1
	packageSizes := []int{23, 31, 53}

	result, err := service.CalculateAndSave(target, packageSizes)
	assert.Nil(t, result, "expected nil result when business logic fails")
	assert.Error(t, err, "expected an error when business logic fails")
	assert.Equal(t, "target must be greater than zero", err.Error(), "error message should match")
}

func TestCalculatorService_RepoError(t *testing.T) {
	repo := &mockRepoFailure{}
	service := NewCalculatorService(repo)

	target := 500000
	packageSizes := []int{23, 31, 53}

	result, err := service.CalculateAndSave(target, packageSizes)
	assert.Nil(t, result, "expected nil result when repository fails")
	assert.Error(t, err, "expected an error when repository fails")
	assert.Equal(t, "repo failure", err.Error(), "error message should match repository error")
}
