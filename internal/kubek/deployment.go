package kubek

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func Deployment(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx done, exiting.")
			return
		case <-ticker.C:
			log.Info().Msgf("Deployment: %s", name)
		}
	}
}
