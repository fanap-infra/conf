package conf

import (
	"math/big"
	"time"
)

func GetString(path string, defaultVal ...string) string {
	return conf.GetString(path, defaultVal...)
}

func GetBoolean(path string, defaultVal ...bool) bool {
	return conf.GetBoolean(path, defaultVal...)
}

func GetInt32(path string, defaultVal ...int32) int32 {
	return conf.GetInt32(path, defaultVal...)
}

func GetInt64(path string, defaultVal ...int64) int64 {
	return conf.GetInt64(path, defaultVal...)
}

func GetFloat32(path string, defaultVal ...float32) float32 {
	return conf.GetFloat32(path, defaultVal...)
}

func GetFloat64(path string, defaultVal ...float64) float64 {
	return conf.GetFloat64(path, defaultVal...)
}

func GetTimeDuration(path string, defaultVal ...time.Duration) time.Duration {
	return conf.GetTimeDuration(path, defaultVal...)
}

func GetStringList(path string) []string {
	return conf.GetStringList(path)
}

func GetBooleanList(path string) []bool {
	return conf.GetBooleanList(path)
}

func GetInt32List(path string) []int32 {
	return conf.GetInt32List(path)
}

func GetInt64List(path string) []int64 {
	return conf.GetInt64List(path)
}

func GetFloat32List(path string) []float32 {
	return conf.GetFloat32List(path)
}

func GetFloat64List(path string) []float64 {
	return conf.GetFloat64List(path)
}

func GetByteList(path string) []byte {
	return conf.GetByteList(path)
}

func GetByteSize(path string) *big.Int {
	return conf.GetByteSize(path)
}

func GetKeys(path string) []string {
	nodes := conf.GetNode(path)
	if nodes == nil {
		return []string{}
	}
	if nodes.IsObject() {
		return nodes.GetObject().GetKeys()
	}

	return nil
}
