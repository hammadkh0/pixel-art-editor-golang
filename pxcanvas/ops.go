package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	default:
		pxCanvas.PxSize = 10
	}
}

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := previousCoord.Position.X - currentCoord.Position.X
	yDiff := previousCoord.Position.Y - currentCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}
