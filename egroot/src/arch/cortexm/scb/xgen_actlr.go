package scb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)

type AUX_Periph struct {
	ACTLR ACTLR
}

func (p *AUX_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var AUX = (*AUX_Periph)(unsafe.Pointer(uintptr(0xe000e008)))

type ACTLR_Bits uint32

func (b ACTLR_Bits) Field(mask ACTLR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACTLR_Bits) J(v int) ACTLR_Bits {
	return ACTLR_Bits(bits.Make32(v, uint32(mask)))
}

type ACTLR struct{ mmio.U32 }

func (r *ACTLR) Bits(mask ACTLR_Bits) ACTLR_Bits { return ACTLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ACTLR) StoreBits(mask, b ACTLR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ACTLR) SetBits(mask ACTLR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ACTLR) ClearBits(mask ACTLR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ACTLR) Load() ACTLR_Bits                { return ACTLR_Bits(r.U32.Load()) }
func (r *ACTLR) Store(b ACTLR_Bits)              { r.U32.Store(uint32(b)) }

type ACTLR_Mask struct{ mmio.UM32 }

func (rm ACTLR_Mask) Load() ACTLR_Bits   { return ACTLR_Bits(rm.UM32.Load()) }
func (rm ACTLR_Mask) Store(b ACTLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *AUX_Periph) DISMCYCINT() ACTLR_Mask {
	return ACTLR_Mask{mmio.UM32{&p.ACTLR.U32, uint32(DISMCYCINT)}}
}

func (p *AUX_Periph) DISDEFWBUF() ACTLR_Mask {
	return ACTLR_Mask{mmio.UM32{&p.ACTLR.U32, uint32(DISDEFWBUF)}}
}

func (p *AUX_Periph) DISFOLD() ACTLR_Mask {
	return ACTLR_Mask{mmio.UM32{&p.ACTLR.U32, uint32(DISFOLD)}}
}

func (p *AUX_Periph) DISFPCA() ACTLR_Mask {
	return ACTLR_Mask{mmio.UM32{&p.ACTLR.U32, uint32(DISFPCA)}}
}

func (p *AUX_Periph) DISOOFP() ACTLR_Mask {
	return ACTLR_Mask{mmio.UM32{&p.ACTLR.U32, uint32(DISOOFP)}}
}
