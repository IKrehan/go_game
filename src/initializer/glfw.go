package initializer

import "github.com/go-gl/glfw/v3.2/glfw"

func Glfw(width int, height int, title string) *glfw.Window {
  if err := glfw.Init(); err != nil {
    panic(err)
  }

  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

  window, err := glfw.CreateWindow(width, height, title, nil, nil)
  if err != nil {
    panic(err)
  }
  window.MakeContextCurrent()

  return window
}
