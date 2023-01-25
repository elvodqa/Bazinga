package core

import "github.com/veandco/go-sdl2/sdl"

type Input struct {
	keys []uint8
}

func (t *Input) IsKeyPressed(key sdl.Scancode) bool {
	t.keys = sdl.GetKeyboardState()
	return t.keys[key] == 1
}
