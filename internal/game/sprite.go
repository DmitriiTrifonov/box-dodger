package game

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Sprite struct {
	File       *goaseprite.File
	AnimPlayer *goaseprite.Player
	Img        *ebiten.Image
}

func NewSprite(path string) (*Sprite, error) {
	dirFS := os.DirFS("./bin")
	file, err := goaseprite.Open(path, dirFS)
	if err != nil {
		return nil, err
	}

	img, _, err := ebitenutil.NewImageFromFile(file.ImagePath)
	if err != nil {
		return nil, err
	}

	animPlayer := file.CreatePlayer()

	return &Sprite{
		File:       file,
		AnimPlayer: animPlayer,
		Img:        img,
	}, nil
}

func (s *Sprite) SetAnimTag(tag string) error {
	return s.AnimPlayer.Play(tag)
}

func (s *Sprite) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	sub := s.Img.SubImage(image.Rect(s.AnimPlayer.CurrentFrameCoords()))
	screen.DrawImage(sub.(*ebiten.Image), opts)
}

func (s *Sprite) SetFrameIdx(idx int) {
	s.AnimPlayer.SetFrameIndexInAnimation(idx)
}

func (s *Sprite) Clone() *Sprite {
	return &Sprite{
		File:       s.File,
		AnimPlayer: s.AnimPlayer.Clone(),
		Img:        s.Img,
	}
}
