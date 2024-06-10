package main

import (
	"os"

	"github.com/amnestia/tnderlike/internal/api/server"
)

func main() {
	os.Exit(server.New().Run())
}
