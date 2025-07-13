package main

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v11"
)

type envValues struct {
	DatabaseName     string `env:"DATABASE_NAME,required"`
	DatabaseHostname string `env:"DATABASE_HOSTNAME,required"`
	DatabasePort     string `env:"DATABASE_PORT,required"`
	DatabaseUsername string `env:"DATABASE_USERNAME"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseSSLMode  bool   `env:"DATABASE_SSL_MODE"             envDefault:"true"`
}

func main() {
	var envValues envValues

	if err := env.Parse(&envValues); err != nil {
		panic(err)
	}

	url := databaseURL(envValues)

	fmt.Printf("%s\n", url)
}

func databaseURL(envValues envValues) string {
	connString := "postgres://"

	if envValues.DatabaseUsername != "" {
		connString = fmt.Sprintf("%s%s", connString, envValues.DatabaseUsername)

		if envValues.DatabasePassword != "" {
			encodedPassword := url.QueryEscape(envValues.DatabasePassword)
			connString = fmt.Sprintf("%s:%s", connString, encodedPassword)
		}

		connString += "@"
	}

	sslMode := "disable"
	if envValues.DatabaseSSLMode {
		sslMode = "require"
	}

	connString = fmt.Sprintf(
		"%s%s:%s/%s?sslmode=%s",
		connString,
		envValues.DatabaseHostname,
		envValues.DatabasePort,
		envValues.DatabaseName,
		sslMode,
	)

	return connString
}
