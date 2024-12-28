package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"regexp"

	"golang.org/x/image/draw"
)

type Util struct {
}

// NewUtil returns a new utility instance and exposes the methods
func NewUtil() *Util {
	return &Util{}
}

// CreateJSONEnvelope assigns a key to a new map and fills it with data
func (u *Util) CreateJSONEnvelope(key string, data any) map[string]any {
	if len(key) < 1 {
		key = "data"
	}

	if data == nil {
		data = make(map[string]any)
	}
	var envelope = make(map[string]any)

	envelope[key] = data

	return envelope
}

func (u *Util) ResizeImage(rdr io.Reader, wtr io.Writer, mimetype string, width int) (err error) {
	var src image.Image

	switch mimetype {
	case "image/jpeg":
		src, err = jpeg.Decode(rdr)
	case "image/png":
		src, err = png.Decode(rdr)
	}

	if err != nil {
		return err
	}

	ratio := (float64)(src.Bounds().Max.Y) / (float64)(src.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	err = jpeg.Encode(wtr, dst, nil)
	if err != nil {
		return err
	}

	return nil

}

// CheckPWStrength does some exremely basic pass word strenth checks and returns a bool if it passes or not
func (u *Util) CheckPWStrength(pw string, minLength int) bool {

	if len(pw) < minLength {
		return false
	}

	charRe, err := regexp.Compile("[^a-zA-Z0-9\\n]")

	if err != nil {
		return false
	}

	hasSpecialChar := charRe.MatchString(pw)

	if !hasSpecialChar {
		return false
	}

	alphaRe, err := regexp.Compile("[a-zA-Z]")

	if err != nil {
		return false
	}

	hasAlpha := alphaRe.MatchString(pw)

	if !hasAlpha {
		return false
	}

	numRe, err := regexp.Compile("[0-9]")

	if err != nil {
		return false
	}

	hasNum := numRe.MatchString(pw)

	if !hasNum {
		return false
	}

	return true

}
