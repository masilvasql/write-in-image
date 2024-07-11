package pkg

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type WriteInImageInput struct {
	Name         string
	TemplatePath string
	OutputPath   string
	FontSize     string
	AlturaTexto  string
	Color        string
	Wg           *sync.WaitGroup
}

func NewWriteInImageInput(name, templatePath, outputPath, fontSize, alturaTexto, color string, wg *sync.WaitGroup) *WriteInImageInput {
	return &WriteInImageInput{
		Name:         name,
		TemplatePath: templatePath,
		OutputPath:   outputPath,
		FontSize:     fontSize,
		AlturaTexto:  alturaTexto,
		Color:        color,
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
	customColor, err := hexToRGBA(input.Color)
	if err != nil {
		log.Fatalf("failed to parse color: %v", err)
	}
	c.SetSrc(image.NewUniform(customColor))

	// Calcula a largura do texto
	face := truetype.NewFace(font, &truetype.Options{Size: tamanhoFonte})
	textWidth := 0
	for _, char := range input.Name {
		awidth, ok := face.GlyphAdvance(rune(char))
		if ok {
			textWidth += awidth.Round()
		}
	}

	alturaTextoInt, err := strconv.Atoi(input.AlturaTexto)
	if err != nil {
		log.Fatalf("failed to parse text height: %v", err)
	}

	// Centraliza o texto
	imageWidth := bounds.Dx()
	x := (imageWidth - textWidth) / 2
	pt := freetype.Pt(x, alturaTextoInt+int(c.PointToFixed(40)>>6))

	_, err = c.DrawString(input.Name, pt)
	if err != nil {
		log.Fatalf("failed to draw string: %v", err)
	}

	fullOutputPath := fmt.Sprintf("%s/%s.jpg", input.OutputPath, input.Name)
	// Verifica se o diretório de saída existe
	if _, err := os.Stat(input.OutputPath); os.IsNotExist(err) {
		// Se não existir, cria
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

func hexToRGBA(hex string) (color.RGBA, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return color.RGBA{}, log.Output(1, "Invalid length of hex string")
	}
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
}
