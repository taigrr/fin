package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed img/background.jpg
var backgroundJpg []byte

var resourceBackgroundJpg = &fyne.StaticResource{
	StaticName:    "background.jpg",
	StaticContent: backgroundJpg,
}
