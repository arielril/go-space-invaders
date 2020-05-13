package main

import (
	"fmt"
	"runtime"

	"github.com/arielril/go-space-invaders/game"
	"github.com/arielril/go-space-invaders/opengl"
	"github.com/arielril/go-space-invaders/util"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wWidth  = 800
	wHeight = 600
)

var fps util.FPS

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("failed to init glfw: %v", err))
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
}

func main() {
	game.InitObjects()

	win, _ := opengl.NewWindow(wWidth, wHeight, "Space Invaders")

	win.SetKeyCallback(opengl.KeyCallback)
	win.SetCharCallback(opengl.CharCallback)

	opengl.Setup()
	game.Init()
	for !win.ShouldClose() {
		opengl.Reshape(win)
		game.Display(win)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
