package pkg

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/golang/freetype"
)

type WriteInImageInput struct {
	Name         string
	TemplatePath string
	OutputPath   string
	FontSize     string
	Wg           *sync.WaitGroup
}

func NewWriteInImageInput(name, templatePath, outputPath, fontSize string, wg *sync.WaitGroup) *WriteInImageInput {
	return &WriteInImageInput{
		Name:         name,
		TemplatePath: templatePath,
		OutputPath:   outputPath,
		FontSize:     fontSize,
		Wg:           wg,
	}

}

func WriteInImage(input *WriteInImageInput) {

	imgFile, err := os.Open(input.TemplatePath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}
	relativePath := filepath.Join(cwd, "fonts/Corinthia-Bold.ttf")
	fontBytes, err := os.ReadFile(relativePath)
	if err != nil {
		log.Fatalf("failed to read font: %v", err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	tamanhoFonte, err := strconv.ParseFloat(input.FontSize, 64)
	if err != nil {
		log.Fatalf("failed to parse font size: %v", err)
	}
	c.SetFontSize(tamanhoFonte)
	c.SetClip(bounds)
	c.SetDst(newImg)
	c.SetSrc(image.Black)

	pt := freetype.Pt(500, 650+int(c.PointToFixed(40)>>6))

	_, err = c.DrawString(input.Name, pt)
	if err != nil {
		log.Fatalf("failed to draw string: %v", err)
	}

	fullOutputPath := fmt.Sprintf("%s/%s.jpg", input.OutputPath, input.Name)
	// verifica se o diretório de saída existe
	if _, err := os.Stat(input.OutputPath); os.IsNotExist(err) {
		// se não existir, cria
		err = os.Mkdir(input.OutputPath, 0755)
		if err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	outFile, err := os.Create(fullOutputPath)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, newImg, nil)
	if err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}

	log.Println("Image processed and saved to output.jpg")
	input.Wg.Done()
}
