package resources

import (
	_ "embed"
)

var (
	//go:embed bgm/relaxed-piano-bgm-242462.mp3
	BGMFile []byte

	//go:embed tilesets/grass.png
	GrassTextureFile []byte

	//go:embed characters/basic-charakter-spritesheet.png
	PlayerTextureFile []byte
)
