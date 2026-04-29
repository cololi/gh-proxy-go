// Package config 处理从环境变量加载的运行配置。
package config

import (
	"os"
	"strconv"
	"time"
)

// Config 存储从环境变量加载的所有配置。
type Config struct {
	Listen          string
	SizeLimit       int64
	BufferSize      int
	UpstreamTimeout time.Duration
	ShutdownTimeout time.Duration
}

// Load 从环境变量读取配置，并设置默认值。
func Load() *Config {
	return &Config{
		Listen:          getEnv("LISTEN", ":8080"),
		SizeLimit:       getEnvInt64("SIZE_LIMIT", 1072668082176),
		BufferSize:      int(getEnvInt64("BUFFER_SIZE", 32*1024)),
		UpstreamTimeout: getEnvDuration("UPSTREAM_TIMEOUT", 30*time.Second),
		ShutdownTimeout: getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return n
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return defaultValue
	}
	return d
}
