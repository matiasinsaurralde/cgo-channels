package main

/*
#include <stdlib.h>

#include "common.h"
#include "message.h"
*/
import "C"

import (
  "unsafe"
)

// This is an exported function. It takes a *C.Char, converts it to a Go string and sends it to the channel.
//export publishString
func publishString(s *C.char) {
  var goString string
  goString = C.GoString(s)
  stringChannel <- goString
  return
}

// This function provides a bridge between the C Message struct and Go. It'll send a Go Message to the channel.
//export publishMessage
func publishMessage(m *C.struct_Message) {
  var data []byte = C.GoBytes(m.data, m.length)

  var message Message = Message{data}

  C.free(unsafe.Pointer(m.data))
  C.free(unsafe.Pointer(m))

  messageChannel <- message

  return
}
