package config

import "flag"

type Flags struct {
	DryRun  bool
	Verbose bool
}

type Config struct {
	InputPath string
	Flags
}

func NewConfigFromFlags() *Config {
	config := new(Config)
	flag.StringVar(&config.InputPath, "i", ".", "input path [file/dir]")
	flag.BoolVar(&config.Verbose, "v", true, "verbose mode")
	flag.BoolVar(&config.DryRun, "d", false, "dry-run mode")
	flag.Parse()
	return config
}
