package utils

import (
	"bytes"
	"encoding/base64"

	"github.com/disintegration/imaging"
)

func ConvertPngToDataUri(path string) (string, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return "", err
	}

	var imageBuffer bytes.Buffer
	if err := imaging.Encode(&imageBuffer, img, imaging.PNG); err != nil {
		return "", err
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageBuffer.Bytes()), nil
}
