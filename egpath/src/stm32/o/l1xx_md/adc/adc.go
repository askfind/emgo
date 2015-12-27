// Peripheral: ADC_Periph  Analog to Digital Converter.
// Instances:
//  ADC1  mmap.ADC1_BASE
// Registers:
//  0x00 32  SR    Status register.
//  0x04 32  CR1   Control register 1.
//  0x08 32  CR2   Control register 2.
//  0x0C 32  SMPR1 Sample time register 1.
//  0x10 32  SMPR2 Sample time register 2.
//  0x14 32  SMPR3 Sample time register 3.
//  0x18 32  JOFR1 Injected channel data offset register 1.
//  0x1C 32  JOFR2 Injected channel data offset register 2.
//  0x20 32  JOFR3 Injected channel data offset register 3.
//  0x24 32  JOFR4 Injected channel data offset register 4.
//  0x28 32  HTR   Watchdog higher threshold register.
//  0x2C 32  LTR   Watchdog lower threshold register.
//  0x30 32  SQR1  Regular sequence register 1.
//  0x34 32  SQR2  Regular sequence register 2.
//  0x38 32  SQR3  Regular sequence register 3.
//  0x3C 32  SQR4  Regular sequence register 4.
//  0x40 32  SQR5  Regular sequence register 5.
//  0x44 32  JSQR  Injected sequence register.
//  0x48 32  JDR1  Injected data register 1.
//  0x4C 32  JDR2  Injected data register 2.
//  0x50 32  JDR3  Injected data register 3.
//  0x54 32  JDR4  Injected data register 4.
//  0x58 32  DR    Regular data register.
//  0x5C 32  SMPR0 Sample time register 0.
// Import:
//  stm32/o/l1xx_md/mmap
package adc

const (
	AWD   SR_Bits = 0x01 << 0 //+ Analog watchdog flag.
	EOC   SR_Bits = 0x01 << 1 //+ End of conversion.
	JEOC  SR_Bits = 0x01 << 2 //+ Injected channel end of conversion.
	JSTRT SR_Bits = 0x01 << 3 //+ Injected channel Start flag.
	STRT  SR_Bits = 0x01 << 4 //+ Regular channel Start flag.
	OVR   SR_Bits = 0x01 << 5 //+ Overrun flag.
	ADONS SR_Bits = 0x01 << 6 //+ ADC ON status.
	RCNR  SR_Bits = 0x01 << 8 //+ Regular channel not ready flag.
	JCNR  SR_Bits = 0x01 << 9 //+ Injected channel not ready flag.
)

const (
	AWDCH     CR1_Bits = 0x1F << 0  //+ AWDCH[4:0] bits (Analog watchdog channel select bits).
	AWDCH_0   CR1_Bits = 0x01 << 0  //  Bit 0.
	AWDCH_1   CR1_Bits = 0x02 << 0  //  Bit 1.
	AWDCH_2   CR1_Bits = 0x04 << 0  //  Bit 2.
	AWDCH_3   CR1_Bits = 0x08 << 0  //  Bit 3.
	AWDCH_4   CR1_Bits = 0x10 << 0  //  Bit 4.
	EOCIE     CR1_Bits = 0x01 << 5  //+ Interrupt enable for EOC.
	AWDIE     CR1_Bits = 0x01 << 6  //+ Analog Watchdog interrupt enable.
	JEOCIE    CR1_Bits = 0x01 << 7  //+ Interrupt enable for injected channels.
	SCAN      CR1_Bits = 0x01 << 8  //+ Scan mode.
	AWDSGL    CR1_Bits = 0x01 << 9  //+ Enable the watchdog on a single channel in scan mode.
	JAUTO     CR1_Bits = 0x01 << 10 //+ Automatic injected group conversion.
	DISCEN    CR1_Bits = 0x01 << 11 //+ Discontinuous mode on regular channels.
	JDISCEN   CR1_Bits = 0x01 << 12 //+ Discontinuous mode on injected channels.
	DISCNUM   CR1_Bits = 0x07 << 13 //+ DISCNUM[2:0] bits (Discontinuous mode channel count).
	DISCNUM_0 CR1_Bits = 0x01 << 13 //  Bit 0.
	DISCNUM_1 CR1_Bits = 0x02 << 13 //  Bit 1.
	DISCNUM_2 CR1_Bits = 0x04 << 13 //  Bit 2.
	PDD       CR1_Bits = 0x01 << 16 //+ Power Down during Delay phase.
	PDI       CR1_Bits = 0x01 << 17 //+ Power Down during Idle phase.
	JAWDEN    CR1_Bits = 0x01 << 22 //+ Analog watchdog enable on injected channels.
	AWDEN     CR1_Bits = 0x01 << 23 //+ Analog watchdog enable on regular channels.
	RES       CR1_Bits = 0x03 << 24 //+ RES[1:0] bits (Resolution).
	RES_0     CR1_Bits = 0x01 << 24 //  Bit 0.
	RES_1     CR1_Bits = 0x02 << 24 //  Bit 1.
	OVRIE     CR1_Bits = 0x01 << 26 //+ Overrun interrupt enable.
)

