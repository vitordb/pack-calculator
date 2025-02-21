package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePacks_Valid(t *testing.T) {
	target := 500000
	packageSizes := []int{23, 31, 53}

	result, err := CalculatePacks(target, packageSizes)
	assert.NoError(t, err, "Unexpected error for valid input")

	expected := map[int]int{
		23: 2,
		31: 7,
		53: 9429,
	}
	assert.Equal(t, expected, result, "Solution does not match expected output")
}

func TestCalculatePacks_ZeroOrNegativeTarget(t *testing.T) {
	tests := []struct {
		name        string
		target      int
		expectedErr string
	}{
		{"ZeroTarget", 0, "target must be greater than zero"},
		{"NegativeTarget", -5, "target must be greater than zero"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CalculatePacks(tc.target, []int{10, 20})
			assert.Nil(t, result, "Expected nil result for %s", tc.name)
			assert.Error(t, err, "Expected an error for %s", tc.name)
			assert.Equal(t, tc.expectedErr, err.Error(), "Error message should match for %s", tc.name)
		})
	}
}

func TestCalculatePacks_NoPackageSizes(t *testing.T) {
	result, err := CalculatePacks(100, []int{})
	assert.Nil(t, result, "Expected nil result when no package sizes provided")
	assert.Error(t, err, "Expected an error when no package sizes provided")
	assert.Equal(t, "no package sizes provided", err.Error(), "Error message should match")
}

func TestCalculatePacks_WithSmallTarget(t *testing.T) {
	target := 5
	packageSizes := []int{4, 6}

	result, err := CalculatePacks(target, packageSizes)
	assert.NoError(t, err, "Did not expect an error for a small target")
	expected := map[int]int{
		6: 1,
	}
	assert.Equal(t, expected, result, "Solution does not match expected output for a small target")
}

func TestCalculatePacks_InvalidPackageSizes(t *testing.T) {
	target := 5
	packageSizes := []int{-2, -3}

	result, err := CalculatePacks(target, packageSizes)
	assert.Nil(t, result, "Expected nil result for negative package sizes")
	assert.Error(t, err, "Expected an error for negative package sizes")
	expectedErr := "package sizes must be greater than zero"
	assert.Equal(t, expectedErr, err.Error(), "Error message should match for invalid package sizes")
}

func TestCalculatePacks_ComputePackagesFuncError(t *testing.T) {
	originalFunc := computePackagesFunc
	defer func() { computePackagesFunc = originalFunc }()

	computePackagesFunc = func(target int, packageSizes []int) (map[int]int, error) {
		return nil, errors.New("simulated error in computePackagesFunc")
	}

	target := 500000
	packageSizes := []int{23, 31, 53}

	result, err := CalculatePacks(target, packageSizes)
	assert.Nil(t, result, "Expected nil result when computePackagesFunc fails")
	assert.Error(t, err, "Expected error when computePackagesFunc fails")
	assert.Equal(t, "simulated error in computePackagesFunc", err.Error(), "Error message should match simulated error")
}
