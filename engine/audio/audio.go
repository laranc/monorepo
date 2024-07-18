package audio

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func Init(rate int, format uint16, channels, chunkSize int) {
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		fmt.Println(err)
		return
	}
	mix.OpenAudio(rate, format, channels, chunkSize)
}

func LoadSound(path string) *mix.Chunk {
	chunk, err := mix.LoadWAV(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return chunk
}

func LoadMusic(path string) *mix.Music {
	music, err := mix.LoadMUS(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return music
}

func PlaySound(sound *mix.Chunk) {
	sound.Play(-1, 0)
}

func PlayMusic(music *mix.Music) {
	music.Play(-1)
}