const (
	ADON      CR2_Bits = 0x01 << 0  //+ A/D Converter ON / OFF.
	CONT      CR2_Bits = 0x01 << 1  //+ Continuous Conversion.
	CFG       CR2_Bits = 0x01 << 2  //+ ADC Configuration.
	DELS      CR2_Bits = 0x07 << 4  //+ DELS[2:0] bits (Delay selection).
	DELS_0    CR2_Bits = 0x01 << 4  //  Bit 0.
	DELS_1    CR2_Bits = 0x02 << 4  //  Bit 1.
	DELS_2    CR2_Bits = 0x04 << 4  //  Bit 2.
	DMA       CR2_Bits = 0x01 << 8  //+ Direct Memory access mode.
	DDS       CR2_Bits = 0x01 << 9  //+ DMA disable selection (Single ADC).
	EOCS      CR2_Bits = 0x01 << 10 //+ End of conversion selection.
	ALIGN     CR2_Bits = 0x01 << 11 //+ Data Alignment.
	JEXTSEL   CR2_Bits = 0x0F << 16 //+ JEXTSEL[3:0] bits (External event select for injected group).
	JEXTSEL_0 CR2_Bits = 0x01 << 16 //  Bit 0.
	JEXTSEL_1 CR2_Bits = 0x02 << 16 //  Bit 1.
	JEXTSEL_2 CR2_Bits = 0x04 << 16 //  Bit 2.
	JEXTSEL_3 CR2_Bits = 0x08 << 16 //  Bit 3.
	JEXTEN    CR2_Bits = 0x03 << 20 //+ JEXTEN[1:0] bits (External Trigger Conversion mode for injected channels).
	JEXTEN_0  CR2_Bits = 0x01 << 20 //  Bit 0.
	JEXTEN_1  CR2_Bits = 0x02 << 20 //  Bit 1.
	JSWSTART  CR2_Bits = 0x01 << 22 //+ Start Conversion of injected channels.
	EXTSEL    CR2_Bits = 0x0F << 24 //+ EXTSEL[3:0] bits (External Event Select for regular group).
	EXTSEL_0  CR2_Bits = 0x01 << 24 //  Bit 0.
	EXTSEL_1  CR2_Bits = 0x02 << 24 //  Bit 1.
	EXTSEL_2  CR2_Bits = 0x04 << 24 //  Bit 2.
	EXTSEL_3  CR2_Bits = 0x08 << 24 //  Bit 3.
	EXTEN     CR2_Bits = 0x03 << 28 //+ EXTEN[1:0] bits (External Trigger Conversion mode for regular channels).
	EXTEN_0   CR2_Bits = 0x01 << 28 //  Bit 0.
	EXTEN_1   CR2_Bits = 0x02 << 28 //  Bit 1.
	SWSTART   CR2_Bits = 0x01 << 30 //+ Start Conversion of regular channels.
)

const (
	SMP20   SMPR1_Bits = 0x07 << 0  //+ SMP20[2:0] bits (Channel 20 Sample time selection).
	SMP20_0 SMPR1_Bits = 0x01 << 0  //  Bit 0.
	SMP20_1 SMPR1_Bits = 0x02 << 0  //  Bit 1.
	SMP20_2 SMPR1_Bits = 0x04 << 0  //  Bit 2.
	SMP21   SMPR1_Bits = 0x07 << 3  //+ SMP21[2:0] bits (Channel 21 Sample time selection).
	SMP21_0 SMPR1_Bits = 0x01 << 3  //  Bit 0.
	SMP21_1 SMPR1_Bits = 0x02 << 3  //  Bit 1.
	SMP21_2 SMPR1_Bits = 0x04 << 3  //  Bit 2.
	SMP22   SMPR1_Bits = 0x07 << 6  //+ SMP22[2:0] bits (Channel 22 Sample time selection).
	SMP22_0 SMPR1_Bits = 0x01 << 6  //  Bit 0.
	SMP22_1 SMPR1_Bits = 0x02 << 6  //  Bit 1.
	SMP22_2 SMPR1_Bits = 0x04 << 6  //  Bit 2.
	SMP23   SMPR1_Bits = 0x07 << 9  //+ SMP23[2:0] bits (Channel 23 Sample time selection).
	SMP23_0 SMPR1_Bits = 0x01 << 9  //  Bit 0.
	SMP23_1 SMPR1_Bits = 0x02 << 9  //  Bit 1.
	SMP23_2 SMPR1_Bits = 0x04 << 9  //  Bit 2.
	SMP24   SMPR1_Bits = 0x07 << 12 //+ SMP24[2:0] bits (Channel 24 Sample time selection).
	SMP24_0 SMPR1_Bits = 0x01 << 12 //  Bit 0.
	SMP24_1 SMPR1_Bits = 0x02 << 12 //  Bit 1.
	SMP24_2 SMPR1_Bits = 0x04 << 12 //  Bit 2.
	SMP25   SMPR1_Bits = 0x07 << 15 //+ SMP25[2:0] bits (Channel 25 Sample time selection).
	SMP25_0 SMPR1_Bits = 0x01 << 15 //  Bit 0.
	SMP25_1 SMPR1_Bits = 0x02 << 15 //  Bit 1.
	SMP25_2 SMPR1_Bits = 0x04 << 15 //  Bit 2.
	SMP26   SMPR1_Bits = 0x07 << 18 //+ SMP26[2:0] bits (Channel 26 Sample time selection).
	SMP26_0 SMPR1_Bits = 0x01 << 18 //  Bit 0.
	SMP26_1 SMPR1_Bits = 0x02 << 18 //  Bit 1.
	SMP26_2 SMPR1_Bits = 0x04 << 18 //  Bit 2.
	SMP27   SMPR1_Bits = 0x07 << 21 //+ SMP27[2:0] bits (Channel 27 Sample time selection).
	SMP27_0 SMPR1_Bits = 0x01 << 21 //  Bit 0.
	SMP27_1 SMPR1_Bits = 0x02 << 21 //  Bit 1.
	SMP27_2 SMPR1_Bits = 0x04 << 21 //  Bit 2.
	SMP28   SMPR1_Bits = 0x07 << 24 //+ SMP28[2:0] bits (Channel 28 Sample time selection).
	SMP28_0 SMPR1_Bits = 0x01 << 24 //  Bit 0.
	SMP28_1 SMPR1_Bits = 0x02 << 24 //  Bit 1.
	SMP28_2 SMPR1_Bits = 0x04 << 24 //  Bit 2.
	SMP29   SMPR1_Bits = 0x07 << 27 //+ SMP29[2:0] bits (Channel 29 Sample time selection).
	SMP29_0 SMPR1_Bits = 0x01 << 27 //  Bit 0.
	SMP29_1 SMPR1_Bits = 0x02 << 27 //  Bit 1.
	SMP29_2 SMPR1_Bits = 0x04 << 27 //  Bit 2.
)

