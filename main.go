//go:generate pioasm -o go blink.pio blink_pio.go

package main

import (
	"fmt"
	"machine"
)

const clockHz = 133000000

func main() {
	pio := machine.PIO0

	pio.Configure()

	offset := pio.AddProgram(&blinkProgram)
	fmt.Printf("Loaded program at %d\n", offset)

	blinkPinForever(&pio.StateMachines[0], offset, machine.LED, 3)
	blinkPinForever(&pio.StateMachines[1], offset, machine.GPIO6, 4)
	blinkPinForever(&pio.StateMachines[2], offset, machine.GPIO11, 1)
}

func blinkPinForever(sm *machine.PIOStateMachine, offset uint8, pin machine.Pin, freq uint) {
	blinkProgramInit(sm, offset, pin)
	sm.SetEnabled(true)

	fmt.Printf("Blinking pin %d at %d Hz\n", pin, freq)
	sm.Tx(uint32(clockHz / (2 * freq)))
}
