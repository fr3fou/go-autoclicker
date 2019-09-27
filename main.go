package main

import (
	"time"

	"github.com/JamesHovious/w32"
)

func main() {
	time.Sleep(time.Second * 5)
	down := w32.INPUT{
		Type: 0,
		Mi: w32.MOUSEINPUT{
			DwFlags: w32.MOUSEEVENTF_LEFTDOWN,
		},
	}

	up := w32.INPUT{
		Type: 0,
		Mi: w32.MOUSEINPUT{
			DwFlags: w32.MOUSEEVENTF_LEFTUP,
		},
	}

	for {
		w32.SendInput([]w32.INPUT{down})
		time.Sleep(25)
		w32.SendInput([]w32.INPUT{up})
	}
}
