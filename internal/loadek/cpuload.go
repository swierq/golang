package loadek

import (
	"fmt"
	"math"
)

// CPULoad creates artificial load on CPU.
func (a *App) CPULoad(number int) (string, error) {
	result := a.countPrimes(number)
	return fmt.Sprintf("%v", result), nil
}

func (a *App) isPrime(number int64) bool {
	sqrt := int64(math.Ceil(math.Pow(float64(number), 0.5)))
	for i := int64(2); i <= sqrt && i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func (a *App) countPrimes(count int) []int64 {
	primeList := make([]int64, 0, count)
	var candidate int64 = 2
	numberOfPrimes := 0
	for numberOfPrimes < count {
		if a.isPrime(candidate) {
			primeList = append(primeList, candidate)
			numberOfPrimes += 1
		}
		candidate += 1
	}
	return primeList
}
