package main

import (
	"embed"

	"github.com/rr13k/pen/cmd"
)

//go:embed temps/*
var EmbedTempsContent embed.FS

func main() {
	cmd.Cli(EmbedTempsContent)
}
