package main

import "fmt"

func (a *App) MakeDualLayerPDF(
	inFile string,
	outFile string,
	dpi int,
	pages string,
	lang string,
) error {
	logger.Printf("inFile: %s, outFile: %s, dpi: %d, pages: %s, lang: %s\n", inFile, outFile, dpi, pages, lang)
	args := []string{"dual", "--dpi", fmt.Sprintf("%d", dpi), "--lang", lang}
	if pages != "" {
		args = append(args, "--page_range", pages)
	}
	if outFile != "" {
		args = append(args, "-o", outFile)
	}
	args = append(args, inFile)
	logger.Println(args)
	// return a.cmdRunner(args, "pdf")
	go_client(args)
	return nil
}
