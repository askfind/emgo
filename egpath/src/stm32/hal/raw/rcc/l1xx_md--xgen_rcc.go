// +build l1xx_md

package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type RCC_Periph struct {
	CR        CR
	ICSCR     ICSCR
	CFGR      CFGR
	CIR       CIR
	AHBRSTR   AHBRSTR
	APB2RSTR  APB2RSTR
	APB1RSTR  APB1RSTR
	AHBENR    AHBENR
	APB2ENR   APB2ENR
	APB1ENR   APB1ENR
	AHBLPENR  AHBLPENR
	APB2LPENR APB2LPENR
	APB1LPENR APB1LPENR
	CSR       CSR
}

func (p *RCC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var RCC = (*RCC_Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

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

func (p *RCC_Periph) HSION() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *RCC_Periph) HSIRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *RCC_Periph) MSION() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MSION)}}
}

func (p *RCC_Periph) MSIRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MSIRDY)}}
}

func (p *RCC_Periph) HSEON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *RCC_Periph) HSERDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *RCC_Periph) HSEBYP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *RCC_Periph) PLLON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *RCC_Periph) PLLRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
}

func (p *RCC_Periph) CSSON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *RCC_Periph) RTCPRE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(RTCPRE)}}
}

type ICSCR_Bits uint32

func (b ICSCR_Bits) Field(mask ICSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICSCR_Bits) J(v int) ICSCR_Bits {
	return ICSCR_Bits(bits.Make32(v, uint32(mask)))
}

type ICSCR struct{ mmio.U32 }

func (r *ICSCR) Bits(mask ICSCR_Bits) ICSCR_Bits { return ICSCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICSCR) StoreBits(mask, b ICSCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICSCR) SetBits(mask ICSCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ICSCR) ClearBits(mask ICSCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ICSCR) Load() ICSCR_Bits                { return ICSCR_Bits(r.U32.Load()) }
func (r *ICSCR) Store(b ICSCR_Bits)              { r.U32.Store(uint32(b)) }

type ICSCR_Mask struct{ mmio.UM32 }

func (rm ICSCR_Mask) Load() ICSCR_Bits   { return ICSCR_Bits(rm.UM32.Load()) }
func (rm ICSCR_Mask) Store(b ICSCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) HSICAL() ICSCR_Mask {
	return ICSCR_Mask{mmio.UM32{&p.ICSCR.U32, uint32(HSICAL)}}
}

func (p *RCC_Periph) HSITRIM() ICSCR_Mask {
	return ICSCR_Mask{mmio.UM32{&p.ICSCR.U32, uint32(HSITRIM)}}
}

func (p *RCC_Periph) MSIRANGE() ICSCR_Mask {
	return ICSCR_Mask{mmio.UM32{&p.ICSCR.U32, uint32(MSIRANGE)}}
}

func (p *RCC_Periph) MSICAL() ICSCR_Mask {
	return ICSCR_Mask{mmio.UM32{&p.ICSCR.U32, uint32(MSICAL)}}
}

func (p *RCC_Periph) MSITRIM() ICSCR_Mask {
	return ICSCR_Mask{mmio.UM32{&p.ICSCR.U32, uint32(MSITRIM)}}
}

type CFGR_Bits uint32

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR_Bits) J(v int) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR struct{ mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.U32.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.U32.Store(uint32(b)) }

type CFGR_Mask struct{ mmio.UM32 }

func (rm CFGR_Mask) Load() CFGR_Bits   { return CFGR_Bits(rm.UM32.Load()) }
func (rm CFGR_Mask) Store(b CFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SW() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *RCC_Periph) SWS() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *RCC_Periph) HPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *RCC_Periph) PPRE1() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *RCC_Periph) PPRE2() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *RCC_Periph) PLLSRC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC)}}
}

func (p *RCC_Periph) PLLMUL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLMUL)}}
}

func (p *RCC_Periph) PLLDIV() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLDIV)}}
}

func (p *RCC_Periph) MCOSEL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCOSEL)}}
}

func (p *RCC_Periph) MCOPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCOPRE)}}
}

type CIR_Bits uint32

