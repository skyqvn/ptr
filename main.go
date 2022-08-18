package main

import (
	"fmt"
	"unsafe"
)

type Ptr struct {
	ptr uint64
}

func (p *Ptr) Add() {
	p.ptr += 1
}

func (p *Ptr) Move(off int64) {
	p.ptr += uint64(off)
}

func (p *Ptr) MoveTo(pos uint64) {
	p.ptr = pos
}

func (p *Ptr) Get() byte {
	return *(*byte)(p.GetPtr())
}

func (p *Ptr) GetPtr() unsafe.Pointer {
	return unsafe.Pointer(uintptr(p.ptr))
}

func (p *Ptr) Set(b byte) {
	*(*byte)(p.GetPtr()) = b
}

func (p *Ptr)GetAt(idx uint8)byte  {
	if idx>7{
		panic(fmt.Errorf("idx can't bigger than 7"))
		return 0
	}
	b:=p.Get()
	return (((1<<idx)^0)&b)>>idx
}
