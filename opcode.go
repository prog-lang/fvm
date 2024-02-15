package machine

const (
	/* Stack manipulation */

	PUSH int32 = iota // PUSH int32 on the stack
	DROP              // DROP top value on the stack

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
	CALL // CALL pushes return address onto CS and then jumps
	BR   // BR will perform a CALL if top DS value is true
	DONE // DONE jumps back to the calling subroutine
	EXIT // EXIT the program
	// EXIT must always remain the last instruction in the set
)

// IS is INSTRUCTION SET, it contains string names and actions for every opcode
// declared in the enum above.
var IS = [EXIT + 1]Action{
	{"PUSH", func(m *Machine) { m.DS.Push(m.OPR) }},
	{"DROP", func(m *Machine) { m.DS.Drop() }},

	{"STORE", func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range Int32AsBytes(m.DS.Peek()) {
			m.RAM[index+offset] = bite
		}
	}},
	{"MOVE", func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range Int32AsBytes(m.DS.Pop()) {
			m.RAM[index+offset] = bite
		}
	}},
	{"LOAD", func(m *Machine) {
		m.DS.Push(BytesAsInt32(m.RAM[m.OPR : m.OPR+4]))
	}},

	{"ADD", func(m *Machine) {
		m.DS.Push(m.DS.Pop() + m.DS.Pop())
	}},
	{"SUB", func(m *Machine) {
		m.DS.Push(-(m.DS.Pop() - m.DS.Pop()))
	}},
	{"MUL", func(m *Machine) {
		m.DS.Push(m.DS.Pop() * m.DS.Pop())
	}},
	{"DIV", func(m *Machine) {
		a := m.DS.Pop()
		b := m.DS.Pop()
		m.DS.Push(a / b)
	}},

	{"LT", func(m *Machine) {
		m.DS.Push(BoolAsInt32(m.DS.Pop() > m.DS.Pop()))
	}},
	{"EQ", func(m *Machine) {
		m.DS.Push(BoolAsInt32(m.DS.Pop() == m.DS.Pop()))
	}},
	{"GT", func(m *Machine) {
		m.DS.Push(BoolAsInt32(m.DS.Pop() < m.DS.Pop()))
	}},

	{"JUMP", func(m *Machine) { m.IP = m.OPR }},
	{"CALL", func(m *Machine) { m.Call() }},
	{"BR", func(m *Machine) {
		if Int32AsBool(m.DS.Pop()) {
			m.Call()
		}
	}},
	{"DONE", func(m *Machine) { m.IP = m.CS.Pop() }},
	{"EXIT", func(m *Machine) { m.OK = false }},
}

type Action struct {
	Name string
	Exec func(*Machine)
}

func (action Action) String() string {
	return action.Name
}