func (b CIR_Bits) Field(mask CIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR_Bits) J(v int) CIR_Bits {
	return CIR_Bits(bits.Make32(v, uint32(mask)))
}

type CIR struct{ mmio.U32 }

func (r *CIR) Bits(mask CIR_Bits) CIR_Bits { return CIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CIR) StoreBits(mask, b CIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CIR) SetBits(mask CIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CIR) ClearBits(mask CIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CIR) Load() CIR_Bits              { return CIR_Bits(r.U32.Load()) }
func (r *CIR) Store(b CIR_Bits)            { r.U32.Store(uint32(b)) }

type CIR_Mask struct{ mmio.UM32 }

func (rm CIR_Mask) Load() CIR_Bits   { return CIR_Bits(rm.UM32.Load()) }
func (rm CIR_Mask) Store(b CIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *RCC_Periph) LSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *RCC_Periph) HSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *RCC_Periph) HSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *RCC_Periph) PLLRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *RCC_Periph) MSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(MSIRDYF)}}
}

func (p *RCC_Periph) LSECSS() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSECSS)}}
}

func (p *RCC_Periph) CSSF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *RCC_Periph) LSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *RCC_Periph) LSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *RCC_Periph) HSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *RCC_Periph) HSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *RCC_Periph) PLLRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *RCC_Periph) MSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(MSIRDYIE)}}
}

func (p *RCC_Periph) LSECSSIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSECSSIE)}}
}

func (p *RCC_Periph) LSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *RCC_Periph) LSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *RCC_Periph) HSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *RCC_Periph) HSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *RCC_Periph) PLLRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *RCC_Periph) MSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(MSIRDYC)}}
}

func (p *RCC_Periph) LSECSSC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSECSSC)}}
}

func (p *RCC_Periph) CSSC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type AHBRSTR_Bits uint32

func (b AHBRSTR_Bits) Field(mask AHBRSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBRSTR_Bits) J(v int) AHBRSTR_Bits {
	return AHBRSTR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBRSTR struct{ mmio.U32 }

func (r *AHBRSTR) Bits(mask AHBRSTR_Bits) AHBRSTR_Bits { return AHBRSTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AHBRSTR) StoreBits(mask, b AHBRSTR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBRSTR) SetBits(mask AHBRSTR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *AHBRSTR) ClearBits(mask AHBRSTR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *AHBRSTR) Load() AHBRSTR_Bits                  { return AHBRSTR_Bits(r.U32.Load()) }
func (r *AHBRSTR) Store(b AHBRSTR_Bits)                { r.U32.Store(uint32(b)) }

type AHBRSTR_Mask struct{ mmio.UM32 }

func (rm AHBRSTR_Mask) Load() AHBRSTR_Bits   { return AHBRSTR_Bits(rm.UM32.Load()) }
func (rm AHBRSTR_Mask) Store(b AHBRSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOARST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOARST)}}
}

func (p *RCC_Periph) GPIOBRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOBRST)}}
}

func (p *RCC_Periph) GPIOCRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOCRST)}}
}

func (p *RCC_Periph) GPIODRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIODRST)}}
}

func (p *RCC_Periph) GPIOERST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOERST)}}
}

func (p *RCC_Periph) GPIOHRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOHRST)}}
}

func (p *RCC_Periph) GPIOFRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOFRST)}}
}

func (p *RCC_Periph) GPIOGRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOGRST)}}
}

func (p *RCC_Periph) CRCRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(CRCRST)}}
}

func (p *RCC_Periph) FLITFRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(FLITFRST)}}
}

func (p *RCC_Periph) DMA1RST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(DMA1RST)}}
}

func (p *RCC_Periph) DMA2RST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(DMA2RST)}}
}

func (p *RCC_Periph) AESRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(AESRST)}}
}

func (p *RCC_Periph) FSMCRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(FSMCRST)}}
}

type APB2RSTR_Bits uint32

