package loadek

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	result, err := isPrime(17)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, true, result, "result should be True")

	result, err = isPrime(20)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, false, result, "result should be False")
}

func TestCPULoad(t *testing.T) {
	result, err := CPULoad(17)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, "true", result, "result should be True")

	result, err = CPULoad(20)
	require.Nil(t, err, "There should be no error")
	require.Equal(t, "false", result, "result should be False")
}
