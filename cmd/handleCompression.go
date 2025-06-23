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

	fileIn := os.Stdin
	fileOut := os.Stdout
	var err error = nil
	if compressionConfig.InFileName != "" {
		fileIn, err = os.Open(compressionConfig.InFileName)
		defer fileIn.Close()
		if err != nil {
			return fmt.Errorf("Could not open %s. Err: %s\n", compressionConfig.InFileName, err)
		}

	}

	if compressionConfig.OutFileName != "" {
		_, err = os.Stat(compressionConfig.OutFileName)
		if err == nil {
			return fmt.Errorf("File %s already exists. Exiting...\n", compressionConfig.OutFileName)
		}

		fileOut, err = os.Create(compressionConfig.OutFileName)
		defer fileOut.Close()
		if err != nil {
			return fmt.Errorf("Could not create %s. Err: %s\n", compressionConfig.OutFileName, err)
		}
	}

	err = internal.HandleCompress(fileIn, fileOut, compressionConfig)
	if err != nil {
		fmt.Printf("Could not compress %s. Error: %s\n", compressionConfig.InFileName, err)
	}
	return nil
}
