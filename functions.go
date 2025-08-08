package main

import "path/filepath"

func hasExtension(filename string) bool {
	return filepath.Ext(filename) != ""
}
