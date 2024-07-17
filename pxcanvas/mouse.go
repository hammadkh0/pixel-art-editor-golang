package pxcanvas

import (
	"pixel-art-editor-golang/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY)) //float value for direct in which wheel is scrolled
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0)) //hide the cursor when mouse is moved outside the canvas
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent // update the coordinate after the panning is done so this value can be used in next panning
}

// Mouseable Interface implementation
func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
