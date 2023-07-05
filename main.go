package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type scene struct {
	t             int
	idx           int
	isPlaying     bool
	framePerImage int
	img           []*ebiten.Image
	width, height int
}

func (s *scene) Update() error {
	s.t++

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		s.isPlaying = !s.isPlaying
	case isKeyLongPressed(ebiten.KeyArrowLeft):
		s.idx = (s.idx - 1 + len(s.img)) % len(s.img)
	case isKeyLongPressed(ebiten.KeyArrowRight):
		s.idx = (s.idx + 1) % len(s.img)
	case isKeyLongPressed(ebiten.KeyArrowUp):
		s.framePerImage = max(1, s.framePerImage-1)
	case isKeyLongPressed(ebiten.KeyArrowDown):
		s.framePerImage++
	}

	if s.isPlaying && s.t%s.framePerImage == 0 {
		s.idx = (s.idx + 1) % len(s.img)
	}

	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img[s.idx], &ebiten.DrawImageOptions{})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%03d/%03d, speed 60/%2d=%.1f [img/s]", s.idx, len(s.img), s.framePerImage, 60/float64(s.framePerImage)))
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

func listImgFiles(dir string) []string {
	var files []string
	dirs, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, e := range dirs {
		if name := e.Name(); !e.IsDir() && strings.HasSuffix(name, ".png") {
			files = append(files, filepath.Join(dir, name))
		}
	}
	return files
}

func main() {
	var files []string

	switch {
	case len(os.Args) == 1:
		files = listImgFiles(".")
	case len(os.Args) == 2:
		stat, err := os.Stat(os.Args[1])
		if err != nil || !stat.IsDir() {
			println(err)
			files = os.Args[1:]
			break
		}
		files = listImgFiles(os.Args[1])
	default:
		files = os.Args[1:]
	}

	if files == nil {
		println("not found img files")
		println("Usage: eanim [dir or files]")
		os.Exit(1)
	}

	scene := &scene{
		framePerImage: 10,
		isPlaying:     true,
	}
	for _, arg := range files {
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
