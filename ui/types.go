package ui

import (
	"pixel-art-editor-golang/apptype"
	"pixel-art-editor-golang/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
