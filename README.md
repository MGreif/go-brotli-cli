# A small and simple brotli CLI tool

A simple CLI tool to compress and decompress files using [brotli](https://github.com/google/brotli).
This projects builds upon [Andybalholm's brotli library](https://github.com/andybalholm/brotli).


## Usage

### `./brotli-cli -h`

```
Usage:
        ./build/brotli-cli {compress,decompress}
Actions:
        compress # Compresses the given file
        decompress # Decompresses the given file
Options:
  -h    Print help
```

### `./brotli-cli compress -h`

```
Usage:
  -bs int
        Buffer Size (default 4096)
  -fi int
        Flush Interval (default 10)
  -i string
        Target file
  -o string
        Output file
```

### `./brotli-cli decompress -h`

```
Usage:
  -bs int
        Buffer Size (default 4096)
  -dont-trim-zeros
        Dont trim zeroes at the end of the file
  -fi int
        Flush Interval (default 10)
  -i string
        Target file
  -o string
        Output file
```

### `./brotli-cli compress -i file.html -o file.br`

```
Start compressing file.html
Successfully compressed file.html to file.br
```

### `./brotli-cli decompress -i file.br -o file.html`

```
Start decompressing file.br
Successfully decompressed file.br to file.html
```