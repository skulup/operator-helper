package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Float64Or get the envVar as float64 otherwise the default
func Float64Or(envVar string, def float64) float64 {
	if val := Value(envVar); val != "" {
		return RequireFloat64(val)
	}
	return def
}

// Int32Or get the envVar as int32 otherwise the default
func Int32Or(envVar string, def int32) int32 {
	return int32(Int64Or(envVar, int64(def)))
}

// Int64Or get the envVar as int64 otherwise the default
func Int64Or(envVar string, def int64) int64 {
	if val := Value(envVar); val != "" {
		return int64(RequireFloat64(val))
	}
	return def
}

// RequireFloat64 returns the env variable as int64 or panic
func RequireFloat64(envVar string) float64 {
	i, err := strconv.ParseFloat(RequireValue(envVar), 64)
	if err != nil {
		panic(err)
	}
	return i
}

// Value returns the environment value with white space trimmed
func Value(envVar string) string {
	return strings.TrimSpace(os.Getenv(envVar))
}

// ValueOr returns the value of Value or default if none exists
func ValueOr(envVar, def string) string {
	if val := strings.TrimSpace(os.Getenv(envVar)); val != "" {
		return val
	}
	return def
}

// RequireValue returns the env Value or panic if none exists
func RequireValue(envVar string) string {
	if val := Value(envVar); val != "" {
		return val
	}
	log.Fatalf(fmt.Sprintf("expecting value for environment variable: %s", envVar))
	return ""
}

// RandomString generates a random base64 string of length len or err
func RandomString(len int) (string, error) {
	bitsNeeded := len * 6
	bytesNeeded := math.Ceil(float64(bitsNeeded) / 8)
	bs, err := RandomBytes(int(bytesNeeded))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bs)[:len], nil
}

// RandomBytes generates a random bytes of len len or error
func RandomBytes(len int) ([]byte, error) {
	bs := make([]byte, len)
	if _, err := rand.Read(bs); err != nil {
		return nil, err
	}
	return bs, nil
}
