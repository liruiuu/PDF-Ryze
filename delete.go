package main

func (a *App) DeletePDF(inFile string, outFile string, pages string) error {
	logger.Printf("inFile: %s, outFile: %s, pages: %s\n", inFile, outFile, pages)
	args := []string{"delete"}
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
