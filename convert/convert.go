package convert

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func FileToBinary(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	binaryData := make([]byte, len(data)*8)
	for i, byteValue := range data {
		for j := 0; j < 8; j++ {
			bit := (byteValue >> (7 - j)) & 1
			binaryData[i*8+j] = bit
		}
	}

	return binaryData, nil
}

func BinaryToImage(binaryData []byte, outputPath string) error {
	dataLen := len(binaryData)
	imageSize := int(math.Ceil(math.Sqrt(float64(dataLen))))

	img := image.NewGray(image.Rect(0, 0, imageSize, imageSize))

	index := 0
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			if index < dataLen && binaryData[index] == 1 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
			index++
		}
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

func ImageToBinary(imagePath string) ([]byte, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	binaryData := make([]byte, width*height)
	index := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			avg := (r + g + b) / 3
			if avg > 32768 {
				binaryData[index] = 1
			} else {
				binaryData[index] = 0
			}
			index++
		}
	}

	return binaryData, nil
}

func BinaryToFile(binaryData []byte, outputPath string) error {
	dataLen := len(binaryData) / 8
	fileData := make([]byte, dataLen)

	for i := 0; i < dataLen; i++ {
		for j := 0; j < 8; j++ {
			fileData[i] |= (binaryData[i*8+j] << (7 - j))
		}
	}

	err := ioutil.WriteFile(outputPath, fileData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write binary data to file: %v", err)
	}

	return nil
}
