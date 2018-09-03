package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	cb := js.NewCallback(func(args []js.Value) {
		move := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value")
		fmt.Println(move)
	})

	js.Global().Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)

	fmt.Println(`Event Listener Started`)

	<-c
}