const (
	SMP10   SMPR2_Bits = 0x07 << 0  //+ SMP10[2:0] bits (Channel 10 Sample time selection).
	SMP10_0 SMPR2_Bits = 0x01 << 0  //  Bit 0.
	SMP10_1 SMPR2_Bits = 0x02 << 0  //  Bit 1.
	SMP10_2 SMPR2_Bits = 0x04 << 0  //  Bit 2.
	SMP11   SMPR2_Bits = 0x07 << 3  //+ SMP11[2:0] bits (Channel 11 Sample time selection).
	SMP11_0 SMPR2_Bits = 0x01 << 3  //  Bit 0.
	SMP11_1 SMPR2_Bits = 0x02 << 3  //  Bit 1.
	SMP11_2 SMPR2_Bits = 0x04 << 3  //  Bit 2.
	SMP12   SMPR2_Bits = 0x07 << 6  //+ SMP12[2:0] bits (Channel 12 Sample time selection).
	SMP12_0 SMPR2_Bits = 0x01 << 6  //  Bit 0.
	SMP12_1 SMPR2_Bits = 0x02 << 6  //  Bit 1.
	SMP12_2 SMPR2_Bits = 0x04 << 6  //  Bit 2.
	SMP13   SMPR2_Bits = 0x07 << 9  //+ SMP13[2:0] bits (Channel 13 Sample time selection).
	SMP13_0 SMPR2_Bits = 0x01 << 9  //  Bit 0.
	SMP13_1 SMPR2_Bits = 0x02 << 9  //  Bit 1.
	SMP13_2 SMPR2_Bits = 0x04 << 9  //  Bit 2.
	SMP14   SMPR2_Bits = 0x07 << 12 //+ SMP14[2:0] bits (Channel 14 Sample time selection).
	SMP14_0 SMPR2_Bits = 0x01 << 12 //  Bit 0.
	SMP14_1 SMPR2_Bits = 0x02 << 12 //  Bit 1.
	SMP14_2 SMPR2_Bits = 0x04 << 12 //  Bit 2.
	SMP15   SMPR2_Bits = 0x07 << 15 //+ SMP15[2:0] bits (Channel 5 Sample time selection).
	SMP15_0 SMPR2_Bits = 0x01 << 15 //  Bit 0.
	SMP15_1 SMPR2_Bits = 0x02 << 15 //  Bit 1.
	SMP15_2 SMPR2_Bits = 0x04 << 15 //  Bit 2.
	SMP16   SMPR2_Bits = 0x07 << 18 //+ SMP16[2:0] bits (Channel 16 Sample time selection).
	SMP16_0 SMPR2_Bits = 0x01 << 18 //  Bit 0.
	SMP16_1 SMPR2_Bits = 0x02 << 18 //  Bit 1.
	SMP16_2 SMPR2_Bits = 0x04 << 18 //  Bit 2.
	SMP17   SMPR2_Bits = 0x07 << 21 //+ SMP17[2:0] bits (Channel 17 Sample time selection).
	SMP17_0 SMPR2_Bits = 0x01 << 21 //  Bit 0.
	SMP17_1 SMPR2_Bits = 0x02 << 21 //  Bit 1.
	SMP17_2 SMPR2_Bits = 0x04 << 21 //  Bit 2.
	SMP18   SMPR2_Bits = 0x07 << 24 //+ SMP18[2:0] bits (Channel 18 Sample time selection).
	SMP18_0 SMPR2_Bits = 0x01 << 24 //  Bit 0.
	SMP18_1 SMPR2_Bits = 0x02 << 24 //  Bit 1.
	SMP18_2 SMPR2_Bits = 0x04 << 24 //  Bit 2.
	SMP19   SMPR2_Bits = 0x07 << 27 //+ SMP19[2:0] bits (Channel 19 Sample time selection).
	SMP19_0 SMPR2_Bits = 0x01 << 27 //  Bit 0.
	SMP19_1 SMPR2_Bits = 0x02 << 27 //  Bit 1.
	SMP19_2 SMPR2_Bits = 0x04 << 27 //  Bit 2.
)

