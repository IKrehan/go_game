package initializers 

import (
  "log"
  "github.com/go-gl/gl/v4.1-core/gl"
)

const (
  vertexShaderSource = `
  #version 410
  in vec3 vp;
  void main() {
    gl_Position = vec4(vp, 1.0);
  }
  ` + "\x00"

  fragmentShaderSource = `
  #version 410
  out vec4 frag_colour;
  void main() {
    frag_colour = vec4(1, 1, 1, 1);
  }
  ` + "\x00"
)



func OpenGL() uint32 {
  if err := gl.Init(); err != nil {
    panic(err)
  }
  version := gl.GoStr(gl.GetString(gl.VERSION))
  log.Println("OpenGL version", version)

  vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
  if err != nil {
    panic(err)
  }
  fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
  if err != nil {
    panic(err)
  }


  program := gl.CreateProgram()
  gl.AttachShader(program, vertexShader)
  gl.AttachShader(program, fragmentShader)
  gl.LinkProgram(program)
  return program
}
