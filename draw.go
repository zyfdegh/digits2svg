package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ajstarks/svgo"
)

const (
	ascii0  = 48
	ascii1  = 49
	asciiLF = 10
)

var (
	// ErrInvalidData returned when data contains other charactors besides 0 and 1
	ErrInvalidData = errors.New("invalid data")
)

func read(file string) (content []byte, err error) {
	_, err = os.Stat(file)
	if err != nil {
		log.Printf("stat file error: %v\n", err)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		log.Printf("open file error: %v\n", err)
		return
	}
	defer f.Close()

	content, err = ioutil.ReadAll(f)
	if err != nil {
		log.Printf("read file error: %v\n", err)
		return
	}
	return
}

func parse(content []byte) (data []uint8, err error) {
	for _, c := range content {
		switch c {
		case ascii0:
			data = append(data, 0)
		case ascii1:
			data = append(data, 1)
		case asciiLF:
			continue
		default:
			err = ErrInvalidData
			return
		}
	}
	return
}

// drawing data
func draw(data []uint8, w io.Writer, yOffset int) (err error) {
	lenData := len(data)
	if lenData == 0 {
		return
	}

	canvas := svg.New(w)
	canvas.Start(50*lenData, 500+yOffset)

	// grids
	for i := 0; i < lenData; i++ {
		canvas.Grid(50*i, yOffset-50, 50, 50, 50, "fill:none;stroke:#eeeeee")
		canvas.Grid(50*i, yOffset, 50, 50, 50, "fill:none;stroke:#eeeeee")
		canvas.Grid(50*i, yOffset+50, 50, 50, 50, "fill:none;stroke:#eeeeee")
	}

	var previous uint8
	for i, d := range data {
		// vertical lines
		if i >= 1 && previous != d {
			canvas.Line(50*i, 0+yOffset, 50*i, 50+yOffset, "fill:none;stroke:black")
		}

		// horizental lines
		if d == 0 {
			canvas.Line(50*i, 50+yOffset, 50*(i+1), 50+yOffset, "fill:none;stroke:black")
		}
		if d == 1 {
			canvas.Line(50*i, 0+yOffset, 50*(i+1), 0+yOffset, "fill:none;stroke:black")
		}
		previous = d
	}
	canvas.End()
	return
}
