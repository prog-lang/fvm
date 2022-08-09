package machine

import (
	"fmt"
	"log"
)

type Machine struct {
	/* COMPILED CODE */

	// Data contains constants from a program.
	Data []byte

	// ROM is READ ONLY MEMORY that contains program instructions.
	//
	// Let's assume that each instruction is 64 bits long with 32 bits opcode
	// and another 32 bits for operand.
	//
	//        #0     #1     #2     #3     #4     #5     #6     #7
	//     *------*------*------*------*------*------*------*------*
	//     | byte |      |      |      |      |      |      |      |
	//     *------*------*------*------*------*------*------*------*
	//     <--------- opcode ---------> <-------- operand --------->
	//
	// Opcode: operation code
	// Operand: operation argument
	//
	// Example sequence of instructions:
	//
	//     // ASM  vs BYTECODE
	//     PUSH 40 // 0x00 0x00 0x00 0x00  0x00 0x00 0x00 0x28
	//     PUSH 2  // 0x00 0x00 0x00 0x00  0x00 0x00 0x00 0x02
	//     ADD     // 0x00 0x00 0x00 0x01  0x00 0x00 0x00 0x00
	//
	// The way each opcode is encoded is decided by the enum in the opcode.go
	// file.
	ROM []int32

	/* MUTABLE MEMORY */

	// RAM is RANDOM ACCESS MEMORY. Here, our program may decide to store large
	// amounts of data if need be. As it's defined now, RAM's capacity is 1MB.
	RAM [1000000]byte

	// DS is DATA STACK, it allows function calls and arithmetic to be possible.
	DS *Stack

	// CS is CALL STACK that stores return addresses.
	CS *Stack

	/* SPECIALISED REGISTERS */

	// OK boolean flag displays machine's health status. If OK == false,
	// execution must be stopped.
	OK bool

	// BCR is BOOLEAN CONDITION REGISTER used by the BR instruction to decide
	// wether to perform a branching call.
	BCR bool

	// IP is INSTRUCTION POINTER that points to an index in ROM from which our
	// machine is supposed to fetch the next instruction.
	IP int32

	// OCR is OPCODE REGISTER. Fetch will put the next instruction's opcode into
	// this register.
	OCR int32

	// OPR is OPERAND REGISTER. Fetch will put the next instruction's operand
	// into this register.
	OPR int32

	// AR is ACTION REGISTER. Decode will put appropriate operation into
	// this register based on the opcode we've fetched during the Fetch stage.
	AR Action
}

func New(data []byte, rom []int32) *Machine {
	return &Machine{
		Data: data,
		ROM:  rom,
		OK:   true,
		DS:   NewStack(),
		CS:   NewStack(),
	}
}

func (m *Machine) Run() {
	for m.OK {
		m.Cycle()
	}
}

func (m *Machine) Cycle() {
	m.Fetch()
	m.Decode()
	m.Execute()
	log.Print(m)
}

func (m *Machine) Fetch() {
	m.OCR = m.ROM[m.IP]
	m.IP++
	m.OPR = m.ROM[m.IP]
	m.IP++
}

func (m *Machine) Decode() {
	m.AR = OPS[m.OCR]
}

func (m *Machine) Execute() {
	m.AR.Exec(m)
}

// Condition must always be checked through this method as BCR flag must be
// reset to false after use.
func (m *Machine) Condition() (cond bool) {
	cond = m.BCR
	m.BCR = false
	return
}

func (m *Machine) String() string {
	return fmt.Sprintf("%-10s %-10d; %c %-20s #%v",
		m.AR, m.OPR, BoolToEmoji(m.BCR), m.DS, m.RAM[:15])
}
