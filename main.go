package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	file := "mertcaliskan.txt"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("Dosya Okunamadı %s", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	_ = pdf.OutputFileAndClose("MertCaliskan.pdf")
	fmt.Println("PDF Oluşturuldu...")
}
