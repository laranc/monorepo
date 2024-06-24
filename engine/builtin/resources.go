package builtin

import "reflect"

type GameState struct {
	state       int
	reflectType reflect.Type
}

func MakeGameState(defaultState int) GameState {
	gs := GameState{
		state: defaultState,
	}
	gs.reflectType = reflect.TypeOf(gs)
	return gs
}

func (gs *GameState) Type() reflect.Type {
	return gs.reflectType
}
