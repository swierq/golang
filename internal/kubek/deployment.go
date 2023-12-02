package kubek

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func DeploymentScaler(ctx context.Context, wg *sync.WaitGroup, cfgFile string, interval int) {
	defer wg.Done()
	tickerFunc(cfgFile)
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx done, exiting.")
			return
		case <-ticker.C:
			tickerFunc(cfgFile)
		}
	}
}

type ScalerConfig struct {
	Name string `yaml:"name"`
}

func tickerFunc(cfgFile string) {
	log.Info().Msgf("Reading Config: %s", cfgFile)
	cfg, err := readScalerConfig(cfgFile)

	if err != nil {
		log.Error().Msg("Could not read config file, continuing.")
		return
	}
	log.Info().Msgf("Config: %v", cfg)
}

func readScalerConfig(cfgFile string) (ScalerConfig, error) {
	cfg := ScalerConfig{}
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Error().Msgf("error: %v", err)
		return cfg, err
	}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Error().Msgf("error: %v", err)
		return cfg, err
	}
	return cfg, nil
}
