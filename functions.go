package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

type Package struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Scripts map[string]string `json:"scripts"`
}

func hasExtension(filename string) bool {
	return filepath.Ext(filename) != ""
}

func NpmVersionExtractor() {
	fileExists := FileExists("VERSION")
	if fileExists {
		file, err := os.Open("VERSION")
		if err != nil {
			log.Error().Err(err).Msg("error opening VERSION file")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		npmVersion = scanner.Text()

		return
	} else {
		file, err := os.ReadFile("package.json")
		if err != nil {
			log.Error().Err(err).Msg("error reading package.json")
			return
		}

		pac := Package{}

		err = json.Unmarshal(file, &pac)
		if err != nil {
			log.Error().Err(err).Msg("error parsing package.json")
			return
		}

		npmVersion = pac.Version
	}

	return
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ExtractPackage() {
	file, err := os.ReadFile("package.json")
	if err != nil {
		log.Error().Err(err).Msg("error reading package.json")
		return
	}

	pac := Package{}

	err = json.Unmarshal(file, &pac)
	if err != nil {
		log.Error().Err(err).Msg("error parsing package.json")
		return
	}

	startScript := pac.Scripts["start:web"]

	if strings.Contains(startScript, "next") && !strings.Contains(startScript, "-p") {
		startScript = fmt.Sprintf("%s -p %d", startScript, Config.ProxyPort)
	}

	startCommand = startScript
}
