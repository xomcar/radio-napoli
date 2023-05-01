package main

import (
	emb "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	fig "github.com/common-nighthawk/go-figure"
	"github.com/eiannone/keyboard"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	//go:embed data/*.mp3
	fs     emb.FS
	tracks []*track
	key    = make(chan rune)
)

type track struct {
	name   string
	stream beep.StreamSeekCloser
	format beep.Format
}

func getTrack(fs emb.FS, filename string) (*track, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return &track{}, err
	}
	s, f, err := mp3.Decode(file)
	if err != nil {
		return &track{}, err
	}
	name := strings.TrimSuffix(strings.TrimPrefix(filename, "data/"), ".mp3")
	return &track{stream: s, format: f, name: name}, nil
}
func main() {
	fig.NewColorFigure("Radio Napoli", "", "cyan", true).Scroll(3500, 300, "left")
	files, _ := fs.ReadDir("data")
	for _, file := range files {
		t, err := getTrack(fs, "data/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		tracks = append(tracks, t)
		defer t.stream.Close()
	}
	if len(tracks) == 0 {
		log.Fatal(errors.New("no tracks found"))
	}
	format := tracks[0].format
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	fmt.Println("q to quit, s to stop")
	for i, track := range tracks {
		fmt.Printf("[%2d] %v\n", i, track.name)
	}
	for {
		go func() {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				log.Fatal(err)
			}
			key <- char
		}()
		char := <-key
		if unicode.IsDigit(char) && int(char-'0') < len(tracks) {
			speaker.Clear()
			for _, t := range tracks {
				t.stream.Seek(0)
			}
			speaker.Play(tracks[int(char-'0')].stream)
		}
		switch unicode.ToLower(char) {
		case 'q':
			os.Exit(0)
		case 's':
			speaker.Clear()
			for _, t := range tracks {
				t.stream.Seek(0)
			}
		}
		if unicode.ToLower(char) == 'q' {
			os.Exit(0)
		}
	}
}
