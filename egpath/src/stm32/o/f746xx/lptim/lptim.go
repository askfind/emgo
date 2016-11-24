// Peripheral: LPTIM_Periph  LPTIMIMER.
// Instances:
//  LPTIM1  mmap.LPTIM1_BASE
// Registers:
//  0x00 32  ISR  Interrupt and Status register.
//  0x04 32  ICR  Interrupt Clear register.
//  0x08 32  IER  Interrupt Enable register.
//  0x0C 32  CFGR Configuration register.
//  0x10 32  CR   Control register.
//  0x14 32  CMP  Compare register.
//  0x18 32  ARR  Autoreload register.
//  0x1C 32  CNT  Counter register.
// Import:
//  stm32/o/f746xx/mmap
package lptim

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CMPM    ISR_Bits = 0x01 << 0 //+ Compare match.
	ARRM    ISR_Bits = 0x01 << 1 //+ Autoreload match.
	EXTTRIG ISR_Bits = 0x01 << 2 //+ External trigger edge event.
	CMPOK   ISR_Bits = 0x01 << 3 //+ Compare register update OK.
	ARROK   ISR_Bits = 0x01 << 4 //+ Autoreload register update OK.
	UP      ISR_Bits = 0x01 << 5 //+ Counter direction change down to up.
	DOWN    ISR_Bits = 0x01 << 6 //+ Counter direction change up to down.
)

const (
	CMPMn    = 0
	ARRMn    = 1
	EXTTRIGn = 2
	CMPOKn   = 3
	ARROKn   = 4
	UPn      = 5
	DOWNn    = 6
)

const (
	CMPMCF    ICR_Bits = 0x01 << 0 //+ Compare match Clear Flag.
	ARRMCF    ICR_Bits = 0x01 << 1 //+ Autoreload match Clear Flag.
	EXTTRIGCF ICR_Bits = 0x01 << 2 //+ External trigger edge event Clear Flag.
	CMPOKCF   ICR_Bits = 0x01 << 3 //+ Compare register update OK Clear Flag.
	ARROKCF   ICR_Bits = 0x01 << 4 //+ Autoreload register update OK Clear Flag.
	UPCF      ICR_Bits = 0x01 << 5 //+ Counter direction change down to up Clear Flag.
	DOWNCF    ICR_Bits = 0x01 << 6 //+ Counter direction change up to down Clear Flag.
)

const (
	CMPMCFn    = 0
	ARRMCFn    = 1
	EXTTRIGCFn = 2
	CMPOKCFn   = 3
	ARROKCFn   = 4
	UPCFn      = 5
	DOWNCFn    = 6
)

const (
	CMPMIE    IER_Bits = 0x01 << 0 //+ Compare match Interrupt Enable.
	ARRMIE    IER_Bits = 0x01 << 1 //+ Autoreload match Interrupt Enable.
	EXTTRIGIE IER_Bits = 0x01 << 2 //+ External trigger edge event Interrupt Enable.
	CMPOKIE   IER_Bits = 0x01 << 3 //+ Compare register update OK Interrupt Enable.
	ARROKIE   IER_Bits = 0x01 << 4 //+ Autoreload register update OK Interrupt Enable.
	UPIE      IER_Bits = 0x01 << 5 //+ Counter direction change down to up Interrupt Enable.
	DOWNIE    IER_Bits = 0x01 << 6 //+ Counter direction change up to down Interrupt Enable.
)

const (
	CMPMIEn    = 0
	ARRMIEn    = 1
	EXTTRIGIEn = 2
	CMPOKIEn   = 3
	ARROKIEn   = 4
	UPIEn      = 5
	DOWNIEn    = 6
)

const (
	CKSEL     CFGR_Bits = 0x01 << 0  //+ Clock selector.
	CKPOL     CFGR_Bits = 0x03 << 1  //+ CKPOL[1:0] bits (Clock polarity).
	CKPOL_0   CFGR_Bits = 0x01 << 1  //  Bit 0.
	CKPOL_1   CFGR_Bits = 0x02 << 1  //  Bit 1.
	CKFLT     CFGR_Bits = 0x03 << 3  //+ CKFLT[1:0] bits (Configurable digital filter for external clock).
	CKFLT_0   CFGR_Bits = 0x01 << 3  //  Bit 0.
	CKFLT_1   CFGR_Bits = 0x02 << 3  //  Bit 1.
	TRGFLT    CFGR_Bits = 0x03 << 6  //+ TRGFLT[1:0] bits (Configurable digital filter for trigger).
	TRGFLT_0  CFGR_Bits = 0x01 << 6  //  Bit 0.
	TRGFLT_1  CFGR_Bits = 0x02 << 6  //  Bit 1.
	PRESC     CFGR_Bits = 0x07 << 9  //+ PRESC[2:0] bits (Clock prescaler).
	PRESC_0   CFGR_Bits = 0x01 << 9  //  Bit 0.
	PRESC_1   CFGR_Bits = 0x02 << 9  //  Bit 1.
	PRESC_2   CFGR_Bits = 0x04 << 9  //  Bit 2.
	TRIGSEL   CFGR_Bits = 0x07 << 13 //+ TRIGSEL[2:0]] bits (Trigger selector).
	TRIGSEL_0 CFGR_Bits = 0x01 << 13 //  Bit 0.
	TRIGSEL_1 CFGR_Bits = 0x02 << 13 //  Bit 1.
	TRIGSEL_2 CFGR_Bits = 0x04 << 13 //  Bit 2.
	TRIGEN    CFGR_Bits = 0x03 << 17 //+ TRIGEN[1:0] bits (Trigger enable and polarity).
	TRIGEN_0  CFGR_Bits = 0x01 << 17 //  Bit 0.
	TRIGEN_1  CFGR_Bits = 0x02 << 17 //  Bit 1.
	TIMOUT    CFGR_Bits = 0x01 << 19 //+ Timout enable.
	WAVE      CFGR_Bits = 0x01 << 20 //+ Waveform shape.
	WAVPOL    CFGR_Bits = 0x01 << 21 //+ Waveform shape polarity.
	PRELOAD   CFGR_Bits = 0x01 << 22 //+ Reg update mode.
	COUNTMODE CFGR_Bits = 0x01 << 23 //+ Counter mode enable.
	ENC       CFGR_Bits = 0x01 << 24 //+ Encoder mode enable.
)

const (
	CKSELn     = 0
	CKPOLn     = 1
	CKFLTn     = 3
	TRGFLTn    = 6
	PRESCn     = 9
	TRIGSELn   = 13
	TRIGENn    = 17
	TIMOUTn    = 19
	WAVEn      = 20
	WAVPOLn    = 21
	PRELOADn   = 22
	COUNTMODEn = 23
	ENCn       = 24
)

const (
	ENABLE  CR_Bits = 0x01 << 0 //+ LPTIMer enable.
	SNGSTRT CR_Bits = 0x01 << 1 //+ Timer start in single mode.
	CNTSTRT CR_Bits = 0x01 << 2 //+ Timer start in continuous mode.
)

const (
	ENABLEn  = 0
	SNGSTRTn = 1
	CNTSTRTn = 2
)
