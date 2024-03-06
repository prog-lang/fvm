package machine

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsyncCmd(t *testing.T) {
	changed := false
	change := Cmd(func() Object {
		time.Sleep(1 * time.Minute)
		changed = true
		return Unit{}
	})
	unit := async_Cmd.Feed(change).(Cmd).Exec()
	assert.Equal(t, Unit{}, unit)
	assert.False(t, changed)
}

func TestMapCmd(t *testing.T) {
	result := map_Cmd.
		Feed(show_I32).(Function).
		Feed(cmd_Cmd.Feed(int32(42))).(Cmd).
		Exec()
	assert.Equal(t, "42", result)
}

func TestSwapCmd(t *testing.T) {
	changed := false
	change := Cmd(func() Object {
		changed = true
		return 42
	})
	result := swap_Cmd.
		Feed(42).(Function).
		Feed(change).(Cmd).
		Exec()
	assert.True(t, changed)
	assert.Equal(t, 42, result)
}

func TestThenCmd(t *testing.T) {
	changed := false
	change := Cmd(func() Object {
		changed = true
		return 666
	})
	number := cmd_Cmd.Feed(42).(Cmd)
	result := then_Cmd.
		Feed(change).(Function).
		Feed(number).(Cmd).
		Exec()
	assert.Equal(t, 42, result)
	assert.True(t, changed)
}

func TestChainCmd(t *testing.T) {
	changed := false
	change := Cmd(func() Object {
		changed = true
		return 666
	})
	increment := MakeFn(1, func(args []Object) Object {
		return args[0].(int) + 1
	})
	result := chain_Cmd.
		Feed(change).(Function).
		Feed(increment).(Cmd).
		Exec()
	assert.Equal(t, 667, result)
	assert.True(t, changed)
}
