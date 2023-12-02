package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/posener/cmd"
	"github.com/rs/zerolog"
	"github.com/swierq/golang/internal/kubek"
)

func main() {
	root := cmd.New()
	deplName := root.String("name", "deployment", "Deployment name..")
	debug := root.Bool("debug", false, "Debug logging")
	_ = root.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go kubek.Deployment(ctx, &wg, *deplName)
	wg.Wait()
	fmt.Println("Finished.")
}
