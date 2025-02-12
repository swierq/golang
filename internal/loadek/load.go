package loadek

import (
	"context"
	"runtime"
	"sync"
	"time"
)

// GenerateCPULoad creates artificial load on CPU based on the given percentage and exits when the context is done.
func (a *App) generateCPULoad(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fullCores := a.Config.CpuLoadMi / 1000
	partialCoreMillicores := a.Config.CpuLoadMi % 1000
	runtime.GOMAXPROCS(fullCores + 1)
	var wg2 sync.WaitGroup
	wg2.Add(fullCores + 1)

	// Fully utilize the calculated number of cores
	for i := 0; i < fullCores; i++ {
		go func() {
			defer wg2.Done()
			runtime.LockOSThread()
			for {
				select {
				case <-ctx.Done():
					return
				//nolint:staticcheck
				default:
					// Busy-wait to fully utilize the core
				}
			}
		}()
	}

	// Utilize the remaining percentage on one core
	go func() {
		defer wg2.Done()
		runtime.LockOSThread()
		busyTime := time.Duration(partialCoreMillicores) * 100 * time.Microsecond
		idleTime := time.Duration(1000-partialCoreMillicores) * 100 * time.Microsecond

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Busy-wait
				start := time.Now()
				for time.Since(start) < busyTime {
				}
				// Idle
				time.Sleep(idleTime)
			}
		}
	}()

	wg2.Wait()
}

func (a *App) reserveMemory(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Calculate the number of bytes to reserve
	bytesToReserve := a.Config.MemLoadMb * 1024 * 1024

	// Allocate a slice of bytes
	memory := make([]byte, bytesToReserve)

	// Optionally fill the memory to ensure it's reserved
	for i := range memory {
		memory[i] = 1
	}

	// Keep the memory reserved until the context is done
	<-ctx.Done()
}

func (a *App) LoadSystem(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go a.generateCPULoad(ctx, &wg2)
	go a.reserveMemory(ctx, &wg2)
	wg2.Wait()
}
