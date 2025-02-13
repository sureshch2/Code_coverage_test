package main

import (
	"os"
	"testing"
	"sync"
)

var githubToken = os.Getenv("GITHUB_TOKEN")
var codecovToken = os.Getenv("CODECOV_TOKEN")
var org = "openshift"

// Benchmark for getAllRepos function
func BenchmarkGetAllRepos(b *testing.B) {
	// Set up for benchmark
	if githubToken == "" || codecovToken == "" {
		b.Skip("Skipping tests as GITHUB_TOKEN or CODECOV_TOKEN is not set")
		return
	}

	// Run the benchmark multiple times
	for i := 0; i < b.N; i++ {
		_, err := getAllRepos(org, githubToken)
		if err != nil {
			b.Errorf("Error fetching repos: %v", err)
		}
	}
}

// Benchmark for getRepoCoverage function
func BenchmarkGetRepoCoverage(b *testing.B) {
	// Set up for benchmark
	if githubToken == "" || codecovToken == "" {
		b.Skip("Skipping tests as GITHUB_TOKEN or CODECOV_TOKEN is not set")
		return
	}

	// Example repo to benchmark
	repo := "configure-goalert-operator"

	// Run the benchmark multiple times
	for i := 0; i < b.N; i++ {
		_, _ = getRepoCoverage(org, repo, codecovToken)
	}
}

// Benchmark for concurrent fetching of repo coverages and generating CSV reports
func BenchmarkConcurrentRepoCoverage(b *testing.B) {
	// Set up for benchmark
	if githubToken == "" || codecovToken == "" {
		b.Skip("Skipping tests as GITHUB_TOKEN or CODECOV_TOKEN is not set")
		return
	}

	repos, err := getAllRepos(org, githubToken)
	if err != nil {
		b.Errorf("Error getting repositories: %v", err)
		return
	}

	// Run the benchmark multiple times
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var wg sync.WaitGroup

			for _, repo := range repos {
				wg.Add(1)
				go func(repo string) {
					defer wg.Done()
					// Fetch coverage for each repo
					_, _ = getRepoCoverage(org, repo, codecovToken)
					// Generate CSV reports if verbose mode is enabled
					// If you want to include verbose, ensure that part is handled efficiently.
				}(repo)
			}
			wg.Wait()
		}
	})
}

// Benchmark the total time to run the main logic (can be longer duration)
func BenchmarkMainLogic(b *testing.B) {
	// Set up for benchmark
	if githubToken == "" || codecovToken == "" {
		b.Skip("Skipping tests as GITHUB_TOKEN or CODECOV_TOKEN is not set")
		return
	}

	b.ResetTimer()

	// Run the main function multiple times to measure the overall performance
	for i := 0; i < b.N; i++ {
		main() // This will run the entire main function logic
	}
}
