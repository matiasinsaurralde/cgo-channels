# cgo-channels

A quick program to explore the idea of sending data to a Go channel from C.

Uses [cgo](https://golang.org/cmd/cgo/).

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
Receiving: From Go!
Receiving: Hello from C
Receiving: From Go!
Receiving: Hello from C
```
