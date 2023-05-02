package figure

import (
	"github.com/IKrehan/go_game/src/util"
	"github.com/go-gl/gl/v2.1/gl"
)

type Figure struct {
	Vertices [][]float32
}

func New(vertices [][]float32) Figure {
	e := Figure{vertices}
	return e
}

func (fig Figure) GetFlattenedPoints() []float32 {
	return util.Flatten(fig.Vertices)
}

func (fig Figure) GetBufferSize() int {
	return 4 * len(fig.GetFlattenedPoints())
}

func (fig Figure) MakeVAO() uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, fig.GetBufferSize(), gl.Ptr(fig.GetFlattenedPoints()), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
