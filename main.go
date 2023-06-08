package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type scene struct {
	t             int
	idx           int
	framePerImage int
	img           []*ebiten.Image
	width, height int
}

func (s *scene) Update() error {
	s.t++
	if s.t%s.framePerImage == 0 {
		s.idx++
		s.idx = (s.idx + 1) % len(s.img)
	}
	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img[s.idx], &ebiten.DrawImageOptions{})
}

func (s *scene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	scene := &scene{
		framePerImage: 10,
		width:         1,
		height:        1,
	}
	for _, arg := range os.Args[1:] {
		fp, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(fp)
		if err != nil {
			log.Fatal(err)
		}
		eimg := ebiten.NewImageFromImage(img)
		scene.img = append(scene.img, eimg)
		scene.width = max(scene.width, eimg.Bounds().Dx())
		scene.height = max(scene.height, eimg.Bounds().Dy())
	}
	ebiten.SetWindowSize(scene.width, scene.height)
	ebiten.SetWindowTitle("eanim")

	if err := ebiten.RunGame(scene); err != nil {
		log.Fatal(err)
	}
}
