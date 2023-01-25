package core

import "github.com/veandco/go-sdl2/sdl"

type GameTime struct {
	now       uint64
	last      uint64
	DeltaTime float64
}

func NewGameTime() *GameTime {
	return &GameTime{
		now:       sdl.GetPerformanceCounter(),
		last:      0,
		DeltaTime: 0,
	}
}

func (g *GameTime) Update() {
	g.last = g.now
	g.now = sdl.GetPerformanceCounter()
	g.DeltaTime = float64(g.now-g.last) / float64(sdl.GetPerformanceFrequency())
}

func (g *GameTime) GetTime() float64 {
	return float64(g.now) / float64(sdl.GetPerformanceFrequency())
}

func (g *GameTime) GetTimeAsSeconds() float64 {
	return g.GetTime() / 1000
}

func (g *GameTime) GetFPS() float64 {
	return 1 / (float64(g.now-g.last) / 1000)
}
