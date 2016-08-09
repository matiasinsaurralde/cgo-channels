package main

/*
#include <unistd.h>

void publishString();

// This function loops forever, and calls publishString (defined in chan.go) every 2 seconds.
void runMessageLoop() {
  while(1) {

    // This will call publishString(), as defined in chan.go:
    publishString("Hello from C");

    // Sleep some time:
    usleep(2000000);

  };
};

*/
import "C"

import(
  "fmt"
  "time"
)

// We declare the channel at this level, chan.go needs to see this too!
var stringChannel chan string

func main() {

  // Initializes a channel with "string" type.
  stringChannel = make(chan string)

  // This goroutine calls the C function:
  go func() {
    C.runMessageLoop()
  }()

  // This goroutine sends a string to the channel every 2 seconds:
  go func() {
    for {
      someString := "From Go!"
      stringChannel <- someString
      time.Sleep( 2 * time.Second )
    }
  }()

  // This is our "main" loop, it prints every message received by the channel:
  for {
    receivedString := <-stringChannel
    fmt.Println("Receiving:", receivedString)
  }
}
