# cgo-channels

A quick program to explore the idea of sending data to a Go channel from C.

The program uses `select` across two channels:

`stringChannel`: this channel passes strings from C to Go.

`messageChannel`: passes `Message` data structures.

The `publishMessage` takes a C Message (as defined in [`message.h`](message.h)) and turns it into a Go `Message`, then sends it over the `messageChannel`.

This program makes use of [cgo](https://golang.org/cmd/cgo/).

## Usage

Fetch the repo:

```bash
% go get github.com/matiasinsaurralde/cgo-channels
```

Chdir somewhere (`~` in this case) and build the program:
```bash
% cd ~
% go build github.com/matiasinsaurralde/cgo-channels
```

Run the program:
```bash
% ./cgo-channels
messageChannel receives: {[1 2 3]}
messageChannel receives: {[104 101 108 108 111 0 0 0]}
stringChannel receives: Hello from C
stringChannel receives: From Go!
stringChannel receives: Hello from C
stringChannel receives: From Go!
```

## OpenMP support

I have included a very simple C program that uses OpenMP to parallelize the communication with Go channels, this has been tested under Linux and OSX.

There's a build tag that will include the right flags for each platform:

### Linux

This command will include the OpenMP sample and `-fopenmp` flag:

```bash
% go build github.com/matiasinsaurralde/cgo-channels -tags 'openmp'
```

Expected output:
```bash
% ./cgo-channels
stringChannel receives: Hello from C
messageChannel receives: {[1 2 3]}
stringChannel receives: From Go!
messageChannel receives: {[104 101 108 108 111 0 0 0]}
stringChannel receives: Hello from thread 0, nthreads 8
stringChannel receives: Hello from thread 3, nthreads 8
stringChannel receives: Hello from thread 6, nthreads 8
stringChannel receives: Hello from thread 4, nthreads 8
stringChannel receives: Hello from thread 7, nthreads 8
stringChannel receives: Hello from thread 2, nthreads 8
stringChannel receives: Hello from thread 5, nthreads 8
stringChannel receives: Hello from thread 1, nthreads 8
```

### OSX

`clang-omp` is a requirement, I installed it with:

```bash
% brew install clang-omp
```

You may find more details in [its website](https://clang-omp.github.io/).

We will need to override the default C compiler. To use `clang-omp` and the `openmp` build tag (for Go), run:

```bash
CC='clang-omp' go build -tags 'openmp'
```
