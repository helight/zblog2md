package config

import (
	"flag"
	"os"
)

// Config is a configuration interface
type Config interface {
	IsSet(name string) bool
	Bool(name string) bool
	Int(name string) int
	IntSlice(name string) []int
	Int64(name string) int64
	Int64Slice(name string) []int64
	String(name string) string
	StringSlice(name string) []string
	Uint(name string) uint
	Uint64(name string) uint64
	Set(name, value string) error
}

var (
	initializers []func(Config)
	config       Config
	Auth         = os.Getenv("AUTH") == "true"
	Online       = flag.Bool("online", false, "online flag")
)

func IsSet(name string) bool           { return config.IsSet(name) }
func Bool(name string) bool            { return config.Bool(name) }
func Int(name string) int              { return config.Int(name) }
func IntSlice(name string) []int       { return config.IntSlice(name) }
func Int64(name string) int64          { return config.Int64(name) }
func Int64Slice(name string) []int64   { return config.Int64Slice(name) }
func String(name string) string        { return config.String(name) }
func StringSlice(name string) []string { return config.StringSlice(name) }
func Uint(name string) uint            { return config.Uint(name) }
func Uint64(name string) uint64        { return config.Uint64(name) }
func Set(name, value string) error     { return config.Set(name, value) }

// AddInitializer Add a initializer, call on initialized
func AddInitializer(fc func(Config)) {
	initializers = append(initializers, fc)
}

// Initialize initialize process configure
func Initialize(c Config) {
	config = c
	initializeDatabase("dev")

	for _, initFunc := range initializers {
		initFunc(c)
	}
}

func setDefault(config Config, name, value string) {
	if config.IsSet(name) {
		return
	}
	config.Set(name, value)
}