func (b APB2RSTR_Bits) Field(mask APB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR_Bits) J(v int) APB2RSTR_Bits {
	return APB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2RSTR struct{ mmio.U32 }

func (r *APB2RSTR) Bits(mask APB2RSTR_Bits) APB2RSTR_Bits {
	return APB2RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB2RSTR) StoreBits(mask, b APB2RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2RSTR) SetBits(mask APB2RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB2RSTR) ClearBits(mask APB2RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB2RSTR) Load() APB2RSTR_Bits             { return APB2RSTR_Bits(r.U32.Load()) }
func (r *APB2RSTR) Store(b APB2RSTR_Bits)           { r.U32.Store(uint32(b)) }

type APB2RSTR_Mask struct{ mmio.UM32 }

func (rm APB2RSTR_Mask) Load() APB2RSTR_Bits   { return APB2RSTR_Bits(rm.UM32.Load()) }
func (rm APB2RSTR_Mask) Store(b APB2RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SYSCFGRST)}}
}

func (p *RCC_Periph) TIM9RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM9RST)}}
}

func (p *RCC_Periph) TIM10RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM10RST)}}
}

func (p *RCC_Periph) TIM11RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM11RST)}}
}

func (p *RCC_Periph) ADC1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(ADC1RST)}}
}

func (p *RCC_Periph) SDIORST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SDIORST)}}
}

func (p *RCC_Periph) SPI1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *RCC_Periph) USART1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

type APB1RSTR_Bits uint32

func (b APB1RSTR_Bits) Field(mask APB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR_Bits) J(v int) APB1RSTR_Bits {
	return APB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1RSTR struct{ mmio.U32 }

func (r *APB1RSTR) Bits(mask APB1RSTR_Bits) APB1RSTR_Bits {
	return APB1RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB1RSTR) StoreBits(mask, b APB1RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1RSTR) SetBits(mask APB1RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB1RSTR) ClearBits(mask APB1RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB1RSTR) Load() APB1RSTR_Bits             { return APB1RSTR_Bits(r.U32.Load()) }
func (r *APB1RSTR) Store(b APB1RSTR_Bits)           { r.U32.Store(uint32(b)) }

type APB1RSTR_Mask struct{ mmio.UM32 }

func (rm APB1RSTR_Mask) Load() APB1RSTR_Bits   { return APB1RSTR_Bits(rm.UM32.Load()) }
func (rm APB1RSTR_Mask) Store(b APB1RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *RCC_Periph) TIM3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *RCC_Periph) TIM4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *RCC_Periph) TIM5RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM5RST)}}
}

func (p *RCC_Periph) TIM6RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *RCC_Periph) TIM7RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *RCC_Periph) LCDRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(LCDRST)}}
}

func (p *RCC_Periph) WWDGRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *RCC_Periph) SPI2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *RCC_Periph) SPI3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *RCC_Periph) USART2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART2RST)}}
}

func (p *RCC_Periph) USART3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *RCC_Periph) UART4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART4RST)}}
}

func (p *RCC_Periph) UART5RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART5RST)}}
}

func (p *RCC_Periph) I2C1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *RCC_Periph) I2C2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *RCC_Periph) USBRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USBRST)}}
}

func (p *RCC_Periph) PWRRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

func (p *RCC_Periph) DACRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(DACRST)}}
}

func (p *RCC_Periph) COMPRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(COMPRST)}}
}

type AHBENR_Bits uint32

func (b AHBENR_Bits) Field(mask AHBENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR_Bits) J(v int) AHBENR_Bits {
	return AHBENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBENR struct{ mmio.U32 }

func (r *AHBENR) Bits(mask AHBENR_Bits) AHBENR_Bits { return AHBENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AHBENR) StoreBits(mask, b AHBENR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBENR) SetBits(mask AHBENR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *AHBENR) ClearBits(mask AHBENR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *AHBENR) Load() AHBENR_Bits                 { return AHBENR_Bits(r.U32.Load()) }
func (r *AHBENR) Store(b AHBENR_Bits)               { r.U32.Store(uint32(b)) }

type AHBENR_Mask struct{ mmio.UM32 }

func (rm AHBENR_Mask) Load() AHBENR_Bits   { return AHBENR_Bits(rm.UM32.Load()) }
func (rm AHBENR_Mask) Store(b AHBENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOAEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOAEN)}}
}

