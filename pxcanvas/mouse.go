package pxcanvas

import (
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
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent // update the coordinate after the panning is done so this value can be used in next panning
}
