package main

import (
	"time"

	"github.com/JamesHovious/w32"
)

func main() {
	down := w32.INPUT{
		Type: 0,
		Mi: w32.MOUSEINPUT{
			DwFlags:   w32.MOUSEEVENTF_LEFTDOWN,
			MouseData: 0,
			Time:      0,
		},
	}

	up := w32.INPUT{
		Type: 0,
		Mi: w32.MOUSEINPUT{
			DwFlags:   w32.MOUSEEVENTF_LEFTUP,
			MouseData: 0,
			Time:      0,
		},
	}

	for {
		w32.SendInput([]w32.INPUT{down})
		time.Sleep(25)
		w32.SendInput([]w32.INPUT{up})
	}
}