func (p *RCC_Periph) GPIOBEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOBEN)}}
}

func (p *RCC_Periph) GPIOCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOCEN)}}
}

func (p *RCC_Periph) GPIODEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIODEN)}}
}

func (p *RCC_Periph) GPIOEEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOEEN)}}
}

func (p *RCC_Periph) GPIOHEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOHEN)}}
}

func (p *RCC_Periph) GPIOFEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOFEN)}}
}

func (p *RCC_Periph) GPIOGEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOGEN)}}
}

func (p *RCC_Periph) CRCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(CRCEN)}}
}

func (p *RCC_Periph) FLITFEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FLITFEN)}}
}

func (p *RCC_Periph) DMA1EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA1EN)}}
}

func (p *RCC_Periph) DMA2EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA2EN)}}
}

func (p *RCC_Periph) AESEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(AESEN)}}
}

func (p *RCC_Periph) FSMCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FSMCEN)}}
}

type APB2ENR_Bits uint32

func (b APB2ENR_Bits) Field(mask APB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR_Bits) J(v int) APB2ENR_Bits {
	return APB2ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2ENR struct{ mmio.U32 }

func (r *APB2ENR) Bits(mask APB2ENR_Bits) APB2ENR_Bits { return APB2ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB2ENR) StoreBits(mask, b APB2ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2ENR) SetBits(mask APB2ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB2ENR) ClearBits(mask APB2ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB2ENR) Load() APB2ENR_Bits                  { return APB2ENR_Bits(r.U32.Load()) }
func (r *APB2ENR) Store(b APB2ENR_Bits)                { r.U32.Store(uint32(b)) }

type APB2ENR_Mask struct{ mmio.UM32 }

func (rm APB2ENR_Mask) Load() APB2ENR_Bits   { return APB2ENR_Bits(rm.UM32.Load()) }
func (rm APB2ENR_Mask) Store(b APB2ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SYSCFGEN)}}
}

func (p *RCC_Periph) TIM9EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM9EN)}}
}

func (p *RCC_Periph) TIM10EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM10EN)}}
}

func (p *RCC_Periph) TIM11EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM11EN)}}
}

func (p *RCC_Periph) ADC1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(ADC1EN)}}
}

func (p *RCC_Periph) SDIOEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SDIOEN)}}
}

func (p *RCC_Periph) SPI1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *RCC_Periph) USART1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

type APB1ENR_Bits uint32

func (b APB1ENR_Bits) Field(mask APB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR_Bits) J(v int) APB1ENR_Bits {
	return APB1ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1ENR struct{ mmio.U32 }

func (r *APB1ENR) Bits(mask APB1ENR_Bits) APB1ENR_Bits { return APB1ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB1ENR) StoreBits(mask, b APB1ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1ENR) SetBits(mask APB1ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB1ENR) ClearBits(mask APB1ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB1ENR) Load() APB1ENR_Bits                  { return APB1ENR_Bits(r.U32.Load()) }
func (r *APB1ENR) Store(b APB1ENR_Bits)                { r.U32.Store(uint32(b)) }

type APB1ENR_Mask struct{ mmio.UM32 }

func (rm APB1ENR_Mask) Load() APB1ENR_Bits   { return APB1ENR_Bits(rm.UM32.Load()) }
func (rm APB1ENR_Mask) Store(b APB1ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *RCC_Periph) TIM3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *RCC_Periph) TIM4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *RCC_Periph) TIM5EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM5EN)}}
}

func (p *RCC_Periph) TIM6EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *RCC_Periph) TIM7EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *RCC_Periph) LCDEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(LCDEN)}}
}

func (p *RCC_Periph) WWDGEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *RCC_Periph) SPI2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *RCC_Periph) SPI3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *RCC_Periph) USART2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *RCC_Periph) USART3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *RCC_Periph) UART4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART4EN)}}
}

func (p *RCC_Periph) UART5EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART5EN)}}
}

func (p *RCC_Periph) I2C1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *RCC_Periph) I2C2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *RCC_Periph) USBEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USBEN)}}
}

