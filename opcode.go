package machine

import (
	"machine/opcode"
)

// IS is INSTRUCTION SET, it contains string names and actions for every opcode
// declared in the enum above.
var IS = [opcode.Count]Action{
	{"PUSH", func(m *Machine) { m.DS.Push(Int32AsBytes(m.OPR)...) }},
	{"DROP", func(m *Machine) { m.DS.Drop(int(m.OPR)) }},

	{"STORE", func(m *Machine) {
		index := int(m.OPR)
		for offset, bite := range m.DS.Glance(Int32Size) {
			m.RAM[index+offset] = bite
		}
	}},
	{"LOAD", func(m *Machine) {
		m.DS.Push(m.RAM[m.OPR : m.OPR+Int32Size]...)
	}},

	{"JUMP", func(m *Machine) { m.IP = m.OPR }},
	{"CALL", func(m *Machine) { m.Call() }},
	{"BR", func(m *Machine) {
		if ByteAsBool(m.DS.Pop()) {
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
