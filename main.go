package main

import (
	"fmt"
	"log"
	"os"
	"phototobinary/convert"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide the mode (encode/decode) and the path to the file or image")
	}

	mode := os.Args[1]
	filePath := os.Args[2]

	switch mode {
	case "encode":
		encode(filePath)
	case "decode":
		if len(os.Args) < 4 {
			log.Fatal("Please provide the output file path for decoding")
		}
		outputFilePath := os.Args[3]
		decode(filePath, outputFilePath)
	default:
		log.Fatalf("Unknown mode: %s. Use 'encode' or 'decode'.", mode)
	}
}

func encode(filePath string) {
	binaryData, err := convert.FileToBinary(filePath)
	if err != nil {
		log.Fatalf("Failed to convert file to binary: %v", err)
	}

	outputImagePath := "output_image.png"
	err = convert.BinaryToImage(binaryData, outputImagePath)
	if err != nil {
		log.Fatalf("Failed to convert binary data to image: %v", err)
	}

	fmt.Printf("Binary data successfully converted to image and saved to %s\n", outputImagePath)
}

func decode(imagePath, outputFilePath string) {
	binaryData, err := convert.ImageToBinary(imagePath)
	if err != nil {
		log.Fatalf("Failed to convert image to binary: %v", err)
	}

	err = convert.BinaryToFile(binaryData, outputFilePath)
	if err != nil {
		log.Fatalf("Failed to convert binary data to file: %v", err)
	}

	fmt.Printf("Binary data successfully converted to file and saved to %s\n", outputFilePath)
}