const (
	SMP0    SMPR3_Bits = 0x07 << 0  //+ SMP0[2:0] bits (Channel 0 Sample time selection).
	SMP0_0  SMPR3_Bits = 0x01 << 0  //  Bit 0.
	SMP0_1  SMPR3_Bits = 0x02 << 0  //  Bit 1.
	SMP0_2  SMPR3_Bits = 0x04 << 0  //  Bit 2.
	SMP1    SMPR3_Bits = 0x07 << 3  //+ SMP1[2:0] bits (Channel 1 Sample time selection).
	SMP1_0  SMPR3_Bits = 0x01 << 3  //  Bit 0.
	SMP1_1  SMPR3_Bits = 0x02 << 3  //  Bit 1.
	SMP1_2  SMPR3_Bits = 0x04 << 3  //  Bit 2.
	SMP2    SMPR3_Bits = 0x07 << 6  //+ SMP2[2:0] bits (Channel 2 Sample time selection).
	SMP2_0  SMPR3_Bits = 0x01 << 6  //  Bit 0.
	SMP2_1  SMPR3_Bits = 0x02 << 6  //  Bit 1.
	SMP2_2  SMPR3_Bits = 0x04 << 6  //  Bit 2.
	SMP3    SMPR3_Bits = 0x07 << 9  //+ SMP3[2:0] bits (Channel 3 Sample time selection).
	SMP3_0  SMPR3_Bits = 0x01 << 9  //  Bit 0.
	SMP3_1  SMPR3_Bits = 0x02 << 9  //  Bit 1.
	SMP3_2  SMPR3_Bits = 0x04 << 9  //  Bit 2.
	SMP4    SMPR3_Bits = 0x07 << 12 //+ SMP4[2:0] bits (Channel 4 Sample time selection).
	SMP4_0  SMPR3_Bits = 0x01 << 12 //  Bit 0.
	SMP4_1  SMPR3_Bits = 0x02 << 12 //  Bit 1.
	SMP4_2  SMPR3_Bits = 0x04 << 12 //  Bit 2.
	SMP5    SMPR3_Bits = 0x07 << 15 //+ SMP5[2:0] bits (Channel 5 Sample time selection).
	SMP5_0  SMPR3_Bits = 0x01 << 15 //  Bit 0.
	SMP5_1  SMPR3_Bits = 0x02 << 15 //  Bit 1.
	SMP5_2  SMPR3_Bits = 0x04 << 15 //  Bit 2.
	SMP6    SMPR3_Bits = 0x07 << 18 //+ SMP6[2:0] bits (Channel 6 Sample time selection).
	SMP6_0  SMPR3_Bits = 0x01 << 18 //  Bit 0.
	SMP6_1  SMPR3_Bits = 0x02 << 18 //  Bit 1.
	SMP6_2  SMPR3_Bits = 0x04 << 18 //  Bit 2.
	SMP7    SMPR3_Bits = 0x07 << 21 //+ SMP7[2:0] bits (Channel 7 Sample time selection).
	SMP7_0  SMPR3_Bits = 0x01 << 21 //  Bit 0.
	SMP7_1  SMPR3_Bits = 0x02 << 21 //  Bit 1.
	SMP7_2  SMPR3_Bits = 0x04 << 21 //  Bit 2.
	SMP8    SMPR3_Bits = 0x07 << 24 //+ SMP8[2:0] bits (Channel 8 Sample time selection).
	SMP8_0  SMPR3_Bits = 0x01 << 24 //  Bit 0.
	SMP8_1  SMPR3_Bits = 0x02 << 24 //  Bit 1.
	SMP8_2  SMPR3_Bits = 0x04 << 24 //  Bit 2.
	SMP9    SMPR3_Bits = 0x07 << 27 //+ SMP9[2:0] bits (Channel 9 Sample time selection).
	SMP9_0  SMPR3_Bits = 0x01 << 27 //  Bit 0.
	SMP9_1  SMPR3_Bits = 0x02 << 27 //  Bit 1.
	SMP9_2  SMPR3_Bits = 0x04 << 27 //  Bit 2.
	SMP30   SMPR3_Bits = 0x07 << 0  //  SMP30[2:0] bits (Channel 30 Sample time selection).
	SMP30_0 SMPR3_Bits = 0x01 << 0  //  Bit 0.
	SMP30_1 SMPR3_Bits = 0x02 << 0  //  Bit 1.
	SMP30_2 SMPR3_Bits = 0x04 << 0  //  Bit 2.
	SMP31   SMPR3_Bits = 0x07 << 3  //  SMP31[2:0] bits (Channel 31 Sample time selection).
	SMP31_0 SMPR3_Bits = 0x01 << 3  //  Bit 0.
	SMP31_1 SMPR3_Bits = 0x02 << 3  //  Bit 1.
	SMP31_2 SMPR3_Bits = 0x04 << 3  //  Bit 2.
)

