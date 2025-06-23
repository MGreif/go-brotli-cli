package internal

import (
	"flag"
	"fmt"
	"os"

	"github.com/andybalholm/brotli"
)

type CompressionConfig struct {
	Help          bool
	InFileName    string
	OutFileName   string
	BufferSize    int
	FlushInterval int
}

func (c *CompressionConfig) SetupFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.InFileName, "i", "", "The target file")
	fs.StringVar(&c.OutFileName, "o", "", "The output file")
	fs.IntVar(&c.BufferSize, "bs", 4096, "Buffer Size")
	fs.IntVar(&c.FlushInterval, "fi", 10, "Flush Interval")
}

func HandleCompress(fileToCompress *os.File, outFile *os.File, config *CompressionConfig) error {
	w := brotli.NewWriterLevel(outFile, brotli.BestCompression)

	bufferSize := config.BufferSize

	var buffer = make([]byte, bufferSize)

	flushInterval := config.FlushInterval
	flushCounter := 0
	for {
		b, err := fileToCompress.Read(buffer)
		if b == 0 {
			// End
			break
		} else if err != nil {
			fmt.Printf("Could not read file %s\n", err)
			return err
		}

		w.Write(buffer)

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
