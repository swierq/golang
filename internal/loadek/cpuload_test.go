package loadek

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	app := NewApp()
	result := app.isPrime(17)
	require.Equal(t, true, result, "result should be True")

	result = app.isPrime(20)
	require.Equal(t, false, result, "result should be False")
}

func TestCPULoad(t *testing.T) {
	app := NewApp()
	result, err := app.CPULoad(2)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, "[2 3]", result, "result should be True")

	result, err = app.CPULoad(6)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, "[2 3 5 7 11 13]", result, "result should be True")

}
