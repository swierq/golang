package loadek

import (
	"fmt"
	"math"
)

// CPULoad creates artificial load on CPU.
func CPULoad(number int64) (string, error) {
	result, err := isPrime(number)
	return fmt.Sprintf("%t", result), err
}

func isPrime(number int64) (bool, error) {
	sqrt := int64(math.Ceil(math.Pow(float64(number), 0.5)))
	for i := int64(2); i < sqrt; i++ {
		if number%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
