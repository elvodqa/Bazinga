package graphics

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ImageRenderer struct {
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
}

func NewImageRenderer(renderer *sdl.Renderer, imagePath string) (*ImageRenderer, error) {
	// load png
	image, err := img.Load(imagePath)
	if err != nil {
		return nil, err
	}
	defer image.Free()

	// create texture
	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, err
	}

	return &ImageRenderer{
		Renderer: renderer,
		Texture:  texture,
	}, nil
}

func (i *ImageRenderer) Render(x, y int32) error {
	_, _, w, h, err := i.Texture.Query()
	if err != nil {
		return err
	}

	rect := sdl.Rect{x, y, w, h}
	return i.Renderer.Copy(i.Texture, nil, &rect)
}

func (i *ImageRenderer) RenderWithSize(x, y, w, h int32) error {
	rect := sdl.Rect{x, y, w, h}
	return i.Renderer.Copy(i.Texture, nil, &rect)
}

func (i *ImageRenderer) RenderWithRect(x, y int32, rect *sdl.Rect) error {
	return i.Renderer.Copy(i.Texture, rect, &sdl.Rect{x, y, rect.W, rect.H})
}

func (i *ImageRenderer) SetImage(path string) error {
	// load png
	image, err := img.Load(path)
	if err != nil {
		return err
	}
	defer image.Free()

	// create texture
	texture, err := i.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		return err
	}

	i.Texture = texture
	return nil
}
