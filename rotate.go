package main

import (
	"fmt"
)

func (a *App) RotatePDF(inFile string, outFile string, rotation int, pages string) error {
	logger.Printf("inFile: %s, outFile: %s, rotation: %d, pages: %s\n", inFile, outFile, rotation, pages)
	args := []string{"rotate"}
	if rotation != 0 {
		args = append(args, "--angle", fmt.Sprintf("%d", rotation))
	}
	if pages != "" {
		args = append(args, "--page_range", pages)
	}
	if outFile != "" {
		args = append(args, "-o", outFile)
	}
	args = append(args, inFile)
	logger.Println(args)
	fmt.Println("调用pdf.py函数args=", args)
	// return a.cmdRunner(args, "pdf")
	go_client(args)
	return nil

}
