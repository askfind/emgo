// +build f411xe

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type GPIO_Periph struct {
	MODER   MODER
	OTYPER  OTYPER
	OSPEEDR OSPEEDR
	PUPDR   PUPDR
	IDR     IDR
	ODR     ODR
	BSRRL   BSRRL
	BSRRH   BSRRH
	LCKR    LCKR
	AFR     [2]AFR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

var GPIOH = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE)))

var GPIOI = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOI_BASE)))

var GPIOJ = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOJ_BASE)))

var GPIOK = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOK_BASE)))

type MODER_Bits uint32

type MODER struct{ mmio.U32 }

func (r *MODER) Bits(mask MODER_Bits) MODER_Bits { return MODER_Bits(r.U32.Bits(uint32(mask))) }
func (r *MODER) StoreBits(mask, b MODER_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MODER) SetBits(mask MODER_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *MODER) ClearBits(mask MODER_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *MODER) Load() MODER_Bits                { return MODER_Bits(r.U32.Load()) }
func (r *MODER) Store(b MODER_Bits)              { r.U32.Store(uint32(b)) }

type MODER_Mask struct{ mmio.UM32 }

func (rm MODER_Mask) Load() MODER_Bits   { return MODER_Bits(rm.UM32.Load()) }
func (rm MODER_Mask) Store(b MODER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODER0() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER0)}}
}

func (p *GPIO_Periph) MODER1() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER1)}}
}

func (p *GPIO_Periph) MODER2() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER2)}}
}

func (p *GPIO_Periph) MODER3() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER3)}}
}

func (p *GPIO_Periph) MODER4() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER4)}}
}

func (p *GPIO_Periph) MODER5() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER5)}}
}

func (p *GPIO_Periph) MODER6() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER6)}}
}

func (p *GPIO_Periph) MODER7() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER7)}}
}

func (p *GPIO_Periph) MODER8() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER8)}}
}

func (p *GPIO_Periph) MODER9() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER9)}}
}

func (p *GPIO_Periph) MODER10() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER10)}}
}

func (p *GPIO_Periph) MODER11() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER11)}}
}

func (p *GPIO_Periph) MODER12() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER12)}}
}

func (p *GPIO_Periph) MODER13() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER13)}}
}

func (p *GPIO_Periph) MODER14() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER14)}}
}

func (p *GPIO_Periph) MODER15() MODER_Mask {
	return MODER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODER15)}}
}

type OTYPER_Bits uint32

type OTYPER struct{ mmio.U32 }

func (r *OTYPER) Bits(mask OTYPER_Bits) OTYPER_Bits { return OTYPER_Bits(r.U32.Bits(uint32(mask))) }
func (r *OTYPER) StoreBits(mask, b OTYPER_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OTYPER) SetBits(mask OTYPER_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *OTYPER) ClearBits(mask OTYPER_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *OTYPER) Load() OTYPER_Bits                 { return OTYPER_Bits(r.U32.Load()) }
func (r *OTYPER) Store(b OTYPER_Bits)               { r.U32.Store(uint32(b)) }

type OTYPER_Mask struct{ mmio.UM32 }

func (rm OTYPER_Mask) Load() OTYPER_Bits   { return OTYPER_Bits(rm.UM32.Load()) }
func (rm OTYPER_Mask) Store(b OTYPER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT_0() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_0)}}
}

func (p *GPIO_Periph) OT_1() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_1)}}
}

func (p *GPIO_Periph) OT_2() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_2)}}
}

func (p *GPIO_Periph) OT_3() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_3)}}
}

func (p *GPIO_Periph) OT_4() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_4)}}
}

func (p *GPIO_Periph) OT_5() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_5)}}
}

func (p *GPIO_Periph) OT_6() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_6)}}
}

func (p *GPIO_Periph) OT_7() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_7)}}
}

func (p *GPIO_Periph) OT_8() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_8)}}
}

func (p *GPIO_Periph) OT_9() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_9)}}
}

func (p *GPIO_Periph) OT_10() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_10)}}
}

func (p *GPIO_Periph) OT_11() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_11)}}
}

func (p *GPIO_Periph) OT_12() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_12)}}
}

func (p *GPIO_Periph) OT_13() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_13)}}
}

func (p *GPIO_Periph) OT_14() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_14)}}
}

func (p *GPIO_Periph) OT_15() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OT_15)}}
}

type OSPEEDR_Bits uint32

type OSPEEDR struct{ mmio.U32 }

func (r *OSPEEDR) Bits(mask OSPEEDR_Bits) OSPEEDR_Bits { return OSPEEDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OSPEEDR) StoreBits(mask, b OSPEEDR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OSPEEDR) SetBits(mask OSPEEDR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OSPEEDR) ClearBits(mask OSPEEDR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OSPEEDR) Load() OSPEEDR_Bits                  { return OSPEEDR_Bits(r.U32.Load()) }
func (r *OSPEEDR) Store(b OSPEEDR_Bits)                { r.U32.Store(uint32(b)) }

