package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MGreif/brotli-cli/internal"
)

type Action int

const (
	Compress Action = iota
	Decompress
)

type GlobalConfig struct {
	help bool
}

func setupGlobalFlags() {
	flag.BoolVar(&globalConfig.help, "h", false, "Print help")
	flag.Usage = func() {
		fmt.Printf("Usage:\n\t%s {compress,decompress}\nActions:\n\tcompress # Compresses the given file\n\tdecompress # Decompresses the given file\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		fsCompression.PrintDefaults()
		fsDecompression.PrintDefaults()
	}
}

func getAction(actionString string) (Action, error) {
	switch actionString {
	case "compress":
		return Compress, nil
	case "decompress":
		return Decompress, nil
	}

	return 0, fmt.Errorf("%s is not a valid action\n", actionString)
}

var (
	globalConfig        GlobalConfig
	compressionConfig   *internal.CompressionConfig   = &internal.CompressionConfig{}
	decompressionConfig *internal.DecompressionConfig = &internal.DecompressionConfig{}
	fsCompression       *flag.FlagSet                 = flag.NewFlagSet("", flag.ExitOnError)
	fsDecompression     *flag.FlagSet                 = flag.NewFlagSet("", flag.ExitOnError)
)

func prepareFiles(inFileName string, outFileName string) (*os.File, *os.File, error) {

	if inFileName == "" || outFileName == "" {
		flag.Usage()
		return nil, nil, fmt.Errorf("Please specify input and output file names")
	}

	fileIn, err := os.Open(inFileName)

	if err != nil {
		return nil, nil, fmt.Errorf("Could not open %s. Err: %s\n", inFileName, err)
	}

	_, err = os.Stat(outFileName)
	if err == nil {
		return nil, nil, fmt.Errorf("File %s already exists. Exiting...\n", outFileName)
	}

	fileOut, err := os.Create(outFileName)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not create %s. Err: %s\n", outFileName, err)
	}
	return fileIn, fileOut, nil
}

func main() {
	setupGlobalFlags()

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	if globalConfig.help {
		flag.Usage()
		os.Exit(0)
	}

	actionString := flag.Arg(0)
	action, err := getAction(actionString)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	switch action {
	case Decompress:
		if err := handleDecompression(); err != nil {
			fmt.Printf("Could not decompress: %s\n", err)
			os.Exit(1)
		}
		break
	case Compress:
		if err := handleCompression(); err != nil {
			fmt.Printf("Could not Compress: %s\n", err)
			os.Exit(1)
		}
		break
	}
}
