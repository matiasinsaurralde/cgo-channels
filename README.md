# cgo-channels

A quick program to explore the idea of sending data to a Go channel from C.

The program uses `select` across two channels:

`stringChannel`: this channel passes strings from C to Go.

`messageChannel`: passes `Message` data structures.

The `publishMessage` takes a C Message (as defined in [`message.h`](message.h)) and turns it into a Go `Message`, then sends it over the `messageChannel`.

This program makes use of [cgo](https://golang.org/cmd/cgo/).

## Usage

Fetch the repo:

```
% go get github.com/matiasinsaurralde/cgo-channels
```

Chdir somewhere (`~` in this case) and build the program:
```
% cd ~
% go build github.com/matiasinsaurralde/cgo-channels
```

Run the program:
```
% ./cgo-channels
messageChannel receives: {[1 2 3]}
messageChannel receives: {[104 101 108 108 111 0 0 0]}
stringChannel receives: Hello from C
stringChannel receives: From Go!
stringChannel receives: Hello from C
stringChannel receives: From Go!
```
