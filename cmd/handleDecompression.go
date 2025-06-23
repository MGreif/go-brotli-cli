package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MGreif/brotli-cli/internal"
)

func handleDecompression() error {
	decompressionConfig.SetupFlags(fsDecompression)
	fsDecompression.Parse(os.Args[2:])
	if decompressionConfig.Help {
		flag.Usage()
		return nil
	}

	fileIn := os.Stdin
	fileOut := os.Stdout
	var err error = nil
	if decompressionConfig.InFileName != "" {
		fileIn, err = os.Open(decompressionConfig.InFileName)
		defer fileIn.Close()
		if err != nil {
			return fmt.Errorf("Could not open %s. Err: %s\n", decompressionConfig.InFileName, err)
		}
	}

	if decompressionConfig.OutFileName != "" {
		_, err = os.Stat(decompressionConfig.OutFileName)
		if err == nil {
			return fmt.Errorf("File %s already exists. Exiting...\n", decompressionConfig.OutFileName)
		}

		fileOut, err = os.Create(decompressionConfig.OutFileName)
		defer fileOut.Close()
		if err != nil {
			return fmt.Errorf("Could not create %s. Err: %s\n", decompressionConfig.OutFileName, err)
		}
	}
	err = internal.HandleDecompress(fileIn, fileOut, decompressionConfig)
	if err != nil {
		fmt.Printf("Could not decompress %s. Error: %s\n", decompressionConfig.InFileName, err)
	}
	return nil
}
