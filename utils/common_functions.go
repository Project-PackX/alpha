package utils

import (
	"bytes"
	"encoding/base64"
	"math/rand"

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

// Generate a random string (letters + numbers) with the given length
func RandomString(length int) string {
	characters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}

// Generate a random string (numbers) with the given length
func RandomPackageCode(length int) string {
	characters := []byte("0123456789")
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}