func (p *RCC_Periph) PWREN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

func (p *RCC_Periph) DACEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(DACEN)}}
}

func (p *RCC_Periph) COMPEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(COMPEN)}}
}

type AHBLPENR_Bits uint32

func (b AHBLPENR_Bits) Field(mask AHBLPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBLPENR_Bits) J(v int) AHBLPENR_Bits {
	return AHBLPENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBLPENR struct{ mmio.U32 }

func (r *AHBLPENR) Bits(mask AHBLPENR_Bits) AHBLPENR_Bits {
	return AHBLPENR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *AHBLPENR) StoreBits(mask, b AHBLPENR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBLPENR) SetBits(mask AHBLPENR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *AHBLPENR) ClearBits(mask AHBLPENR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *AHBLPENR) Load() AHBLPENR_Bits             { return AHBLPENR_Bits(r.U32.Load()) }
func (r *AHBLPENR) Store(b AHBLPENR_Bits)           { r.U32.Store(uint32(b)) }

type AHBLPENR_Mask struct{ mmio.UM32 }

func (rm AHBLPENR_Mask) Load() AHBLPENR_Bits   { return AHBLPENR_Bits(rm.UM32.Load()) }
func (rm AHBLPENR_Mask) Store(b AHBLPENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) GPIOALPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOALPEN)}}
}

func (p *RCC_Periph) GPIOBLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOBLPEN)}}
}

func (p *RCC_Periph) GPIOCLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOCLPEN)}}
}

func (p *RCC_Periph) GPIODLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIODLPEN)}}
}

func (p *RCC_Periph) GPIOELPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOELPEN)}}
}

func (p *RCC_Periph) GPIOHLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOHLPEN)}}
}

func (p *RCC_Periph) GPIOFLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOFLPEN)}}
}

func (p *RCC_Periph) GPIOGLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(GPIOGLPEN)}}
}

func (p *RCC_Periph) CRCLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(CRCLPEN)}}
}

func (p *RCC_Periph) FLITFLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(FLITFLPEN)}}
}

func (p *RCC_Periph) SRAMLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(SRAMLPEN)}}
}

func (p *RCC_Periph) DMA1LPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(DMA1LPEN)}}
}

func (p *RCC_Periph) DMA2LPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(DMA2LPEN)}}
}

func (p *RCC_Periph) AESLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(AESLPEN)}}
}

func (p *RCC_Periph) FSMCLPEN() AHBLPENR_Mask {
	return AHBLPENR_Mask{mmio.UM32{&p.AHBLPENR.U32, uint32(FSMCLPEN)}}
}

type APB2LPENR_Bits uint32

func (b APB2LPENR_Bits) Field(mask APB2LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2LPENR_Bits) J(v int) APB2LPENR_Bits {
	return APB2LPENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2LPENR struct{ mmio.U32 }

func (r *APB2LPENR) Bits(mask APB2LPENR_Bits) APB2LPENR_Bits {
	return APB2LPENR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB2LPENR) StoreBits(mask, b APB2LPENR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2LPENR) SetBits(mask APB2LPENR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB2LPENR) ClearBits(mask APB2LPENR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB2LPENR) Load() APB2LPENR_Bits             { return APB2LPENR_Bits(r.U32.Load()) }
func (r *APB2LPENR) Store(b APB2LPENR_Bits)           { r.U32.Store(uint32(b)) }

type APB2LPENR_Mask struct{ mmio.UM32 }

func (rm APB2LPENR_Mask) Load() APB2LPENR_Bits   { return APB2LPENR_Bits(rm.UM32.Load()) }
func (rm APB2LPENR_Mask) Store(b APB2LPENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGLPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(SYSCFGLPEN)}}
}

func (p *RCC_Periph) TIM9LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM9LPEN)}}
}

func (p *RCC_Periph) TIM10LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM10LPEN)}}
}

func (p *RCC_Periph) TIM11LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(TIM11LPEN)}}
}

func (p *RCC_Periph) ADC1LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(ADC1LPEN)}}
}

