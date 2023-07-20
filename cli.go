package main

import (
	"github.com/alexflint/go-arg"
)

// Config struct to hold the configuration data
type Args struct {
	ConfigFile string `arg:"positional,required" help:"Path to the YAML config file"`
	IsDryRun   bool   `arg:"--dryrun" help:"Set this flag to enable dry run mode"`
	Verbose    bool   `arg:"--verbose,-v" help:"Show what is happening"`
}

func (Config) Description() string {
	return "A simple hash based photo collation and merge program for UNIX systems."
}

func cli() Args {
	var args Args
	arg.MustParse(&args)
	return args
}
