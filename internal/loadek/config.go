package loadek

import "fmt"

type Config struct {
	CpuLoadMi int
	MemLoadMb int
}

type ConfigOption func(*Config) error

func NewConfig(options ...ConfigOption) (*Config, error) {
	cfg := &Config{
		CpuLoadMi: 100,
		MemLoadMb: 100,
	}

	for _, option := range options {
		err := option(cfg)
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func WithCpuLoadMi(cpuLoadMi int) ConfigOption {
	return func(c *Config) error {
		if cpuLoadMi < 0 {
			return fmt.Errorf("cpuLoadMi must be greater than 0")
		}
		c.CpuLoadMi = cpuLoadMi
		return nil
	}
}

func WithMemLoadMb(memLoadMb int) ConfigOption {
	return func(c *Config) error {
		if memLoadMb < 0 {
			return fmt.Errorf("memLoadMb must be greater than 0")
		}
		c.MemLoadMb = memLoadMb
		return nil
	}
}
