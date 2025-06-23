package internal

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/andybalholm/brotli"
)

type DecompressionConfig struct {
	Help          bool
	InFileName    string
	OutFileName   string
	BufferSize    int
	FlushInterval int
	DontTrimZeros bool
}

func (d *DecompressionConfig) SetupFlags(fs *flag.FlagSet) {
	fs.StringVar(&d.InFileName, "i", "", "The target file")
	fs.StringVar(&d.OutFileName, "o", "", "The output file")
	fs.IntVar(&d.BufferSize, "bs", 4096, "Buffer Size")
	fs.IntVar(&d.FlushInterval, "fi", 10, "Flush Interval")
	fs.BoolVar(&d.DontTrimZeros, "dont-trim-zeros", false, "Dont trim zeroes at the end of the file")
}

func HandleDecompress(fileToDecompress *os.File, outFile *os.File, config *DecompressionConfig) error {
	r := brotli.NewReader(fileToDecompress)
	bufferSize := config.BufferSize

	w := bufio.NewWriterSize(outFile, bufferSize)

	var buffer = make([]byte, bufferSize)

	flushInterval := config.FlushInterval
	flushCounter := 0
	for {
		b, err := r.Read(buffer)
		if b == 0 {
			// End
			break
		} else if err != nil {
			fmt.Printf("Could not read file %s\n", err)
			return err
		}

		if config.DontTrimZeros {
			w.Write(buffer)
		} else {
			// Trim trailing \x00
			w.Write([]byte(strings.TrimRight(string(buffer), "\x00")))
		}

		if flushCounter == flushInterval {
			flushCounter = 0
			w.Flush()
		} else {
			flushCounter++
		}

		clear(buffer)
	}
	w.Flush()

	return nil
}
