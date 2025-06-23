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

	fileIn, fileOut, err := prepareFiles(decompressionConfig.InFileName, decompressionConfig.OutFileName)
	defer fileIn.Close()
	defer fileOut.Close()
	if err != nil {
		return fmt.Errorf("Could not prepare files: %s\n", err)
	}

	fmt.Printf("Start decompressing %s\n", decompressionConfig.OutFileName)
	err = internal.HandleDecompress(fileIn, fileOut, decompressionConfig)
	if err != nil {
		fmt.Printf("Could not decompress %s. Error: %s\n", decompressionConfig.InFileName, err)
	} else {
		fmt.Printf("Successfully decompressed %s to %s\n", decompressionConfig.InFileName, decompressionConfig.OutFileName)
	}
	return nil
}
