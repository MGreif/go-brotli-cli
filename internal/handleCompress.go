package internal

import (
	"flag"
	"fmt"
	"io"

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
	fs.StringVar(&c.InFileName, "i", "", "Input file (default stdin)")
	fs.StringVar(&c.OutFileName, "o", "", "Output file (default stdout)")
	fs.IntVar(&c.BufferSize, "bs", 4096, "Buffer Size")
	fs.IntVar(&c.FlushInterval, "fi", 10, "Flush Interval")
}

func HandleCompress(toCompress io.Reader, out io.Writer, config *CompressionConfig) error {
	w := brotli.NewWriterLevel(out, brotli.BestCompression)

	bufferSize := config.BufferSize

	var buffer = make([]byte, bufferSize)

	flushInterval := config.FlushInterval
	flushCounter := 0
	for {
		b, err := toCompress.Read(buffer)
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
