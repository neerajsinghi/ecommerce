package model

import "time"

func GetJWTSecret() string {
	return "secret"
}

func GetJWTExpiration() time.Duration {
	return time.Hour * 24 * 30
}
