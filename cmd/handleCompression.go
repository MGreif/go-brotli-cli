package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MGreif/brotli-cli/internal"
)

func handleCompression() error {
	compressionConfig.SetupFlags(fsCompression)
	fsCompression.Parse(os.Args[2:])
	if compressionConfig.Help {
		flag.Usage()
		return nil
	}

	fileIn, fileOut, err := prepareFiles(compressionConfig.InFileName, compressionConfig.OutFileName)
	defer fileIn.Close()
	defer fileOut.Close()
	if err != nil {
		return fmt.Errorf("Could not prepare files: %s\n", err)
	}

	fmt.Printf("Start decompressing %s\n", compressionConfig.OutFileName)
	err = internal.HandleCompress(fileIn, fileOut, compressionConfig)
	if err != nil {
		fmt.Printf("Could not compress %s. Error: %s\n", compressionConfig.InFileName, err)
	} else {
		fmt.Printf("Successfully compressed %s to %s\n", compressionConfig.InFileName, compressionConfig.OutFileName)
	}
	return nil
}
