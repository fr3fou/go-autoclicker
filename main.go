package main

import (
	"github.com/JamesHovious/w32"
)

func main() {
	i := w32.INPUT{
		Type: 0,
		Mi: w32.MOUSEINPUT{
			DwFlags: 0x0002,
		},
	}
	w32.SendInput([]w32.INPUT{i})
}
