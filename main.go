package main

/*
#cgo pkg-config: glad
#cgo pkg-config: glfw3

#include <stdlib.h>
#include "glad/gl.h"
#include "GLFW/glfw3.h"
*/
import "C"
import (
	"log"
	"runtime"
	"unsafe"

	"github.com/BoB1Edition/godos/shared"
)

type config struct {
	Fullscreen bool       `json:"fullscreen"`
	Resolution Resolution `json:"resolution"`
}
type Resolution struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func getError() {
	var description *C.char
	e := C.glfwGetError(&description)
	log.Printf("e: %v, description: %v", e, description)
	C.free(unsafe.Pointer(description))
	return
}

func main() {
	conf := new(config)
	shared.CheckErrorP(shared.LoadFromFile("config.json", conf))
	v := C.glfwInit()
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MAJOR, 3)
	C.glfwWindowHint(C.GLFW_CONTEXT_VERSION_MINOR, 2)
	C.glfwWindowHint(C.GLFW_OPENGL_FORWARD_COMPAT, C.GL_TRUE)
	C.glfwWindowHint(C.GLFW_OPENGL_PROFILE, C.GLFW_OPENGL_CORE_PROFILE)
	if v != C.GLFW_TRUE {
		getError()
		return
	}
	defer C.glfwTerminate()
	window := C.glfwCreateWindow(C.int(conf.Resolution.X), C.int(conf.Resolution.Y), C.CString("My Title"), nil, nil)
	if window == nil {
		getError()
		return
	}
	defer C.glfwDestroyWindow(window)
	C.glfwMakeContextCurrent(window)
	C.gladLoadGL(C.GLADloadfunc(C.glfwGetProcAddress))
	//for C.glfwWindowShouldClose(window) == 0 {
	var width, height C.int
	C.glfwGetFramebufferSize(window, &width, &height)
	//C.glViewport(0, 0, width, height)
	//}
	C.glfwSwapInterval(1)
	log.Printf("window: %v", window)
	log.Printf("width: %v", width)
	log.Printf("height: %v", height)
	runtime.LockOSThread()
	for C.glfwWindowShouldClose(window) == 0 {
		C.glfwSwapBuffers(window)
		C.glfwPollEvents()
	}
}
