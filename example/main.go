package main

import (
	"fmt"
	"machine"
	. "machine/opcode"

	"github.com/charmbracelet/log"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	url := "https://google.com"
	data := []byte(url)
	code := []byte{
		I32_PUSH, 0, 1, 0, 0,
		DONE, 0, 0, 0, 0,
	}
	result := machine.NewRoutine(
		machine.NewROM(data),
		machine.NewROM(code),
		0,
	).Call()
	fmt.Println(result)
}
