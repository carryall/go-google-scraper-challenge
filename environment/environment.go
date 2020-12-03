package environment

import (
	"os"
	"strings"
)

const (
	PROD Env = "production"
	DEV  Env = "development"
)

type Env string

func (e Env) String() string {
	return string(e)
}

func ReadEnv(key string, def Env) Env {
	v := Getenv(key, def.String())
	if v == "" {
		return def
	}

	env := Env(strings.ToLower(v))
	switch env {
	case PROD, DEV: // allowed.
	default:
		panic("unexpected environment " + v)
	}

	return env
}

func Getenv(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
