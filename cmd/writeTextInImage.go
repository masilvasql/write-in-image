/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/masilvasql/write-in-image/pkg"
	"github.com/spf13/cobra"
)

var startTime time.Time
var endTime time.Time

// writeTextInImageCmd represents the writeTextInImage command
var writeTextInImageCmd = &cobra.Command{
	Use:   "writeTextInImage",
	Short: "Write text in an image",
	Long:  `You will be able to write text in an image with this command and some parameters like file, output-path, template, font-size, text-height and font-color`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		outputPath, _ := cmd.Flags().GetString("outoput-path")
		templatePath, _ := cmd.Flags().GetString("template")
		fontSize, _ := cmd.Flags().GetString("font-size")
		textHeight, _ := cmd.Flags().GetString("text-height")
		fontColor, _ := cmd.Flags().GetString("font-color")

		arquivo, err := os.Open(filePath)
		defer arquivo.Close()
		if err != nil {
			log.Fatalf("Falha ao abrir arquivo: %v", err)
		}

		scanner := bufio.NewScanner(arquivo)
		var wg sync.WaitGroup
		for scanner.Scan() {
			wg.Add(1)
			input := pkg.NewWriteInImageInput(scanner.Text(), templatePath, outputPath, fontSize, textHeight, fontColor, &wg)
			go pkg.WriteInImage(input)
		}
		wg.Wait()

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting writeTextInImage command")

		filePath, _ := cmd.Flags().GetString("file")
		outputPath, _ := cmd.Flags().GetString("outoput-path")
		templatePath, _ := cmd.Flags().GetString("template")

		if filePath == "" || outputPath == "" || templatePath == "" {
			cmd.Println("The flags file, outoput-path and template are required")
			os.Exit(1)
		}

		startTime = time.Now()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Finishing writeTextInImage command")
		endTime = time.Now()
		elapsedTime := endTime.Sub(startTime)
		cmd.Printf("Total execution time: %v\n", elapsedTime)
	},
}

func init() {
	rootCmd.AddCommand(writeTextInImageCmd)

	writeTextInImageCmd.Flags().StringP("file", "f", "", "File path With the text to write in the image (required)")
	writeTextInImageCmd.MarkFlagRequired("file")
	writeTextInImageCmd.Flags().StringP("outoput-path", "o", "", "Output path to save the new image (required)")
	writeTextInImageCmd.MarkFlagRequired("outoput-path")
	writeTextInImageCmd.Flags().StringP("template", "t", "", "Template image path (required)")
	writeTextInImageCmd.MarkFlagRequired("template")
	writeTextInImageCmd.Flags().StringP("font-size", "s", "72", "Font size (optional) default 72")
	writeTextInImageCmd.Flags().StringP("text-height", "a", "650", "Text height (optional) default 650")
	writeTextInImageCmd.Flags().StringP("font-color", "c", "000000", "Font color (optional) default is black")

}