const (
	JOFFSET1 JOFR1_Bits = 0xFFF << 0 //+ Data offset for injected channel 1.
)

const (
	JOFFSET2 JOFR2_Bits = 0xFFF << 0 //+ Data offset for injected channel 2.
)

const (
	JOFFSET3 JOFR3_Bits = 0xFFF << 0 //+ Data offset for injected channel 3.
)

const (
	JOFFSET4 JOFR4_Bits = 0xFFF << 0 //+ Data offset for injected channel 4.
)

const (
	HT HTR_Bits = 0xFFF << 0 //+ Analog watchdog high threshold.
)

const (
	LT LTR_Bits = 0xFFF << 0 //+ Analog watchdog low threshold.
)

const (
	L      SQR1_Bits = 0x0F << 20 //+ L[3:0] bits (Regular channel sequence length).
	L_0    SQR1_Bits = 0x01 << 20 //  Bit 0.
	L_1    SQR1_Bits = 0x02 << 20 //  Bit 1.
	L_2    SQR1_Bits = 0x04 << 20 //  Bit 2.
	L_3    SQR1_Bits = 0x08 << 20 //  Bit 3.
	SQ28   SQR1_Bits = 0x1F << 15 //+ SQ28[4:0] bits (25th conversion in regular sequence).
	SQ28_0 SQR1_Bits = 0x01 << 15 //  Bit 0.
	SQ28_1 SQR1_Bits = 0x02 << 15 //  Bit 1.
	SQ28_2 SQR1_Bits = 0x04 << 15 //  Bit 2.
	SQ28_3 SQR1_Bits = 0x08 << 15 //  Bit 3.
	SQ28_4 SQR1_Bits = 0x10 << 15 //  Bit 4.
	SQ27   SQR1_Bits = 0x1F << 10 //+ SQ27[4:0] bits (27th conversion in regular sequence).
	SQ27_0 SQR1_Bits = 0x01 << 10 //  Bit 0.
	SQ27_1 SQR1_Bits = 0x02 << 10 //  Bit 1.
	SQ27_2 SQR1_Bits = 0x04 << 10 //  Bit 2.
	SQ27_3 SQR1_Bits = 0x08 << 10 //  Bit 3.
	SQ27_4 SQR1_Bits = 0x10 << 10 //  Bit 4.
	SQ26   SQR1_Bits = 0x1F << 5  //+ SQ26[4:0] bits (26th conversion in regular sequence).
	SQ26_0 SQR1_Bits = 0x01 << 5  //  Bit 0.
	SQ26_1 SQR1_Bits = 0x02 << 5  //  Bit 1.
	SQ26_2 SQR1_Bits = 0x04 << 5  //  Bit 2.
	SQ26_3 SQR1_Bits = 0x08 << 5  //  Bit 3.
	SQ26_4 SQR1_Bits = 0x10 << 5  //  Bit 4.
	SQ25   SQR1_Bits = 0x1F << 0  //+ SQ25[4:0] bits (25th conversion in regular sequence).
	SQ25_0 SQR1_Bits = 0x01 << 0  //  Bit 0.
	SQ25_1 SQR1_Bits = 0x02 << 0  //  Bit 1.
	SQ25_2 SQR1_Bits = 0x04 << 0  //  Bit 2.
	SQ25_3 SQR1_Bits = 0x08 << 0  //  Bit 3.
	SQ25_4 SQR1_Bits = 0x10 << 0  //  Bit 4.
)

