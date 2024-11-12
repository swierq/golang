package webapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *App) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.GetRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		app.Logger.Info("shutting down server", "signal", s.String())
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.Logger.Info("starting server", "addr", srv.Addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.Logger.Info("stopped server", "addr", srv.Addr)
	return nil
}
