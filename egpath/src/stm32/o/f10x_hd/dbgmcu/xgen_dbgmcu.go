package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type DBGMCU_Periph struct {
	IDCODE IDCODE
	CR     CR
}

func (p *DBGMCU_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var DBGMCU = (*DBGMCU_Periph)(unsafe.Pointer(uintptr(mmap.DBGMCU_BASE)))

type IDCODE_Bits uint32

func (b IDCODE_Bits) Field(mask IDCODE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDCODE_Bits) J(v int) IDCODE_Bits {
	return IDCODE_Bits(bits.Make32(v, uint32(mask)))
}

type IDCODE struct{ mmio.U32 }

func (r *IDCODE) Bits(mask IDCODE_Bits) IDCODE_Bits { return IDCODE_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDCODE) StoreBits(mask, b IDCODE_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDCODE) SetBits(mask IDCODE_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *IDCODE) ClearBits(mask IDCODE_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *IDCODE) Load() IDCODE_Bits                 { return IDCODE_Bits(r.U32.Load()) }
func (r *IDCODE) Store(b IDCODE_Bits)               { r.U32.Store(uint32(b)) }

type IDCODE_Mask struct{ mmio.UM32 }

func (rm IDCODE_Mask) Load() IDCODE_Bits   { return IDCODE_Bits(rm.UM32.Load()) }
func (rm IDCODE_Mask) Store(b IDCODE_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DBGMCU_Periph) DEV_ID() IDCODE_Mask {
	return IDCODE_Mask{mmio.UM32{&p.IDCODE.U32, uint32(DEV_ID)}}
}

func (p *DBGMCU_Periph) REV_ID() IDCODE_Mask {
	return IDCODE_Mask{mmio.UM32{&p.IDCODE.U32, uint32(REV_ID)}}
}

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

func (p *DBGMCU_Periph) DBG_SLEEP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_SLEEP)}}
}

func (p *DBGMCU_Periph) DBG_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_STANDBY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_STANDBY)}}
}

func (p *DBGMCU_Periph) TRACE_IOEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TRACE_IOEN)}}
}

func (p *DBGMCU_Periph) TRACE_MODE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TRACE_MODE)}}
}

func (p *DBGMCU_Periph) DBG_IWDG_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_IWDG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_WWDG_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_WWDG_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM1_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM1_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM2_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM2_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM3_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM3_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM4_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM4_STOP)}}
}

func (p *DBGMCU_Periph) DBG_CAN1_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_CAN1_STOP)}}
}

func (p *DBGMCU_Periph) DBG_I2C1_SMBUS_TIMEOUT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_I2C1_SMBUS_TIMEOUT)}}
}

func (p *DBGMCU_Periph) DBG_I2C2_SMBUS_TIMEOUT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_I2C2_SMBUS_TIMEOUT)}}
}

func (p *DBGMCU_Periph) DBG_TIM8_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM8_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM5_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM5_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM6_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM6_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM7_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM7_STOP)}}
}

func (p *DBGMCU_Periph) DBG_CAN2_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_CAN2_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM15_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM15_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM16_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM16_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM17_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM17_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM12_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM12_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM13_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM13_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM14_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM14_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM9_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM9_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM10_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM10_STOP)}}
}

func (p *DBGMCU_Periph) DBG_TIM11_STOP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DBG_TIM11_STOP)}}
}
