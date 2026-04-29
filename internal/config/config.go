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
	AssetURL        string
	SizeLimit       int64
	BufferSize      int
	UpstreamTimeout time.Duration
	ShutdownTimeout time.Duration
}

// Load 从环境变量读取配置，并设置默认值。
func Load() *Config {
	return &Config{
		Listen:          env("LISTEN", ":8080"),
		AssetURL:        env("ASSET_URL", "https://hunshcn.github.io/gh-proxy"),
		SizeLimit:       envInt64("SIZE_LIMIT", 1072668082176),
		BufferSize:      int(envInt64("BUFFER_SIZE", 32*1024)),
		UpstreamTimeout: envDuration("UPSTREAM_TIMEOUT", 30*time.Second),
		ShutdownTimeout: envDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt64(key string, def int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return def
	}
	return n
}

func envDuration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}
