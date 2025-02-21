package application

import (
	"pack-calculator/internal/domain"
	"pack-calculator/internal/domain/ports"
)

// CalculatorService orchestrates business logic (calculation) and persistence.
type CalculatorService struct {
	Repo ports.DBInterface
}

// NewCalculatorService creates a new CalculatorService.
func NewCalculatorService(repo ports.DBInterface) *CalculatorService {
	return &CalculatorService{Repo: repo}
}

// CalculateAndSave computes the solution and persists the result.
func (s *CalculatorService) CalculateAndSave(amount int, packSizes []int) (*ports.Result, error) {
	solution, err := domain.CalculatePacks(amount, packSizes)
	if err != nil {
		return nil, err
	}

	result := &ports.Result{
		Amount:    amount,
		PackSizes: packSizes,
		Solution:  solution,
	}

	if err := s.Repo.SaveResult(result); err != nil {
		return nil, err
	}

	return result, nil
}
