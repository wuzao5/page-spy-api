package config

import (
	"io/fs"
)

type CorsConfig struct {
	AllowOrigins  []string `json:"allowOrigins"`
	AllowMethods  []string `json:"allowMethods"`
	AllowHeaders  []string `json:"allowHeaders"`
	ExposeHeaders []string `json:"exposeHeaders"`
}

type Config struct {
	Port                string      `json:"port"`
	Debug               bool        `json:"debug"`
	NotAllowedDeleteLog bool        `json:"notAllowedDeleteLog"`
	RpcAddress          []*Address  `json:"rpcAddress"`
	CorsConfig          *CorsConfig `json:"corsConfig"`
	MaxRoomNumber       int         `json:"maxRoomNumber"`
	// max log file size, unit is mb
	MaxLogFileSizeOfMB int64 `json:"maxLogFileSizeOfMB"`
	// max log file size, unit is day
	MaxLogLifeTimeOfHour int64 `json:"maxLogLifeTimeOfHour"`
}

func (c *Config) GetMaxLogLifeTimeOfHour() int64 {
	if c.MaxLogLifeTimeOfHour <= 0 {
		return 30 * 24 // default log life 30 day
	}

	return c.MaxLogLifeTimeOfHour
}

func (c *Config) GetMaxLogFileSizeOfMB() int64 {
	if c.MaxLogFileSizeOfMB <= 0 {
		return 10 * 1024 // default log size 10GB
	}

	return c.MaxLogFileSizeOfMB
}

func (c *Config) GetMaxRoomNumber() int {
	if c.MaxRoomNumber <= 0 {
		return 500
	}

	return c.MaxRoomNumber
}

type Address struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type StaticConfig struct {
	DirName string
	Files   fs.FS
	GitHash string
	Version string
}
