package domain

import (
	"errors"
	"sort"
)

var computePackagesFunc = computePackages

// CalculatePacks determines the minimum combination of packages needed to reach at least
// the target number of items. It returns a map where keys are package sizes and values are
// the counts of packages used.
func CalculatePacks(target int, packageSizes []int) (map[int]int, error) {
	if target <= 0 {
		return nil, errors.New("target must be greater than zero")
	}
	// Ensure there is at least one package size provided.
	if len(packageSizes) == 0 {
		return nil, errors.New("no package sizes provided")
	}

	// Sort package sizes in ascending order.
	sort.Ints(packageSizes)

	return computePackagesFunc(target, packageSizes)
}

// computePackages applies a dynamic programming approach to determine the fewest packages
// needed so that their sum is at least the target. It returns a map detailing the combination.
func computePackages(target int, packageSizes []int) (map[int]int, error) {
	// Validate that each package size is positive.
	for _, size := range packageSizes {
		if size <= 0 {
			return nil, errors.New("package sizes must be greater than zero")
		}
	}

	// The largest available package is the last element (since the slice is sorted).
	largest := packageSizes[len(packageSizes)-1]
	// Define an upper limit for our search: target plus the largest package.
	limit := target + largest

	// Initialize the dp array, where dp[i] will hold the minimum number of packages
	// needed to exactly reach i items.
	dp := make([]int, limit+1)
	// prevPackage records which package size was used to achieve a given total.
	prevPackage := make([]int, limit+1)
	const INF = 1e9

	// Set initial values: "infinity" for dp and -1 for prevPackage.
	for i := 0; i <= limit; i++ {
		dp[i] = INF
		prevPackage[i] = -1
	}
	dp[0] = 0 // Base case: zero items require zero packages.

	// Build the dp table: for every reachable total, try adding each package.
	for i := 0; i <= limit; i++ {
		if dp[i] == INF {
			continue // Skip unreachable totals.
		}
		for _, size := range packageSizes {
			next := i + size
			// If next is within our limit and a better (fewer packages) solution is found:
			if next <= limit && dp[i]+1 < dp[next] {
				dp[next] = dp[i] + 1
				prevPackage[next] = size
			}
		}
	}

	// Find the smallest total at or above target that can be reached.
	bestTotal := -1
	for i := target; i <= limit; i++ {
		if dp[i] < INF {
			bestTotal = i
			break
		}
	}
	if bestTotal == -1 {
		return nil, errors.New("could not achieve the target with the provided package sizes")
	}

	// Backtrack to reconstruct which packages were used to reach bestTotal.
	solution := make(map[int]int)
	remaining := bestTotal
	for remaining > 0 {
		used := prevPackage[remaining]
		if used == -1 {
			break // Safety check: if no package is recorded, exit.
		}
		solution[used]++
		remaining -= used
	}

	return solution, nil
}
