package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/sameer1703/csv-wsuitemetric-beat/beater"
)

func main() {
	err := beat.Run("csvbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
