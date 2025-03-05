package main

import (
	"embed"
	"github.com/liushiqi1001/go-whatsapp-web-multidevice/cmd"
)

//go:embed views/index.html
var embedIndex embed.FS

//go:embed views
var embedViews embed.FS

func main() {
	cmd.Execute(embedIndex, embedViews)
}
