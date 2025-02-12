package loadek

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestGenerateCPULoad(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	config, _ := NewConfig(WithCpuLoadMi(10), WithMemLoadMb(10))
	app := NewApp(WithConfig(config))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	startTime := time.Now()
	go app.generateCPULoad(ctx, &wg)

	// Wait for the load generation to complete
	wg.Wait()

	// Check if the function respects the context timeout
	elapsedTime := time.Since(startTime)
	if elapsedTime < 1*time.Second {
		t.Errorf("generateCPULoad finished too early, elapsed time: %v", elapsedTime)
	}
}

func TestReserveMemory(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, _ := NewConfig(WithCpuLoadMi(10), WithMemLoadMb(10))
	app := NewApp(WithConfig(config))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	startTime := time.Now()
	go app.reserveMemory(ctx, &wg)

	// Wait for the load generation to complete
	wg.Wait()

	// Check if the function respects the context timeout
	elapsedTime := time.Since(startTime)
	if elapsedTime < 1*time.Second {
		t.Errorf("reserveMemory finished too early, elapsed time: %v", elapsedTime)
	}
}

func TestLoadSystem(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, _ := NewConfig(WithCpuLoadMi(10), WithMemLoadMb(10))
	app := NewApp(WithConfig(config))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	startTime := time.Now()
	go app.LoadSystem(ctx, &wg)

	// Wait for the load generation to complete
	wg.Wait()

	// Check if the function respects the context timeout
	elapsedTime := time.Since(startTime)
	if elapsedTime < 1*time.Second {
		t.Errorf("LoadSystem finished too early, elapsed time: %v", elapsedTime)
	}
}
