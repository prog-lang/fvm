package machine

const (
	/* Stack manipulation */
	PUSH int32 = iota // PUSH int32 on the stack
	DROP              // DROP top value on the stack
	COND              // COND moves top stack value into the BCR

	/* RAM manipulation */
	STORE // STORE top stack value at some location in RAM
	MOVE  // MOVE top stack value to some location in RAM
	LOAD  // LOAD value from RAM onto a stack

	/* Arithmetic operations */

	ADD // ADD top two numbers on the Stack
	SUB // SUB top two numbers on the Stack
	MUL // MUL top two numbers on the Stack
	DIV // DIV top two numbers on the Stack

	/* Comparisons */

	LT // LT is LESS THAN
	EQ // EQ is EQUAL
	GT // GT is GREATER THAN

	/* Program flow */

	JUMP // JUMP IP to the specified instruction address in ROM
	CALL // CALL pushed return address onto the call stack and then jumps
	BR   // BR will perform a CALL if BCR is true
	DONE // DONE jumps back to the calling subroutine
	EXIT // EXIT the program
	// EXIT must always remain the last instruction in the set
)

var OPS = [EXIT + 1]Operation{
	{Name: "PUSH", Exec: func(m *Machine) { m.Push(m.ODR) }},
	{Name: "DROP", Exec: func(m *Machine) { m.Drop() }},
	{Name: "COND", Exec: func(m *Machine) { m.BCR = m.Pop() != 0 }},

	{Name: "STORE", Exec: func(m *Machine) {
		index := int(m.ODR)
		for offset, bite := range Int32AsBytes(m.Peek()) {
			m.RAM[index+offset] = bite
		}
	}},
	{Name: "MOVE", Exec: func(m *Machine) {
		index := int(m.ODR)
		for offset, bite := range Int32AsBytes(m.Pop()) {
			m.RAM[index+offset] = bite
		}
	}},
	{Name: "LOAD", Exec: func(m *Machine) {
		bytes := make([]byte, 4)
		index := int(m.ODR)
		for offset := 0; offset < 4; offset++ {
			bytes[offset] = m.RAM[index+offset]
		}
		m.Push(BytesToInt32(bytes))
	}},

	{Name: "ADD", Exec: func(m *Machine) { m.Push(m.Pop() + m.Pop()) }},
	{Name: "SUB", Exec: func(m *Machine) { m.Push(-(m.Pop() - m.Pop())) }},
	{Name: "MUL", Exec: func(m *Machine) { m.Push(m.Pop() * m.Pop()) }},
	{Name: "DIV", Exec: func(m *Machine) {
		a := m.Pop()
		b := m.Pop()
		m.Push(a / b)
	}},

	{Name: "LT", Exec: func(m *Machine) {
		m.Push(BoolToInt32(m.Pop() > m.Pop()))
	}},
	{Name: "EQ", Exec: func(m *Machine) {
		m.Push(BoolToInt32(m.Pop() == m.Pop()))
	}},
	{Name: "GT", Exec: func(m *Machine) {
		m.Push(BoolToInt32(m.Pop() < m.Pop()))
	}},

	{Name: "JUMP", Exec: func(m *Machine) { m.IP = m.ODR }},
	{Name: "CALL", Exec: func(m *Machine) { m.Call.Push(m.IP); m.IP = m.ODR }},
	{Name: "BR", Exec: func(m *Machine) {
		if m.Condition() {
			m.IP = m.ODR
		}
	}},
	{Name: "DONE", Exec: func(m *Machine) { m.IP = m.Call.Pop() }},
	{Name: "EXIT", Exec: func(m *Machine) { m.OK = false }},
}

type Operation struct {
	Name string
	Exec func(*Machine)
}

func (opcode Operation) String() string {
	return opcode.Name
}