const (
	SQ19   SQR2_Bits = 0x1F << 0  //+ SQ19[4:0] bits (19th conversion in regular sequence).
	SQ19_0 SQR2_Bits = 0x01 << 0  //  Bit 0.
	SQ19_1 SQR2_Bits = 0x02 << 0  //  Bit 1.
	SQ19_2 SQR2_Bits = 0x04 << 0  //  Bit 2.
	SQ19_3 SQR2_Bits = 0x08 << 0  //  Bit 3.
	SQ19_4 SQR2_Bits = 0x10 << 0  //  Bit 4.
	SQ20   SQR2_Bits = 0x1F << 5  //+ SQ20[4:0] bits (20th conversion in regular sequence).
	SQ20_0 SQR2_Bits = 0x01 << 5  //  Bit 0.
	SQ20_1 SQR2_Bits = 0x02 << 5  //  Bit 1.
	SQ20_2 SQR2_Bits = 0x04 << 5  //  Bit 2.
	SQ20_3 SQR2_Bits = 0x08 << 5  //  Bit 3.
	SQ20_4 SQR2_Bits = 0x10 << 5  //  Bit 4.
	SQ21   SQR2_Bits = 0x1F << 10 //+ SQ21[4:0] bits (21th conversion in regular sequence).
	SQ21_0 SQR2_Bits = 0x01 << 10 //  Bit 0.
	SQ21_1 SQR2_Bits = 0x02 << 10 //  Bit 1.
	SQ21_2 SQR2_Bits = 0x04 << 10 //  Bit 2.
	SQ21_3 SQR2_Bits = 0x08 << 10 //  Bit 3.
	SQ21_4 SQR2_Bits = 0x10 << 10 //  Bit 4.
	SQ22   SQR2_Bits = 0x1F << 15 //+ SQ22[4:0] bits (22th conversion in regular sequence).
	SQ22_0 SQR2_Bits = 0x01 << 15 //  Bit 0.
	SQ22_1 SQR2_Bits = 0x02 << 15 //  Bit 1.
	SQ22_2 SQR2_Bits = 0x04 << 15 //  Bit 2.
	SQ22_3 SQR2_Bits = 0x08 << 15 //  Bit 3.
	SQ22_4 SQR2_Bits = 0x10 << 15 //  Bit 4.
	SQ23   SQR2_Bits = 0x1F << 20 //+ SQ23[4:0] bits (23th conversion in regular sequence).
	SQ23_0 SQR2_Bits = 0x01 << 20 //  Bit 0.
	SQ23_1 SQR2_Bits = 0x02 << 20 //  Bit 1.
	SQ23_2 SQR2_Bits = 0x04 << 20 //  Bit 2.
	SQ23_3 SQR2_Bits = 0x08 << 20 //  Bit 3.
	SQ23_4 SQR2_Bits = 0x10 << 20 //  Bit 4.
	SQ24   SQR2_Bits = 0x1F << 25 //+ SQ24[4:0] bits (24th conversion in regular sequence).
	SQ24_0 SQR2_Bits = 0x01 << 25 //  Bit 0.
	SQ24_1 SQR2_Bits = 0x02 << 25 //  Bit 1.
	SQ24_2 SQR2_Bits = 0x04 << 25 //  Bit 2.
	SQ24_3 SQR2_Bits = 0x08 << 25 //  Bit 3.
	SQ24_4 SQR2_Bits = 0x10 << 25 //  Bit 4.
)

const (
	SQ13   SQR3_Bits = 0x1F << 0  //+ SQ13[4:0] bits (13th conversion in regular sequence).
	SQ13_0 SQR3_Bits = 0x01 << 0  //  Bit 0.
	SQ13_1 SQR3_Bits = 0x02 << 0  //  Bit 1.
	SQ13_2 SQR3_Bits = 0x04 << 0  //  Bit 2.
	SQ13_3 SQR3_Bits = 0x08 << 0  //  Bit 3.
	SQ13_4 SQR3_Bits = 0x10 << 0  //  Bit 4.
	SQ14   SQR3_Bits = 0x1F << 5  //+ SQ14[4:0] bits (14th conversion in regular sequence).
	SQ14_0 SQR3_Bits = 0x01 << 5  //  Bit 0.
	SQ14_1 SQR3_Bits = 0x02 << 5  //  Bit 1.
	SQ14_2 SQR3_Bits = 0x04 << 5  //  Bit 2.
	SQ14_3 SQR3_Bits = 0x08 << 5  //  Bit 3.
	SQ14_4 SQR3_Bits = 0x10 << 5  //  Bit 4.
	SQ15   SQR3_Bits = 0x1F << 10 //+ SQ15[4:0] bits (15th conversion in regular sequence).
	SQ15_0 SQR3_Bits = 0x01 << 10 //  Bit 0.
	SQ15_1 SQR3_Bits = 0x02 << 10 //  Bit 1.
	SQ15_2 SQR3_Bits = 0x04 << 10 //  Bit 2.
	SQ15_3 SQR3_Bits = 0x08 << 10 //  Bit 3.
	SQ15_4 SQR3_Bits = 0x10 << 10 //  Bit 4.
	SQ16   SQR3_Bits = 0x1F << 15 //+ SQ16[4:0] bits (16th conversion in regular sequence).
	SQ16_0 SQR3_Bits = 0x01 << 15 //  Bit 0.
	SQ16_1 SQR3_Bits = 0x02 << 15 //  Bit 1.
	SQ16_2 SQR3_Bits = 0x04 << 15 //  Bit 2.
	SQ16_3 SQR3_Bits = 0x08 << 15 //  Bit 3.
	SQ16_4 SQR3_Bits = 0x10 << 15 //  Bit 4.
	SQ17   SQR3_Bits = 0x1F << 20 //+ SQ17[4:0] bits (17th conversion in regular sequence).
	SQ17_0 SQR3_Bits = 0x01 << 20 //  Bit 0.
	SQ17_1 SQR3_Bits = 0x02 << 20 //  Bit 1.
	SQ17_2 SQR3_Bits = 0x04 << 20 //  Bit 2.
	SQ17_3 SQR3_Bits = 0x08 << 20 //  Bit 3.
	SQ17_4 SQR3_Bits = 0x10 << 20 //  Bit 4.
	SQ18   SQR3_Bits = 0x1F << 25 //+ SQ18[4:0] bits (18th conversion in regular sequence).
	SQ18_0 SQR3_Bits = 0x01 << 25 //  Bit 0.
	SQ18_1 SQR3_Bits = 0x02 << 25 //  Bit 1.
	SQ18_2 SQR3_Bits = 0x04 << 25 //  Bit 2.
	SQ18_3 SQR3_Bits = 0x08 << 25 //  Bit 3.
	SQ18_4 SQR3_Bits = 0x10 << 25 //  Bit 4.
)

