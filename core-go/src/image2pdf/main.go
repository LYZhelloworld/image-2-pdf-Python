package main

import "github.com/jung-kurt/gofpdf"
import (
	"fmt"
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"
	"sort"
)

func createPDF(pdfname string, files []string) error {
	pdf := gofpdf.New("P", "pt", "", "")
	for i, file := range files {
		width, height := getImageDimension(file)
		pdf.AddPageFormat("", gofpdf.SizeType{Wd: width, Ht: height})
		pdf.Image(file, 0, 0, width, height, false, "", 0, "")
		fmt.Printf("%d: %s\n", i + 1, filepath.Base(file))
	}
	return pdf.OutputFileAndClose(pdfname)
}

func getImageDimension(file string) (float64, float64) {
	f, err := os.Open(file)
	if err != nil {
		return 0, 0
	}
	defer f.Close()
	
	image, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0
	}
	return float64(image.Width), float64(image.Height)
}

func getImages(path string) []string {
	result := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		
		if strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") ||
			strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			result = append(result, path)
		}
		
		return nil
	})
	if err != nil {
		fmt.Println(fmt.Errorf("An error occurred when walking directory: %s\n%w", path, err))
		return result
	}
	
	sort.Strings(result)
	return result
}

func getPDFName(path string) string {
	return filepath.Join(filepath.Dir(path), filepath.Base(path) + ".pdf")
}

func main() {
	if len(os.Args) > 1 {
		for _, directory := range os.Args[1:] {
			images := getImages(directory)
			if len(images) > 0 {
				pdf := getPDFName(directory)
				err := createPDF(pdf, images)
				if(err == nil) {
					fmt.Printf("PDF file created: %s with %d image file(s).\n",
						pdf, len(images))
				} else {
					fmt.Println(fmt.Errorf("An error occurred when creating PDF file: %s\n%w",
						pdf, err))
				}
			}
		}
		fmt.Print("Press any key to continue...")
		fmt.Scanln()
	} else {
		fmt.Println(fmt.Errorf(
			"Usage: %s <directory that contains images> [<directory2 that contains images> ...]\n" +
			"You can drag and drop folders onto the script.", os.Args[0]))
	}
}
