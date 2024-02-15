package machine

const (
	/* Stack manipulation */

	PUSH int32 = iota // PUSH int32 on the stack
	DROP              // DROP top value on the stack

	/* RAM manipulation */
	STORE // STORE top stack value at some location in RAM
	LOAD  // LOAD value from RAM onto a stack

	/* Arithmetic operations */

	ADD // ADD top two numbers on the Stack
	SUB // SUB top two numbers on the Stack
	MUL // MUL top two numbers on the Stack
	DIV // DIV top two numbers on the Stack

	/* Conversions */
	I2B // I2B converts int32 -> byte
	B2I // B2I converts byte -> int32

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
	{"PUSH", func(m *Machine) { m.DS.Push(Int32AsBytes(m.OPR)...) }},
	{"DROP", func(m *Machine) { m.DS.Drop(int(m.OPR)) }},

	{"STORE", func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range m.DS.Glance(int32Size) {
			m.RAM[index+offset] = bite
		}
	}},
	{"LOAD", func(m *Machine) {
		m.DS.Push(m.RAM[m.OPR : m.OPR+4]...)
	}},

	{"ADD", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(Int32AsBytes(x + y)...)
	}},
	{"SUB", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(Int32AsBytes(x - y)...)
	}},
	{"MUL", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(Int32AsBytes(x * y)...)
	}},
	{"DIV", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(Int32AsBytes(x / y)...)
	}},

	{"I2B", func(m *Machine) {
		m.DS.Push(BoolAsByte(Int32AsBool(BytesAsInt32(m.DS.Take(int32Size)))))
	}},
	{"B2I", func(m *Machine) { m.DS.Push(Int32AsBytes(int32(m.DS.Pop()))...) }},

	{"LT", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(BoolAsByte(x > y))
	}},
	{"EQ", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(BoolAsByte(x == y))
	}},
	{"GT", func(m *Machine) {
		x, y := pop2Int32(m)
		m.DS.Push(BoolAsByte(x < y))
	}},

	{"JUMP", func(m *Machine) { m.IP = m.OPR }},
	{"CALL", call},
	{"BR", func(m *Machine) {
		if ByteAsBool(m.DS.Pop()) {
			call(m)
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

/* ACTION HELPERS */

func pop2Int32(m *Machine) (x, y int32) {
	y = BytesAsInt32(m.DS.Take(int32Size))
	x = BytesAsInt32(m.DS.Take(int32Size))
	return
}

func call(m *Machine) {
	m.CS.Push(m.IP)
	m.IP = m.OPR
}
