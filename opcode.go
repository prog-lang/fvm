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

var OPS = [EXIT + 1]Action{
	{Name: "PUSH", Exec: func(m *Machine) { m.DS.Push(m.OPR) }},
	{Name: "DROP", Exec: func(m *Machine) { m.DS.Drop() }},
	{Name: "COND", Exec: func(m *Machine) { m.BCR = m.DS.Pop() != 0 }},

	{Name: "STORE", Exec: func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range Int32AsBytes(m.DS.Peek()) {
			m.RAM[index+offset] = bite
		}
	}},
	{Name: "MOVE", Exec: func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range Int32AsBytes(m.DS.Pop()) {
			m.RAM[index+offset] = bite
		}
	}},
	{Name: "LOAD", Exec: func(m *Machine) {
		bytes := make([]byte, 4)
		index := int(m.OPR)
		for offset := 0; offset < 4; offset++ {
			bytes[offset] = m.RAM[index+offset]
		}
		m.DS.Push(BytesToInt32(bytes))
	}},

	{Name: "ADD", Exec: func(m *Machine) { m.DS.Push(m.DS.Pop() + m.DS.Pop()) }},
	{Name: "SUB", Exec: func(m *Machine) { m.DS.Push(-(m.DS.Pop() - m.DS.Pop())) }},
	{Name: "MUL", Exec: func(m *Machine) { m.DS.Push(m.DS.Pop() * m.DS.Pop()) }},
	{Name: "DIV", Exec: func(m *Machine) {
		a := m.DS.Pop()
		b := m.DS.Pop()
		m.DS.Push(a / b)
	}},

	{Name: "LT", Exec: func(m *Machine) {
		m.DS.Push(BoolToInt32(m.DS.Pop() > m.DS.Pop()))
	}},
	{Name: "EQ", Exec: func(m *Machine) {
		m.DS.Push(BoolToInt32(m.DS.Pop() == m.DS.Pop()))
	}},
	{Name: "GT", Exec: func(m *Machine) {
		m.DS.Push(BoolToInt32(m.DS.Pop() < m.DS.Pop()))
	}},

	{Name: "JUMP", Exec: func(m *Machine) { m.IP = m.OPR }},
	{Name: "CALL", Exec: func(m *Machine) { m.CS.Push(m.IP); m.IP = m.OPR }},
	{Name: "BR", Exec: func(m *Machine) {
		if m.Condition() {
			m.IP = m.OPR
		}
	}},
	{Name: "DONE", Exec: func(m *Machine) { m.IP = m.CS.Pop() }},
	{Name: "EXIT", Exec: func(m *Machine) { m.OK = false }},
}

type Action struct {
	Name string
	Exec func(*Machine)
}

func (action Action) String() string {
	return action.Name
}