type OSPEEDR_Mask struct{ mmio.UM32 }

func (rm OSPEEDR_Mask) Load() OSPEEDR_Bits   { return OSPEEDR_Bits(rm.UM32.Load()) }
func (rm OSPEEDR_Mask) Store(b OSPEEDR_Bits) { rm.UM32.Store(uint32(b)) }

type PUPDR_Bits uint32

type PUPDR struct{ mmio.U32 }

func (r *PUPDR) Bits(mask PUPDR_Bits) PUPDR_Bits { return PUPDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PUPDR) StoreBits(mask, b PUPDR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PUPDR) SetBits(mask PUPDR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PUPDR) ClearBits(mask PUPDR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PUPDR) Load() PUPDR_Bits                { return PUPDR_Bits(r.U32.Load()) }
func (r *PUPDR) Store(b PUPDR_Bits)              { r.U32.Store(uint32(b)) }

type PUPDR_Mask struct{ mmio.UM32 }

func (rm PUPDR_Mask) Load() PUPDR_Bits   { return PUPDR_Bits(rm.UM32.Load()) }
func (rm PUPDR_Mask) Store(b PUPDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPDR0() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR0)}}
}

func (p *GPIO_Periph) PUPDR1() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR1)}}
}

func (p *GPIO_Periph) PUPDR2() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR2)}}
}

func (p *GPIO_Periph) PUPDR3() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR3)}}
}

func (p *GPIO_Periph) PUPDR4() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR4)}}
}

func (p *GPIO_Periph) PUPDR5() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR5)}}
}

func (p *GPIO_Periph) PUPDR6() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR6)}}
}

func (p *GPIO_Periph) PUPDR7() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR7)}}
}

func (p *GPIO_Periph) PUPDR8() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR8)}}
}

func (p *GPIO_Periph) PUPDR9() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR9)}}
}

func (p *GPIO_Periph) PUPDR10() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR10)}}
}

func (p *GPIO_Periph) PUPDR11() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR11)}}
}

func (p *GPIO_Periph) PUPDR12() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR12)}}
}

func (p *GPIO_Periph) PUPDR13() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR13)}}
}

func (p *GPIO_Periph) PUPDR14() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR14)}}
}

func (p *GPIO_Periph) PUPDR15() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PUPDR15)}}
}

type IDR_Bits uint32

type IDR struct{ mmio.U32 }

func (r *IDR) Bits(mask IDR_Bits) IDR_Bits { return IDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDR) StoreBits(mask, b IDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDR) SetBits(mask IDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IDR) ClearBits(mask IDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IDR) Load() IDR_Bits              { return IDR_Bits(r.U32.Load()) }
func (r *IDR) Store(b IDR_Bits)            { r.U32.Store(uint32(b)) }

type IDR_Mask struct{ mmio.UM32 }

func (rm IDR_Mask) Load() IDR_Bits   { return IDR_Bits(rm.UM32.Load()) }
func (rm IDR_Mask) Store(b IDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) IDR_0() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_0)}}
}

func (p *GPIO_Periph) IDR_1() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_1)}}
}

func (p *GPIO_Periph) IDR_2() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_2)}}
}

func (p *GPIO_Periph) IDR_3() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_3)}}
}

func (p *GPIO_Periph) IDR_4() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_4)}}
}

func (p *GPIO_Periph) IDR_5() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_5)}}
}

func (p *GPIO_Periph) IDR_6() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_6)}}
}

func (p *GPIO_Periph) IDR_7() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_7)}}
}

func (p *GPIO_Periph) IDR_8() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_8)}}
}

func (p *GPIO_Periph) IDR_9() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_9)}}
}

func (p *GPIO_Periph) IDR_10() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_10)}}
}

func (p *GPIO_Periph) IDR_11() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_11)}}
}

func (p *GPIO_Periph) IDR_12() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_12)}}
}

func (p *GPIO_Periph) IDR_13() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_13)}}
}

func (p *GPIO_Periph) IDR_14() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_14)}}
}

func (p *GPIO_Periph) IDR_15() IDR_Mask {
	return IDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(IDR_15)}}
}

type ODR_Bits uint32

type ODR struct{ mmio.U32 }

func (r *ODR) Bits(mask ODR_Bits) ODR_Bits { return ODR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ODR) StoreBits(mask, b ODR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ODR) SetBits(mask ODR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ODR) ClearBits(mask ODR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ODR) Load() ODR_Bits              { return ODR_Bits(r.U32.Load()) }
func (r *ODR) Store(b ODR_Bits)            { r.U32.Store(uint32(b)) }