const (
	SQ7    SQR4_Bits = 0x1F << 0  //+ SQ7[4:0] bits (7th conversion in regular sequence).
	SQ7_0  SQR4_Bits = 0x01 << 0  //  Bit 0.
	SQ7_1  SQR4_Bits = 0x02 << 0  //  Bit 1.
	SQ7_2  SQR4_Bits = 0x04 << 0  //  Bit 2.
	SQ7_3  SQR4_Bits = 0x08 << 0  //  Bit 3.
	SQ7_4  SQR4_Bits = 0x10 << 0  //  Bit 4.
	SQ8    SQR4_Bits = 0x1F << 5  //+ SQ8[4:0] bits (8th conversion in regular sequence).
	SQ8_0  SQR4_Bits = 0x01 << 5  //  Bit 0.
	SQ8_1  SQR4_Bits = 0x02 << 5  //  Bit 1.
	SQ8_2  SQR4_Bits = 0x04 << 5  //  Bit 2.
	SQ8_3  SQR4_Bits = 0x08 << 5  //  Bit 3.
	SQ8_4  SQR4_Bits = 0x10 << 5  //  Bit 4.
	SQ9    SQR4_Bits = 0x1F << 10 //+ SQ9[4:0] bits (9th conversion in regular sequence).
	SQ9_0  SQR4_Bits = 0x01 << 10 //  Bit 0.
	SQ9_1  SQR4_Bits = 0x02 << 10 //  Bit 1.
	SQ9_2  SQR4_Bits = 0x04 << 10 //  Bit 2.
	SQ9_3  SQR4_Bits = 0x08 << 10 //  Bit 3.
	SQ9_4  SQR4_Bits = 0x10 << 10 //  Bit 4.
	SQ10   SQR4_Bits = 0x1F << 15 //+ SQ10[4:0] bits (10th conversion in regular sequence).
	SQ10_0 SQR4_Bits = 0x01 << 15 //  Bit 0.
	SQ10_1 SQR4_Bits = 0x02 << 15 //  Bit 1.
	SQ10_2 SQR4_Bits = 0x04 << 15 //  Bit 2.
	SQ10_3 SQR4_Bits = 0x08 << 15 //  Bit 3.
	SQ10_4 SQR4_Bits = 0x10 << 15 //  Bit 4.
	SQ11   SQR4_Bits = 0x1F << 20 //+ SQ11[4:0] bits (11th conversion in regular sequence).
	SQ11_0 SQR4_Bits = 0x01 << 20 //  Bit 0.
	SQ11_1 SQR4_Bits = 0x02 << 20 //  Bit 1.
	SQ11_2 SQR4_Bits = 0x04 << 20 //  Bit 2.
	SQ11_3 SQR4_Bits = 0x08 << 20 //  Bit 3.
	SQ11_4 SQR4_Bits = 0x10 << 20 //  Bit 4.
	SQ12   SQR4_Bits = 0x1F << 25 //+ SQ12[4:0] bits (12th conversion in regular sequence).
	SQ12_0 SQR4_Bits = 0x01 << 25 //  Bit 0.
	SQ12_1 SQR4_Bits = 0x02 << 25 //  Bit 1.
	SQ12_2 SQR4_Bits = 0x04 << 25 //  Bit 2.
	SQ12_3 SQR4_Bits = 0x08 << 25 //  Bit 3.
	SQ12_4 SQR4_Bits = 0x10 << 25 //  Bit 4.
)

