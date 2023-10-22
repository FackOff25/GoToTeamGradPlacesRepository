package main

import (
	"flag"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/app"
)

const _DEFAULT_CONFIG_FILE_PATH = "configs/config.yaml"

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config_file_path", _DEFAULT_CONFIG_FILE_PATH, "config file path")
}

func main() {
	flag.Parse()

	app.Run(configFilePath)
}
