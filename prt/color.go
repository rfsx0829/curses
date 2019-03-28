package prt

type Mode uint32
type Color uint32

const (
	Default Mode = iota
	HighLight
	Light
	Italic
	UnderLine
	Flashing
	Normal
	Invert
	Invisible
)

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Fuchsia
	Azure
	White
)