type ODR_Mask struct{ mmio.UM32 }

func (rm ODR_Mask) Load() ODR_Bits   { return ODR_Bits(rm.UM32.Load()) }
func (rm ODR_Mask) Store(b ODR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ODR_0() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_0)}}
}

func (p *GPIO_Periph) ODR_1() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_1)}}
}

func (p *GPIO_Periph) ODR_2() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_2)}}
}

func (p *GPIO_Periph) ODR_3() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_3)}}
}

func (p *GPIO_Periph) ODR_4() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_4)}}
}

func (p *GPIO_Periph) ODR_5() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_5)}}
}

func (p *GPIO_Periph) ODR_6() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_6)}}
}

func (p *GPIO_Periph) ODR_7() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_7)}}
}

func (p *GPIO_Periph) ODR_8() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_8)}}
}

func (p *GPIO_Periph) ODR_9() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_9)}}
}

func (p *GPIO_Periph) ODR_10() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_10)}}
}

func (p *GPIO_Periph) ODR_11() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_11)}}
}

func (p *GPIO_Periph) ODR_12() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_12)}}
}

func (p *GPIO_Periph) ODR_13() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_13)}}
}

func (p *GPIO_Periph) ODR_14() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_14)}}
}

func (p *GPIO_Periph) ODR_15() ODR_Mask {
	return ODR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ODR_15)}}
}

type BSRRL_Bits uint16

type BSRRL struct{ mmio.U16 }

func (r *BSRRL) Bits(mask BSRRL_Bits) BSRRL_Bits { return BSRRL_Bits(r.U16.Bits(uint16(mask))) }
func (r *BSRRL) StoreBits(mask, b BSRRL_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *BSRRL) SetBits(mask BSRRL_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *BSRRL) ClearBits(mask BSRRL_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *BSRRL) Load() BSRRL_Bits                { return BSRRL_Bits(r.U16.Load()) }
func (r *BSRRL) Store(b BSRRL_Bits)              { r.U16.Store(uint16(b)) }

type BSRRL_Mask struct{ mmio.UM16 }

func (rm BSRRL_Mask) Load() BSRRL_Bits   { return BSRRL_Bits(rm.UM16.Load()) }
func (rm BSRRL_Mask) Store(b BSRRL_Bits) { rm.UM16.Store(uint16(b)) }

type BSRRH_Bits uint16

type BSRRH struct{ mmio.U16 }

func (r *BSRRH) Bits(mask BSRRH_Bits) BSRRH_Bits { return BSRRH_Bits(r.U16.Bits(uint16(mask))) }
func (r *BSRRH) StoreBits(mask, b BSRRH_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *BSRRH) SetBits(mask BSRRH_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *BSRRH) ClearBits(mask BSRRH_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *BSRRH) Load() BSRRH_Bits                { return BSRRH_Bits(r.U16.Load()) }
func (r *BSRRH) Store(b BSRRH_Bits)              { r.U16.Store(uint16(b)) }

type BSRRH_Mask struct{ mmio.UM16 }

func (rm BSRRH_Mask) Load() BSRRH_Bits   { return BSRRH_Bits(rm.UM16.Load()) }
func (rm BSRRH_Mask) Store(b BSRRH_Bits) { rm.UM16.Store(uint16(b)) }

type LCKR_Bits uint32

type LCKR struct{ mmio.U32 }

func (r *LCKR) Bits(mask LCKR_Bits) LCKR_Bits { return LCKR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LCKR) StoreBits(mask, b LCKR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LCKR) SetBits(mask LCKR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LCKR) ClearBits(mask LCKR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LCKR) Load() LCKR_Bits               { return LCKR_Bits(r.U32.Load()) }
func (r *LCKR) Store(b LCKR_Bits)             { r.U32.Store(uint32(b)) }

type LCKR_Mask struct{ mmio.UM32 }

func (rm LCKR_Mask) Load() LCKR_Bits   { return LCKR_Bits(rm.UM32.Load()) }
func (rm LCKR_Mask) Store(b LCKR_Bits) { rm.UM32.Store(uint32(b)) }

type AFR_Bits uint32

type AFR struct{ mmio.U32 }

func (r *AFR) Bits(mask AFR_Bits) AFR_Bits { return AFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AFR) StoreBits(mask, b AFR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AFR) SetBits(mask AFR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *AFR) ClearBits(mask AFR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *AFR) Load() AFR_Bits              { return AFR_Bits(r.U32.Load()) }
func (r *AFR) Store(b AFR_Bits)            { r.U32.Store(uint32(b)) }

type AFR_Mask struct{ mmio.UM32 }

func (rm AFR_Mask) Load() AFR_Bits   { return AFR_Bits(rm.UM32.Load()) }
func (rm AFR_Mask) Store(b AFR_Bits) { rm.UM32.Store(uint32(b)) }
