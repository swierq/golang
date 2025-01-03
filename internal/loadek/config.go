package loadek

type Config struct {
	CpuLoadMi int
	MemLoadMb int
}

type ConfigOption func(*Config)

func NewConfig(options ...ConfigOption) *Config {
	cfg := &Config{
		CpuLoadMi: 100,
		MemLoadMb: 100,
	}

	for _, option := range options {
		option(cfg)
	}
	return cfg
}

func WithCpuLoadMi(cpuLoadMi int) ConfigOption {
	return func(c *Config) {
		c.CpuLoadMi = cpuLoadMi
	}
}

func WithMemLoadMb(memLoadMb int) ConfigOption {
	return func(c *Config) {
		c.MemLoadMb = memLoadMb
	}
}
