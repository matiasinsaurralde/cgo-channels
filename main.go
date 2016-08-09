package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "common.h"
#include "message.h"

// These are the functions exported from chan.go:

// This function loops forever, and calls publishString every 2 seconds.
void sendToStringChannel() {
  while(1) {

    // This will call publishString(), as defined in chan.go:
    publishString("Hello from C");

    // Sleep some time:
    usleep(2000000);

  };
};

// This function loops forever, and calls publishMessage every 3 seconds.
void sendToMessageChannel() {
  while(1) {

    struct Message* message = (struct Message*) malloc(sizeof(struct Message));

    char *data;
    data = (char*)malloc( 8 * sizeof(char) );
    strcpy(data, "hello");

    message->data = data;
    message->length = sizeof(data);

    // This will call publishMessage(), as defined in chan.go:
    publishMessage(message);

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

type Message struct {
  data []byte
}

// We declare the channel at this level, chan.go needs to see this too!
var stringChannel chan string

var messageChannel chan Message

func main() {

  // Initializes a channel with "string" type.
  stringChannel = make(chan string)

  // Initialize a channel with "Message" type.
  messageChannel = make(chan Message)

  // This goroutine calls the C function that sends data to stringChannel:
  go func() {
    C.sendToStringChannel()
  }()

  // This goroutine calls the C function that sends data to messageChannel:
  go func() {
    C.sendToMessageChannel()
  }()

  // This goroutine sends a string to the channel every 2 seconds:
  go func() {
    for {
      someString := "From Go!"
      stringChannel <- someString
      time.Sleep( 2 * time.Second )
    }
  }()

  // This goroutine sends a "Message" to the channel every 3 seconds:
  go func() {
    for {
      someMessage := Message{ data: []byte{1, 2, 3} }
      messageChannel <- someMessage
      time.Sleep( 3 * time.Second )
    }
  }()

  // This is our "main" loop, it prints every message received by both channels:
  for {
    select {
    case receivedString := <-stringChannel:
      fmt.Println("stringChannel receives:", receivedString)
    case receivedMessage := <-messageChannel:
      fmt.Println("messageChannel receives:", receivedMessage)
    }
  }
}
