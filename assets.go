package main

import (
	"embed"
	"io/fs"
)

//go:embed static
var bundledAssets embed.FS

func mustStaticFS() fs.FS {
	files, err := fs.Sub(bundledAssets, "static")
	if err != nil {
		panic(err)
	}
	return files
}
