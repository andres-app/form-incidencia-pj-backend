package helper

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// Extrae los datos de imagen de la cadena base64
func ExtractImageData(base64Str string) []byte {
	data := strings.Split(base64Str, ",")
	if len(data) != 2 {
		return nil
	}

	imageData, err := base64.StdEncoding.DecodeString(data[1])
	if err != nil {
		return nil
	}

	return imageData
}

// Extrae el tipo de imagen desde la cadena base64
func ExtractImageType(base64Str string) string {
	data := strings.Split(base64Str, ";")
	if len(data) != 2 {
		return ""
	}

	mimeTypeData := strings.Split(data[0], ":")
	if len(mimeTypeData) != 2 {
		return ""
	}

	return strings.Split(mimeTypeData[1], "/")[1]
}

// Guarda los datos de imagen en un archivo

func SaveImage(imageData []byte, imageType, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image
	switch imageType {
	case "png":
		img, err = png.Decode(strings.NewReader(string(imageData)))
	case "jpg", "jpeg":
		img, err = jpeg.Decode(strings.NewReader(string(imageData)))
	default:
		return fmt.Errorf("El formato no se pudo decodificar.")
	}

	if err != nil {
		return err
	}

	switch imageType {
	case "png":
		err = png.Encode(file, img)
	case "jpg", "jpeg":
		err = jpeg.Encode(file, img, nil)
	default:
		return fmt.Errorf("El formato no se pudo guardar.")
	}

	if err != nil {
		return err
	}

	return nil
}

func NumbertToString(number int) string {
	if number < 9 {
		return "0" + strconv.Itoa(number)
	}

	return strconv.Itoa(number)
}
