package machine

import (
	"fmt"
	"machine/std"
)

var Std = [std.Count]Action{
	{"NOP", func(m *Machine) {}},

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
		m.DS.Push(BoolAsByte(Int32AsBool(BytesAsInt32(m.DS.Take(Int32Size)))))
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

	{"PRINT_DATA", func(m *Machine) {
		addr, length := pop2Int32(m)
		fmt.Print(string(m.Data[addr : addr+length+1]))
	}},
}

func pop2Int32(m *Machine) (x, y int32) {
	y = BytesAsInt32(m.DS.Take(Int32Size))
	x = BytesAsInt32(m.DS.Take(Int32Size))
	return
}
