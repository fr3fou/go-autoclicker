package main

import (
	"fmt"
	"time"

	"github.com/JamesHovious/w32"
)

var kbHook w32.HHOOK

func main() {
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

	kbHook = w32.SetWindowsHookEx(w32.WH_KEYBOARD_LL, cb, 0, 0)

	w32.SendInput([]w32.INPUT{down})
	time.Sleep(25)
	w32.SendInput([]w32.INPUT{up})
}

func cb(nCode int, wp w32.WPARAM, lp w32.LPARAM) w32.LRESULT {
	if nCode >= 0 {
		fmt.Println("Here")
	}

	return w32.CallNextHookEx(kbHook, nCode, wp, lp)
}
