package SimpleTest

import "time"

// To create a covarage test you can run
// go test -cover ./... -coverprofile=coverage.out
// and to see it in the html file
// go tool cover -html coverage
func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	//fix the first fuzz error
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}

// Create to use in benchmark tests
func CalculateTaxTime(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
