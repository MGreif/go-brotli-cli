# A small and simple brotli CLI tool

A simple CLI tool to compress and decompress data using [brotli](https://github.com/google/brotli).
This projects builds upon [Andybalholm's brotli library](https://github.com/andybalholm/brotli).

## Features

Take input from:
- Stdin
- A file (`-i`)

Put output to:
- Stdout
- A file (`-o`)


## Usage

### `./brotli-cli -h`

```
Usage:
        ./build/brotli-cli {compress,decompress}
Actions:
        compress # Compresses the given input
        decompress # Decompresses the given input
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
        Input file (default stdin)
  -o string
        Output file (default stdout)
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
        Input file (default stdin)
  -o string
        Output file (default stdout)
```

### `./brotli-cli compress -i file.html -o file.br`
### `./brotli-cli decompress -i file.br -o file.html`

### `echo "<title>Fancy Page</title>" | ./brotli-cli compress -o out.br`


### `echo "<title>Fancy Page</title>" | ./brotli-cli compress | xxd`

```
00000000: 8bff 07f8 8f94 aced 9112 c028 2cc7 dcf2  ...........(,...
00000010: ba90 4354 794c 23b0 c125 04c1 3a4f 60ff  ..CTyL#..%..:O`.
00000020: 3c00 8001                                <...
```

### `echo "<title>Fancy Page</title>" | ./brotli-cli compress | ./brotli-cli decompress`
```
<title>Fancy Page</title>
```