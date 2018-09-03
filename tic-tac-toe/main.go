package main

import (
	"fmt"
	"syscall/js"
)

func play(i []js.Value) {
	elements := js.Global().
		Get("document").
		Get("forms").
		Get("tictac").
		Get("elements")

	e := ticTac(elements)

	if hasWinner(e) {
		js.Global().Get("alert").Invoke(`We have a winner`)
	}
}

func ticTac(e js.Value) []string {
	v := make([]string, 9)

	for i := 0; i < 9; i++ {
		v[i] = e.Index(i).Get("value").String()
	}

	return v
}

func hasWinner(e []string) bool {
	if checkEmpty(e[0], e[1], e[2]) && e[0] == e[1] && e[1] == e[2] {
		return true
	}

	if checkEmpty(e[3], e[4], e[5]) && e[3] == e[4] && e[4] == e[5] {
		return true
	}

	if checkEmpty(e[6], e[7], e[8]) && e[6] == e[7] && e[7] == e[8] {
		return true
	}

	if checkEmpty(e[0], e[3], e[6]) && e[0] == e[3] && e[3] == e[6] {
		return true
	}

	if checkEmpty(e[1], e[4], e[7]) && e[1] == e[4] && e[4] == e[7] {
		return true
	}

	if checkEmpty(e[2], e[5], e[8]) && e[2] == e[5] && e[5] == e[8] {
		return true
	}

	if checkEmpty(e[0], e[4], e[8]) && e[0] == e[4] && e[4] == e[8] {
		return true
	}

	if checkEmpty(e[2], e[4], e[6]) && e[2] == e[4] && e[4] == e[6] {
		return true
	}

	return false
}

func checkEmpty(t ...string) bool {
	for _, v := range t {
		if v == `` {
			return false
		}
	}

	return true
}

func registerCallbacks() {
	js.Global().Set("play", js.NewCallback(play))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println(`Tic Tac Toe Initialized`)

	registerCallbacks()
	<-c
}
