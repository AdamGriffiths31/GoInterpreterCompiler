package vm

import (
	"interpreter/code"
	"interpreter/object"
)

type Frame struct {
	cl                 *object.Closure
	instructionPointer int
	basePointer        int
}

func NewFrame(fn *object.Closure, basePointer int) *Frame {
	return &Frame{cl: fn, instructionPointer: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