const (
	SQ1   SQR5_Bits = 0x1F << 0  //+ SQ1[4:0] bits (1st conversion in regular sequence).
	SQ1_0 SQR5_Bits = 0x01 << 0  //  Bit 0.
	SQ1_1 SQR5_Bits = 0x02 << 0  //  Bit 1.
	SQ1_2 SQR5_Bits = 0x04 << 0  //  Bit 2.
	SQ1_3 SQR5_Bits = 0x08 << 0  //  Bit 3.
	SQ1_4 SQR5_Bits = 0x10 << 0  //  Bit 4.
	SQ2   SQR5_Bits = 0x1F << 5  //+ SQ2[4:0] bits (2nd conversion in regular sequence).
	SQ2_0 SQR5_Bits = 0x01 << 5  //  Bit 0.
	SQ2_1 SQR5_Bits = 0x02 << 5  //  Bit 1.
	SQ2_2 SQR5_Bits = 0x04 << 5  //  Bit 2.
	SQ2_3 SQR5_Bits = 0x08 << 5  //  Bit 3.
	SQ2_4 SQR5_Bits = 0x10 << 5  //  Bit 4.
	SQ3   SQR5_Bits = 0x1F << 10 //+ SQ3[4:0] bits (3rd conversion in regular sequence).
	SQ3_0 SQR5_Bits = 0x01 << 10 //  Bit 0.
	SQ3_1 SQR5_Bits = 0x02 << 10 //  Bit 1.
	SQ3_2 SQR5_Bits = 0x04 << 10 //  Bit 2.
	SQ3_3 SQR5_Bits = 0x08 << 10 //  Bit 3.
	SQ3_4 SQR5_Bits = 0x10 << 10 //  Bit 4.
	SQ4   SQR5_Bits = 0x1F << 15 //+ SQ4[4:0] bits (4th conversion in regular sequence).
	SQ4_0 SQR5_Bits = 0x01 << 15 //  Bit 0.
	SQ4_1 SQR5_Bits = 0x02 << 15 //  Bit 1.
	SQ4_2 SQR5_Bits = 0x04 << 15 //  Bit 2.
	SQ4_3 SQR5_Bits = 0x08 << 15 //  Bit 3.
	SQ4_4 SQR5_Bits = 0x10 << 15 //  Bit 4.
	SQ5   SQR5_Bits = 0x1F << 20 //+ SQ5[4:0] bits (5th conversion in regular sequence).
	SQ5_0 SQR5_Bits = 0x01 << 20 //  Bit 0.
	SQ5_1 SQR5_Bits = 0x02 << 20 //  Bit 1.
	SQ5_2 SQR5_Bits = 0x04 << 20 //  Bit 2.
	SQ5_3 SQR5_Bits = 0x08 << 20 //  Bit 3.
	SQ5_4 SQR5_Bits = 0x10 << 20 //  Bit 4.
	SQ6   SQR5_Bits = 0x1F << 25 //+ SQ6[4:0] bits (6th conversion in regular sequence).
	SQ6_0 SQR5_Bits = 0x01 << 25 //  Bit 0.
	SQ6_1 SQR5_Bits = 0x02 << 25 //  Bit 1.
	SQ6_2 SQR5_Bits = 0x04 << 25 //  Bit 2.
	SQ6_3 SQR5_Bits = 0x08 << 25 //  Bit 3.
	SQ6_4 SQR5_Bits = 0x10 << 25 //  Bit 4.
)

const (
	JSQ1   JSQR_Bits = 0x1F << 0  //+ JSQ1[4:0] bits (1st conversion in injected sequence).
	JSQ1_0 JSQR_Bits = 0x01 << 0  //  Bit 0.
	JSQ1_1 JSQR_Bits = 0x02 << 0  //  Bit 1.
	JSQ1_2 JSQR_Bits = 0x04 << 0  //  Bit 2.
	JSQ1_3 JSQR_Bits = 0x08 << 0  //  Bit 3.
	JSQ1_4 JSQR_Bits = 0x10 << 0  //  Bit 4.
	JSQ2   JSQR_Bits = 0x1F << 5  //+ JSQ2[4:0] bits (2nd conversion in injected sequence).
	JSQ2_0 JSQR_Bits = 0x01 << 5  //  Bit 0.
	JSQ2_1 JSQR_Bits = 0x02 << 5  //  Bit 1.
	JSQ2_2 JSQR_Bits = 0x04 << 5  //  Bit 2.
	JSQ2_3 JSQR_Bits = 0x08 << 5  //  Bit 3.
	JSQ2_4 JSQR_Bits = 0x10 << 5  //  Bit 4.
	JSQ3   JSQR_Bits = 0x1F << 10 //+ JSQ3[4:0] bits (3rd conversion in injected sequence).
	JSQ3_0 JSQR_Bits = 0x01 << 10 //  Bit 0.
	JSQ3_1 JSQR_Bits = 0x02 << 10 //  Bit 1.
	JSQ3_2 JSQR_Bits = 0x04 << 10 //  Bit 2.
	JSQ3_3 JSQR_Bits = 0x08 << 10 //  Bit 3.
	JSQ3_4 JSQR_Bits = 0x10 << 10 //  Bit 4.
	JSQ4   JSQR_Bits = 0x1F << 15 //+ JSQ4[4:0] bits (4th conversion in injected sequence).
	JSQ4_0 JSQR_Bits = 0x01 << 15 //  Bit 0.
	JSQ4_1 JSQR_Bits = 0x02 << 15 //  Bit 1.
	JSQ4_2 JSQR_Bits = 0x04 << 15 //  Bit 2.
	JSQ4_3 JSQR_Bits = 0x08 << 15 //  Bit 3.
	JSQ4_4 JSQR_Bits = 0x10 << 15 //  Bit 4.
	JL     JSQR_Bits = 0x03 << 20 //+ JL[1:0] bits (Injected Sequence length).
	JL_0   JSQR_Bits = 0x01 << 20 //  Bit 0.
	JL_1   JSQR_Bits = 0x02 << 20 //  Bit 1.
)

const (
	JDATA JDR1_Bits = 0xFFFF << 0 //+ Injected data.
)

const (
	JDATA JDR2_Bits = 0xFFFF << 0 //+ Injected data.
)

const (
	JDATA JDR3_Bits = 0xFFFF << 0 //+ Injected data.
)

const (
	JDATA JDR4_Bits = 0xFFFF << 0 //+ Injected data.
)

const (
	DATA DR_Bits = 0xFFFF << 0 //+ Regular data.
)