func (p *RCC_Periph) SDIOLPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(SDIOLPEN)}}
}

func (p *RCC_Periph) SPI1LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(SPI1LPEN)}}
}

func (p *RCC_Periph) USART1LPEN() APB2LPENR_Mask {
	return APB2LPENR_Mask{mmio.UM32{&p.APB2LPENR.U32, uint32(USART1LPEN)}}
}

type APB1LPENR_Bits uint32

func (b APB1LPENR_Bits) Field(mask APB1LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1LPENR_Bits) J(v int) APB1LPENR_Bits {
	return APB1LPENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1LPENR struct{ mmio.U32 }

func (r *APB1LPENR) Bits(mask APB1LPENR_Bits) APB1LPENR_Bits {
	return APB1LPENR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB1LPENR) StoreBits(mask, b APB1LPENR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1LPENR) SetBits(mask APB1LPENR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB1LPENR) ClearBits(mask APB1LPENR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB1LPENR) Load() APB1LPENR_Bits             { return APB1LPENR_Bits(r.U32.Load()) }
func (r *APB1LPENR) Store(b APB1LPENR_Bits)           { r.U32.Store(uint32(b)) }

type APB1LPENR_Mask struct{ mmio.UM32 }

func (rm APB1LPENR_Mask) Load() APB1LPENR_Bits   { return APB1LPENR_Bits(rm.UM32.Load()) }
func (rm APB1LPENR_Mask) Store(b APB1LPENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM2LPEN)}}
}

func (p *RCC_Periph) TIM3LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM3LPEN)}}
}

func (p *RCC_Periph) TIM4LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM4LPEN)}}
}

func (p *RCC_Periph) TIM5LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM5LPEN)}}
}

func (p *RCC_Periph) TIM6LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM6LPEN)}}
}

func (p *RCC_Periph) TIM7LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(TIM7LPEN)}}
}

func (p *RCC_Periph) LCDLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(LCDLPEN)}}
}

func (p *RCC_Periph) WWDGLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(WWDGLPEN)}}
}

func (p *RCC_Periph) SPI2LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI2LPEN)}}
}

func (p *RCC_Periph) SPI3LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(SPI3LPEN)}}
}

func (p *RCC_Periph) USART2LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(USART2LPEN)}}
}

func (p *RCC_Periph) USART3LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(USART3LPEN)}}
}

func (p *RCC_Periph) UART4LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(UART4LPEN)}}
}

func (p *RCC_Periph) UART5LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(UART5LPEN)}}
}

func (p *RCC_Periph) I2C1LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C1LPEN)}}
}

func (p *RCC_Periph) I2C2LPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(I2C2LPEN)}}
}

func (p *RCC_Periph) USBLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(USBLPEN)}}
}

func (p *RCC_Periph) PWRLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(PWRLPEN)}}
}

func (p *RCC_Periph) DACLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(DACLPEN)}}
}

func (p *RCC_Periph) COMPLPEN() APB1LPENR_Mask {
	return APB1LPENR_Mask{mmio.UM32{&p.APB1LPENR.U32, uint32(COMPLPEN)}}
}

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSION() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *RCC_Periph) LSIRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *RCC_Periph) LSEON() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSEON)}}
}

func (p *RCC_Periph) LSERDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSERDY)}}
}

func (p *RCC_Periph) LSEBYP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSEBYP)}}
}

func (p *RCC_Periph) LSECSSON() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSECSSON)}}
}

func (p *RCC_Periph) LSECSSD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSECSSD)}}
}

func (p *RCC_Periph) RTCSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RTCSEL)}}
}

func (p *RCC_Periph) RTCEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RTCEN)}}
}

func (p *RCC_Periph) RTCRST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RTCRST)}}
}

func (p *RCC_Periph) RMVF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *RCC_Periph) OBLRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OBLRSTF)}}
}

func (p *RCC_Periph) PINRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PINRSTF)}}
}

func (p *RCC_Periph) PORRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *RCC_Periph) SFTRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *RCC_Periph) IWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(IWDGRSTF)}}
}

func (p *RCC_Periph) WWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *RCC_Periph) LPWRRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}
