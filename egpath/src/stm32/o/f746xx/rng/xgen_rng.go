package rng

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type RNG_Periph struct {
	CR CR
	SR SR
	DR DR
}

func (p *RNG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var RNG = (*RNG_Periph)(unsafe.Pointer(uintptr(mmap.RNG_BASE)))

type CR_Bits uint32

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RNG_Periph) RNGEN() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RNGEN)}}
}

func (p *RNG_Periph) IE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IE)}}
}

type SR_Bits uint32

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RNG_Periph) DRDY() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(DRDY)}}
}

func (p *RNG_Periph) CECS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CECS)}}
}

func (p *RNG_Periph) SECS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SECS)}}
}

func (p *RNG_Periph) CEIS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CEIS)}}
}

func (p *RNG_Periph) SEIS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SEIS)}}
}

type DR_Bits uint32

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }
