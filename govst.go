package main

/*
#cgo CPPFLAGS: -Ivstsdk2.4 -D__cdecl=""
*/
import "C"

//export Process
func Process(gain float64, input float64) float64 {
     return input*gain
}

func main() {}
