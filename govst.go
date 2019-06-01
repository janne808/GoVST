package main

/*
#cgo CPPFLAGS: -Ivstsdk2.4 -D__cdecl=""
*/
import "C"

import "unsafe"

// lol
const arrayLength = 1<<30
     
//export ProcessReplacing
func ProcessReplacing(input1 *float32, input2 *float32, output1 *float32, output2 *float32,
     	   	      sampleFrames int, gain float32) {
     // make arrays
     in1 := (*[arrayLength]C.float)(unsafe.Pointer(input1))
     in2 := (*[arrayLength]C.float)(unsafe.Pointer(input2))
     out1 := (*[arrayLength]C.float)(unsafe.Pointer(output1))
     out2 := (*[arrayLength]C.float)(unsafe.Pointer(output2))

     // process
     for i := 0; i < sampleFrames; i++ {
     	 out1[i] = (C.float)(gainfunc((float64)(in1[i]), (float64)(gain)))
     	 out2[i] = (C.float)(gainfunc((float64)(in2[i]), (float64)(gain)))
     }
}

//export ProcessDoubleReplacing
func ProcessDoubleReplacing(input1 *float64, input2 *float64, output1 *float64, output2 *float64,
     			    sampleFrames int, gain float64) {
     // make arrays
     in1 := (*[arrayLength]C.double)(unsafe.Pointer(input1))
     in2 := (*[arrayLength]C.double)(unsafe.Pointer(input2))
     out1 := (*[arrayLength]C.double)(unsafe.Pointer(output1))
     out2 := (*[arrayLength]C.double)(unsafe.Pointer(output2))

     // process
     for i := 0; i < sampleFrames; i++ {
     	 out1[i] = (C.double)(gainfunc((float64)(in1[i]), gain))
     	 out2[i] = (C.double)(gainfunc((float64)(in2[i]), gain))
     }
}

func gainfunc(input float64, gain float64) float64 {
     return input*gain
}

func main() {}
