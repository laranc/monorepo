package ai

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

type AIHandler struct {
	l       *lua.LState
	scripts []*AIScript
}

type AIScript struct {
	state string
	Table map[string]any
}

func MakeAIHandler() AIHandler {
	l := lua.NewState()
	l.OpenLibs()
	return AIHandler{l: l, scripts: make([]*AIScript, 1)}
}

func (state *AIHandler) Destroy() {
	state.l.Close()
}

func NewAIScript(initState string, entityID uint64) *AIScript {
	table := make(map[string]any)
	table["entity"] = entityID
	return &AIScript{state: initState, Table: table}
}

func (state *AIHandler) RegisterScript(script *AIScript, path string) error {
	if err := state.l.DoFile(path); err != nil {
		fmt.Printf("Error loading Lua script '%s': %e\n", script, err)
		return err
	}
	state.scripts = append(state.scripts, script)
	return nil
}

func (state *AIHandler) Update() {
	for _, script := range state.scripts {
		table := state.newTable(script.Table)
		err := state.l.CallByParam(lua.P{
			Fn:      state.l.GetGlobal(script.state),
			NRet:    1,
			Protect: true,
		}, table)
		if err != nil {
			fmt.Printf("Script error: %e\n", err)
			return
		}
		script.state = state.l.Get(-1).String()
		state.l.Pop(1)
	}
}

func (state *AIHandler) newTable(data map[string]any) *lua.LTable {
	table := state.l.NewTable()
	for key, value := range data {
		table.RawSetString(key, lua.LString(fmt.Sprintf("%v", value)))
	}
	return table
}
