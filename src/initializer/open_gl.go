package initializer

import (
  "log"
  "fmt"
  "strings"
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

func compileShader(source string, shaderType uint32) (uint32, error) {
    shader := gl.CreateShader(shaderType)
    
    csources, free := gl.Strs(source)
    gl.ShaderSource(shader, 1, csources, nil)
    free()
    gl.CompileShader(shader)
    
    var status int32
    gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
    if status == gl.FALSE {
        var logLength int32
        gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
        
        log := strings.Repeat("\x00", int(logLength+1))
        gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
        
        return 0, fmt.Errorf("failed to compile %v: %v", source, log)
    }
    
    return shader, nil
}

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
