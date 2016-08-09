package main

import "C"

// This is an exported function. It takes a *C.Char, converts it to a Go string and sends it to the channel.
//export publishString
func publishString(s *C.char) {
  var goString string
  goString = C.GoString(s)
  stringChannel <- goString
  return
}
