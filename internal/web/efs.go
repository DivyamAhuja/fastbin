package web

import (
	"embed"
	"io/fs"
)

//go:embed assets
var f embed.FS

var Files fs.FS

func init() {
	Files, _ = fs.Sub(f, "assets")
}
