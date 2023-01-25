package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextRenderer struct {
	Font     *ttf.Font
	Renderer *sdl.Renderer
	Size     int
	Color    sdl.Color
	fontPath string
}

func NewTextRenderer(renderer *sdl.Renderer, fontPath string, size int, color sdl.Color) (*TextRenderer, error) {
	font, err := ttf.OpenFont(fontPath, size)
	if err != nil {
		return nil, err
	}

	return &TextRenderer{
		Font:     font,
		Renderer: renderer,
		Size:     size,
		Color:    color,
		fontPath: fontPath,
	}, nil
}

func (t *TextRenderer) Render(text string, x, y int32) error {
	surface, err := t.Font.RenderUTF8Blended(text, t.Color)
	if err != nil {
		return err
	}
	defer surface.Free()

	texture, err := t.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	_, _, w, h, err := texture.Query()
	if err != nil {
		return err
	}

	rect := sdl.Rect{x, y, w, h}
	return t.Renderer.Copy(texture, nil, &rect)
}

func (t *TextRenderer) SetColor(color sdl.Color) {
	t.Color = color
}

func (t *TextRenderer) SetColorRGB(r, g, b, a uint8) {
	t.Color = sdl.Color{r, g, b, a}
}

func (t *TextRenderer) SetSize(size int) error {
	t.Size = size
	font, err := ttf.OpenFont(t.fontPath, t.Size)
	if err != nil {
		return err
	}
	t.Font = font
	return nil
}

func (t *TextRenderer) SetFont(fontPath string) error {
	font, err := ttf.OpenFont(fontPath, t.Size)
	if err != nil {
		return err
	}
	t.Font = font
	return nil
}

func (t *TextRenderer) Close() {
	t.Font.Close()
}
