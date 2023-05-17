package main

import (
	"embed"

	"github.com/rr13k/pen/cmd"
	"github.com/rr13k/pen/common"
)

//go:embed temps/*
var EmbedTempsContent embed.FS

func main() {
	common.EmbedTempsContent = EmbedTempsContent
	cmd.Cli()
}
