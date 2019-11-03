package talib

// #cgo LDFLAGS: -lta_lib -lm
// #include "ta-lib/ta_libc.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	n, err := C.TA_Initialize()
	if n != 0 {
		panic(fmt.Sprintf("ta-lib status is %d %s", n, err))
	}
}

const MAType_SMA = 0
const MAType_EMA = 1
const MAType_WMA = 2
const MAType_DEMA = 3
const MAType_TEMA = 4
const MAType_TRIMA = 5
const MAType_KAMA = 6
const MAType_MAMA = 7
const MAType_T3 = 8

/*Acos - Vector Trigonometric ACos

Input = double

Output = double

*/
func Acos(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ACOS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Ad - Chaikin A/D Line

Input = High, Low, Close, Volume

Output = double

*/
func Ad(high, low, close, volume []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_AD(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Add - Vector Arithmetic Add

Input = double, double

Output = double

*/
func Add(real0, real1 []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_ADD(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*AdOsc - Chaikin A/D Oscillator

Input = High, Low, Close, Volume

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

*/
func AdOsc(high, low, close, volume []float64, fastPeriod, slowPeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_ADOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), C.int(fastPeriod), C.int(slowPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Adx - Average Directional Movement Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Adx(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_ADX(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Adxr - Average Directional Movement Index Rating

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Adxr(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_ADXR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Apo - Absolute Price Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInMAType:

Type of Moving Average

*/
func Apo(real []float64, fastPeriod, slowPeriod, mAType int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_APO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*AroOn - Aroon

Input = High, Low

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func AroOn(high, low []float64, timePeriod int, outAroonDown []float64, outAroonUp []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outAroonDown == nil {
		outAroonDown = make([]float64, len(high))
	}
	if outAroonUp == nil {
		outAroonUp = make([]float64, len(high))
	}
	C.TA_AROON(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outAroonDown[0])), (*C.double)(unsafe.Pointer(&outAroonUp[0])))
	return outAroonDown[:outNBElement], outAroonUp[:outNBElement], int(outBegIdx)
}

/*AroOnOsc - Aroon Oscillator

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func AroOnOsc(high, low []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_AROONOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Asin - Vector Trigonometric ASin

Input = double

Output = double

*/
func Asin(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ASIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Atan - Vector Trigonometric ATan

Input = double

Output = double

*/
func Atan(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ATAN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Atr - Average True Range

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Atr(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_ATR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*AvgPrice - Average Price

Input = Open, High, Low, Close

Output = double

*/
func AvgPrice(open, high, low, close []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(open))
	}
	C.TA_AVGPRICE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*BBands - Bollinger Bands

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInNbDevUp:(From TA_REAL_MIN to TA_REAL_MAX)

Deviation multiplier for upper band

optInNbDevDn:(From TA_REAL_MIN to TA_REAL_MAX)

Deviation multiplier for lower band

optInMAType:

Type of Moving Average

*/
func BBands(real []float64, timePeriod int, nbDevUp, nbDevDn float64, mAType int, outRealUpperBand []float64, outRealMiddleBand []float64, outRealLowerBand []float64) ([]float64, []float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outRealUpperBand == nil {
		outRealUpperBand = make([]float64, len(real))
	}
	if outRealMiddleBand == nil {
		outRealMiddleBand = make([]float64, len(real))
	}
	if outRealLowerBand == nil {
		outRealLowerBand = make([]float64, len(real))
	}
	C.TA_BBANDS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDevUp), C.double(nbDevDn), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outRealUpperBand[0])), (*C.double)(unsafe.Pointer(&outRealMiddleBand[0])), (*C.double)(unsafe.Pointer(&outRealLowerBand[0])))
	return outRealUpperBand[:outNBElement], outRealMiddleBand[:outNBElement], outRealLowerBand[:outNBElement], int(outBegIdx)
}

/*Beta - Beta

Input = double, double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Beta(real0, real1 []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_BETA(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Bop - Balance Of Power

Input = Open, High, Low, Close

Output = double

*/
func Bop(open, high, low, close []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(open))
	}
	C.TA_BOP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Cci - Commodity Channel Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Cci(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_CCI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Cdl2Crows - Two Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdl2Crows(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL2CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3BlackCrows - Three Black Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdl3BlackCrows(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3BLACKCROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3Inside - Three Inside Up/Down

Input = Open, High, Low, Close

Output = int

*/
func Cdl3Inside(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3INSIDE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3LineStrike - Three-Line Strike

Input = Open, High, Low, Close

Output = int

*/
func Cdl3LineStrike(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3LINESTRIKE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3Outside - Three Outside Up/Down

Input = Open, High, Low, Close

Output = int

*/
func Cdl3Outside(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3OUTSIDE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3StarsinSouth - Three Stars In The South

Input = Open, High, Low, Close

Output = int

*/
func Cdl3StarsinSouth(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3STARSINSOUTH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Cdl3WhiteSoldiers - Three Advancing White Soldiers

Input = Open, High, Low, Close

Output = int

*/
func Cdl3WhiteSoldiers(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDL3WHITESOLDIERS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlAbandonedBaby - Abandoned Baby

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlAbandonedBaby(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLABANDONEDBABY(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlAdvanceBlock - Advance Block

Input = Open, High, Low, Close

Output = int

*/
func CdlAdvanceBlock(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLADVANCEBLOCK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlBelthold - Belt-hold

Input = Open, High, Low, Close

Output = int

*/
func CdlBelthold(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLBELTHOLD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlBreakaway - Breakaway

Input = Open, High, Low, Close

Output = int

*/
func CdlBreakaway(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLBREAKAWAY(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlClosingMarubozu - Closing Marubozu

Input = Open, High, Low, Close

Output = int

*/
func CdlClosingMarubozu(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLCLOSINGMARUBOZU(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlConcealBabySwall - Concealing Baby Swallow

Input = Open, High, Low, Close

Output = int

*/
func CdlConcealBabySwall(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLCONCEALBABYSWALL(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlCounterattack - Counterattack

Input = Open, High, Low, Close

Output = int

*/
func CdlCounterattack(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLCOUNTERATTACK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlDarkCloudCover - Dark Cloud Cover

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlDarkCloudCover(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLDARKCLOUDCOVER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlDoji - Doji

Input = Open, High, Low, Close

Output = int

*/
func CdlDoji(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlDojiStar - Doji Star

Input = Open, High, Low, Close

Output = int

*/
func CdlDojiStar(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlDragonflyDoji - Dragonfly Doji

Input = Open, High, Low, Close

Output = int

*/
func CdlDragonflyDoji(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLDRAGONFLYDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlEngulfing - Engulfing Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlEngulfing(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLENGULFING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlEveningDojiStar - Evening Doji Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlEveningDojiStar(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLEVENINGDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlEveningStar - Evening Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlEveningStar(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLEVENINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlGapSidesideWhite - Up/Down-gap side-by-side white lines

Input = Open, High, Low, Close

Output = int

*/
func CdlGapSidesideWhite(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLGAPSIDESIDEWHITE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlGravestoneDoji - Gravestone Doji

Input = Open, High, Low, Close

Output = int

*/
func CdlGravestoneDoji(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLGRAVESTONEDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHammer - Hammer

Input = Open, High, Low, Close

Output = int

*/
func CdlHammer(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHAMMER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHangingMan - Hanging Man

Input = Open, High, Low, Close

Output = int

*/
func CdlHangingMan(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHANGINGMAN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHarami - Harami Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlHarami(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHARAMI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHaramiCross - Harami Cross Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlHaramiCross(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHARAMICROSS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHighWave - High-Wave Candle

Input = Open, High, Low, Close

Output = int

*/
func CdlHighWave(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHIGHWAVE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHikkake - Hikkake Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlHikkake(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHIKKAKE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHikkakeMod - Modified Hikkake Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlHikkakeMod(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHIKKAKEMOD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlHomingPigeon - Homing Pigeon

Input = Open, High, Low, Close

Output = int

*/
func CdlHomingPigeon(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLHOMINGPIGEON(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlIdentical3Crows - Identical Three Crows

Input = Open, High, Low, Close

Output = int

*/
func CdlIdentical3Crows(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLIDENTICAL3CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlInNeck - In-Neck Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlInNeck(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLINNECK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlInvertedHammer - Inverted Hammer

Input = Open, High, Low, Close

Output = int

*/
func CdlInvertedHammer(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLINVERTEDHAMMER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlKicking - Kicking

Input = Open, High, Low, Close

Output = int

*/
func CdlKicking(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLKICKING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlKickingByLength - Kicking - bull/bear determined by the longer marubozu

Input = Open, High, Low, Close

Output = int

*/
func CdlKickingByLength(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLKICKINGBYLENGTH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlLadderBottom - Ladder Bottom

Input = Open, High, Low, Close

Output = int

*/
func CdlLadderBottom(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLLADDERBOTTOM(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlLongLeggedDoji - Long Legged Doji

Input = Open, High, Low, Close

Output = int

*/
func CdlLongLeggedDoji(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLLONGLEGGEDDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlLongLine - Long Line Candle

Input = Open, High, Low, Close

Output = int

*/
func CdlLongLine(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLLONGLINE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlMarubozu - Marubozu

Input = Open, High, Low, Close

Output = int

*/
func CdlMarubozu(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLMARUBOZU(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlMatchingLow - Matching Low

Input = Open, High, Low, Close

Output = int

*/
func CdlMatchingLow(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLMATCHINGLOW(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlMatHold - Mat Hold

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlMatHold(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLMATHOLD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlMorningDojiStar - Morning Doji Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlMorningDojiStar(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLMORNINGDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlMorningStar - Morning Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func CdlMorningStar(open, high, low, close []float64, penetration float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLMORNINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlOnNeck - On-Neck Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlOnNeck(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLONNECK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlPiercing - Piercing Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlPiercing(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLPIERCING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlRickshawMan - Rickshaw Man

Input = Open, High, Low, Close

Output = int

*/
func CdlRickshawMan(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLRICKSHAWMAN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlRiseFall3Methods - Rising/Falling Three Methods

Input = Open, High, Low, Close

Output = int

*/
func CdlRiseFall3Methods(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLRISEFALL3METHODS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlSeparatingLines - Separating Lines

Input = Open, High, Low, Close

Output = int

*/
func CdlSeparatingLines(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSEPARATINGLINES(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlShootingStar - Shooting Star

Input = Open, High, Low, Close

Output = int

*/
func CdlShootingStar(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSHOOTINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlShortLine - Short Line Candle

Input = Open, High, Low, Close

Output = int

*/
func CdlShortLine(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSHORTLINE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlSpinningTop - Spinning Top

Input = Open, High, Low, Close

Output = int

*/
func CdlSpinningTop(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSPINNINGTOP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlStalledPattern - Stalled Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlStalledPattern(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSTALLEDPATTERN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlStickSandwich - Stick Sandwich

Input = Open, High, Low, Close

Output = int

*/
func CdlStickSandwich(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLSTICKSANDWICH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlTakuri - Takuri (Dragonfly Doji with very long lower shadow)

Input = Open, High, Low, Close

Output = int

*/
func CdlTakuri(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLTAKURI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlTasukiGap - Tasuki Gap

Input = Open, High, Low, Close

Output = int

*/
func CdlTasukiGap(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLTASUKIGAP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlThrusting - Thrusting Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlThrusting(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLTHRUSTING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlTristar - Tristar Pattern

Input = Open, High, Low, Close

Output = int

*/
func CdlTristar(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLTRISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlUnique3River - Unique 3 River

Input = Open, High, Low, Close

Output = int

*/
func CdlUnique3River(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLUNIQUE3RIVER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlUpsideGap2Crows - Upside Gap Two Crows

Input = Open, High, Low, Close

Output = int

*/
func CdlUpsideGap2Crows(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLUPSIDEGAP2CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*CdlxSideGap3Methods - Upside/Downside Gap Three Methods

Input = Open, High, Low, Close

Output = int

*/
func CdlxSideGap3Methods(open, high, low, close []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(open))
	}
	C.TA_CDLXSIDEGAP3METHODS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Ceil - Vector Ceil

Input = double

Output = double

*/
func Ceil(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_CEIL(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Cmo - Chande Momentum Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Cmo(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_CMO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Correl - Pearson's Correlation Coefficient (r)

Input = double, double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Correl(real0, real1 []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_CORREL(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Cos - Vector Trigonometric Cos

Input = double

Output = double

*/
func Cos(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_COS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Cosh - Vector Trigonometric Cosh

Input = double

Output = double

*/
func Cosh(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_COSH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Dema - Double Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Dema(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_DEMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Div - Vector Arithmetic Div

Input = double, double

Output = double

*/
func Div(real0, real1 []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_DIV(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Dx - Directional Movement Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Dx(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_DX(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Ema - Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Ema(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_EMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Exp - Vector Arithmetic Exp

Input = double

Output = double

*/
func Exp(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_EXP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Floor - Vector Floor

Input = double

Output = double

*/
func Floor(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_FLOOR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*HtDcPeriod - Hilbert Transform - Dominant Cycle Period

Input = double

Output = double

*/
func HtDcPeriod(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_HT_DCPERIOD(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*HtDcPhase - Hilbert Transform - Dominant Cycle Phase

Input = double

Output = double

*/
func HtDcPhase(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_HT_DCPHASE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*HtPhasor - Hilbert Transform - Phasor Components

Input = double

Output = double, double

*/
func HtPhasor(real []float64, outInPhase []float64, outQuadrature []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInPhase == nil {
		outInPhase = make([]float64, len(real))
	}
	if outQuadrature == nil {
		outQuadrature = make([]float64, len(real))
	}
	C.TA_HT_PHASOR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outInPhase[0])), (*C.double)(unsafe.Pointer(&outQuadrature[0])))
	return outInPhase[:outNBElement], outQuadrature[:outNBElement], int(outBegIdx)
}

/*HtSine - Hilbert Transform - SineWave

Input = double

Output = double, double

*/
func HtSine(real []float64, outSine []float64, outLeadSine []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outSine == nil {
		outSine = make([]float64, len(real))
	}
	if outLeadSine == nil {
		outLeadSine = make([]float64, len(real))
	}
	C.TA_HT_SINE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outSine[0])), (*C.double)(unsafe.Pointer(&outLeadSine[0])))
	return outSine[:outNBElement], outLeadSine[:outNBElement], int(outBegIdx)
}

/*HtTrendLine - Hilbert Transform - Instantaneous Trendline

Input = double

Output = double

*/
func HtTrendLine(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_HT_TRENDLINE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*HtTrendMode - Hilbert Transform - Trend vs Cycle Mode

Input = double

Output = int

*/
func HtTrendMode(real []float64, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(real))
	}
	C.TA_HT_TRENDMODE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*Kama - Kaufman Adaptive Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Kama(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_KAMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*LinearReg - Linear Regression

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearReg(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LINEARREG(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*LinearRegAngle - Linear Regression Angle

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearRegAngle(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LINEARREG_ANGLE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*LinearRegIntercept - Linear Regression Intercept

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearRegIntercept(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LINEARREG_INTERCEPT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*LinearRegSlope - Linear Regression Slope

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearRegSlope(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LINEARREG_SLOPE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Ln - Vector Log Natural

Input = double

Output = double

*/
func Ln(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Log10 - Vector Log10

Input = double

Output = double

*/
func Log10(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_LOG10(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Ma - Moving average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

optInMaType:

Type of Moving Average

*/
func Ma(real []float64, timePeriod, mAType int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Macd - Moving Average Convergence/Divergence

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

*/
func Macd(real []float64, fastPeriod, slowPeriod, signalPeriod int, outMACD []float64, outMACDSignal []float64, outMACDHist []float64) ([]float64, []float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMACD == nil {
		outMACD = make([]float64, len(real))
	}
	if outMACDSignal == nil {
		outMACDSignal = make([]float64, len(real))
	}
	if outMACDHist == nil {
		outMACDHist = make([]float64, len(real))
	}
	C.TA_MACD(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.int(signalPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD[:outNBElement], outMACDSignal[:outNBElement], outMACDHist[:outNBElement], int(outBegIdx)
}

/*MacdExt - MACD with controllable MA type

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInFastMAType:

Type of Moving Average for fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInSlowMAType:

Type of Moving Average for slow MA

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

optInSignalMAType:

Type of Moving Average for signal line

*/
func MacdExt(real []float64, fastPeriod, fastMAType, slowPeriod, slowMAType, signalPeriod, signalMAType int, outMACD []float64, outMACDSignal []float64, outMACDHist []float64) ([]float64, []float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMACD == nil {
		outMACD = make([]float64, len(real))
	}
	if outMACDSignal == nil {
		outMACDSignal = make([]float64, len(real))
	}
	if outMACDHist == nil {
		outMACDHist = make([]float64, len(real))
	}
	C.TA_MACDEXT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.TA_MAType(fastMAType), C.int(slowPeriod), C.TA_MAType(slowMAType), C.int(signalPeriod), C.TA_MAType(signalMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD[:outNBElement], outMACDSignal[:outNBElement], outMACDHist[:outNBElement], int(outBegIdx)
}

/*MacdFix - Moving Average Convergence/Divergence Fix 12/26

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

*/
func MacdFix(real []float64, signalPeriod int, outMACD []float64, outMACDSignal []float64, outMACDHist []float64) ([]float64, []float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMACD == nil {
		outMACD = make([]float64, len(real))
	}
	if outMACDSignal == nil {
		outMACDSignal = make([]float64, len(real))
	}
	if outMACDHist == nil {
		outMACDHist = make([]float64, len(real))
	}
	C.TA_MACDFIX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(signalPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD[:outNBElement], outMACDSignal[:outNBElement], outMACDHist[:outNBElement], int(outBegIdx)
}

/*Mama - MESA Adaptive Moving Average

Input = double

Output = double, double

Optional Parameters

-------------------

optInFastLimit:(From 0.01 to 0.99)

Upper limit use in the adaptive algorithm

optInSlowLimit:(From 0.01 to 0.99)

Lower limit use in the adaptive algorithm

*/
func Mama(real []float64, fastLimit, slowLimit float64, outMAMA []float64, outFAMA []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMAMA == nil {
		outMAMA = make([]float64, len(real))
	}
	if outFAMA == nil {
		outFAMA = make([]float64, len(real))
	}
	C.TA_MAMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.double(fastLimit), C.double(slowLimit), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMAMA[0])), (*C.double)(unsafe.Pointer(&outFAMA[0])))
	return outMAMA[:outNBElement], outFAMA[:outNBElement], int(outBegIdx)
}

/*Mavp - Moving average with variable period

Input = double, double

Output = double

Optional Parameters

-------------------

optInMinPeriod:(From 2 to 100000)

Value less than minimum will be changed to Minimum period

optInMaxPeriod:(From 2 to 100000)

Value higher than maximum will be changed to Maximum period

optInMAType:

Type of Moving Average

*/
func Mavp(real, periods []float64, minPeriod, maxPeriod, mAType int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MAVP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&periods[0])), C.int(minPeriod), C.int(maxPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Max - Highest value over a specified period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Max(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MAX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*MaxIndex - Index of highest value over a specified period

Input = double

Output = int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MaxIndex(real []float64, timePeriod int, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(real))
	}
	C.TA_MAXINDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*MedPrice - Median Price

Input = High, Low

Output = double

*/
func MedPrice(high, low []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_MEDPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Mfi - Money Flow Index

Input = High, Low, Close, Volume

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Mfi(high, low, close, volume []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_MFI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*MidPoint - MidPoint over period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MidPoint(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MIDPOINT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*MidPrice - Midpoint Price over period

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MidPrice(high, low []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_MIDPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Min - Lowest value over a specified period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Min(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*MinIndex - Index of lowest value over a specified period

Input = double

Output = int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MinIndex(real []float64, timePeriod int, outInteger []int) ([]int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outInteger == nil {
		outInteger = make([]int, len(real))
	}
	C.TA_MININDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger[:outNBElement], int(outBegIdx)
}

/*MinMax - Lowest and highest values over a specified period

Input = double

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MinMax(real []float64, timePeriod int, outMin []float64, outMax []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMin == nil {
		outMin = make([]float64, len(real))
	}
	if outMax == nil {
		outMax = make([]float64, len(real))
	}
	C.TA_MINMAX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMin[0])), (*C.double)(unsafe.Pointer(&outMax[0])))
	return outMin[:outNBElement], outMax[:outNBElement], int(outBegIdx)
}

/*MinMaxIndex - Indexes of lowest and highest values over a specified period

Input = double

Output = int, int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func MinMaxIndex(real []float64, timePeriod int, outMinIdx []int, outMaxIdx []int) ([]int, []int, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outMinIdx == nil {
		outMinIdx = make([]int, len(real))
	}
	if outMaxIdx == nil {
		outMaxIdx = make([]int, len(real))
	}
	C.TA_MINMAXINDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outMinIdx[0])), (*C.int)(unsafe.Pointer(&outMaxIdx[0])))
	return outMinIdx[:outNBElement], outMaxIdx[:outNBElement], int(outBegIdx)
}

/*MinusDi - Minus Directional Indicator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func MinusDi(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_MINUS_DI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*MinusDm - Minus Directional Movement

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func MinusDm(high, low []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_MINUS_DM(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Mom - Momentum

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Mom(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_MOM(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Mult - Vector Arithmetic Mult

Input = double, double

Output = double

*/
func Mult(real0, real1 []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_MULT(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Natr - Normalized Average True Range

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Natr(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_NATR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Obv - On Balance Volume

Input = double, Volume

Output = double

*/
func Obv(real, volume []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_OBV(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*PlusDi - Plus Directional Indicator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func PlusDi(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_PLUS_DI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*PlusDm - Plus Directional Movement

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func PlusDm(high, low []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_PLUS_DM(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Ppo - Percentage Price Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInMAType:

Type of Moving Average

*/
func Ppo(real []float64, fastPeriod, slowPeriod, mAType int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_PPO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Roc - Rate of change : ((price/prevPrice)-1)*100

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Roc(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ROC(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Rocp - Rate of change Percentage: (price-prevPrice)/prevPrice

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocp(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ROCP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Rocr - Rate of change ratio: (price/prevPrice)

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocr(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ROCR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Rocr100 - Rate of change ratio 100 scale: (price/prevPrice)*100

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocr100(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_ROCR100(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Rsi - Relative Strength Index

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Rsi(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_RSI(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sar - Parabolic Sar

Input = High, Low

Output = double

Optional Parameters

-------------------

optInAcceleration:(From 0 to TA_REAL_MAX)

Acceleration Factor used up to the Maximum value

optInMaximum:(From 0 to TA_REAL_MAX)

Acceleration Factor Maximum value

*/
func Sar(high, low []float64, acceleration, maximum float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_SAR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.double(acceleration), C.double(maximum), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*SarExt - Parabolic SAR - Extended

Input = High, Low

Output = double

Optional Parameters

-------------------

optInStartValue:(From TA_REAL_MIN to TA_REAL_MAX)

Start value and direction. 0 for Auto, >0 for Long, <0 for Short

optInOffsetOnReverse:(From 0 to TA_REAL_MAX)

Percent offset added/removed to initial stop on short/long reversal

optInAccelerationInitLong:(From 0 to TA_REAL_MAX)

Acceleration Factor initial value for the Long direction

optInAccelerationLong:(From 0 to TA_REAL_MAX)

Acceleration Factor for the Long direction

optInAccelerationMaxLong:(From 0 to TA_REAL_MAX)

Acceleration Factor maximum value for the Long direction

optInAccelerationInitShort:(From 0 to TA_REAL_MAX)

Acceleration Factor initial value for the Short direction

optInAccelerationShort:(From 0 to TA_REAL_MAX)

Acceleration Factor for the Short direction

optInAccelerationMaxShort:(From 0 to TA_REAL_MAX)

Acceleration Factor maximum value for the Short direction

*/
func SarExt(high, low []float64, startValue, offsetOnReverse, accelerationInitLong, accelerationLong, accelerationMaxLong, accelerationInitShort, accelerationShort, accelerationMaxShort float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_SAREXT(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.double(startValue), C.double(offsetOnReverse), C.double(accelerationInitLong), C.double(accelerationLong), C.double(accelerationMaxLong), C.double(accelerationInitShort), C.double(accelerationShort), C.double(accelerationMaxShort), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sin - Vector Trigonometric Sin

Input = double

Output = double

*/
func Sin(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_SIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sinh - Vector Trigonometric Sinh

Input = double

Output = double

*/
func Sinh(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_SINH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sma - Simple Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Sma(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_SMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sqrt - Vector Square Root

Input = double

Output = double

*/
func Sqrt(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_SQRT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*StdDev - Standard Deviation

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)

Nb of deviations

*/
func StdDev(real []float64, timePeriod int, nbDev float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_STDDEV(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDev), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Stoch - Stochastic

Input = High, Low, Close

Output = double, double

Optional Parameters

-------------------

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInSlowK_Period:(From 1 to 100000)

Smoothing for making the Slow-K line. Usually set to 3

optInSlowK_MAType:

Type of Moving Average for Slow-K

optInSlowD_Period:(From 1 to 100000)

Smoothing for making the Slow-D line

optInSlowD_MAType:

Type of Moving Average for Slow-D

*/
func Stoch(high, low, close []float64, fastKPeriod, slowKPeriod, slowKMAType, slowDPeriod, slowDMAType int, outSlowK []float64, outSlowD []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outSlowK == nil {
		outSlowK = make([]float64, len(high))
	}
	if outSlowD == nil {
		outSlowD = make([]float64, len(high))
	}
	C.TA_STOCH(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(fastKPeriod), C.int(slowKPeriod), C.TA_MAType(slowKMAType), C.int(slowDPeriod), C.TA_MAType(slowDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outSlowK[0])), (*C.double)(unsafe.Pointer(&outSlowD[0])))
	return outSlowK[:outNBElement], outSlowD[:outNBElement], int(outBegIdx)
}

/*Stochf - Stochastic Fast

Input = High, Low, Close

Output = double, double

Optional Parameters

-------------------

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInFastD_Period:(From 1 to 100000)

Smoothing for making the Fast-D line. Usually set to 3

optInFastD_MAType:

Type of Moving Average for Fast-D

*/
func Stochf(high, low, close []float64, fastKPeriod, fastDPeriod, fastDMAType int, outFastK []float64, outFastD []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outFastK == nil {
		outFastK = make([]float64, len(high))
	}
	if outFastD == nil {
		outFastD = make([]float64, len(high))
	}
	C.TA_STOCHF(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(fastKPeriod), C.int(fastDPeriod), C.TA_MAType(fastDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outFastK[0])), (*C.double)(unsafe.Pointer(&outFastD[0])))
	return outFastK[:outNBElement], outFastD[:outNBElement], int(outBegIdx)
}

/*StochRsi - Stochastic Relative Strength Index

Input = double

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInFastD_Period:(From 1 to 100000)

Smoothing for making the Fast-D line. Usually set to 3

optInFastD_MAType:

Type of Moving Average for Fast-D

*/
func StochRsi(real []float64, timePeriod, fastKPeriod, fastDPeriod, fastDMAType int, outFastK []float64, outFastD []float64) ([]float64, []float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outFastK == nil {
		outFastK = make([]float64, len(real))
	}
	if outFastD == nil {
		outFastD = make([]float64, len(real))
	}
	C.TA_STOCHRSI(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.int(fastKPeriod), C.int(fastDPeriod), C.TA_MAType(fastDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outFastK[0])), (*C.double)(unsafe.Pointer(&outFastD[0])))
	return outFastK[:outNBElement], outFastD[:outNBElement], int(outBegIdx)
}

/*Sub - Vector Arithmetic Substraction

Input = double, double

Output = double

*/
func Sub(real0, real1 []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real0))
	}
	C.TA_SUB(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Sum - Summation

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Sum(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_SUM(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*T3 - Triple Exponential Moving Average (T3)

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInVFactor:(From 0 to 1)

Volume Factor

*/
func T3(real []float64, timePeriod int, vFactor float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_T3(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(vFactor), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Tan - Vector Trigonometric Tan

Input = double

Output = double

*/
func Tan(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TAN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Tanh - Vector Trigonometric Tanh

Input = double

Output = double

*/
func Tanh(real []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TANH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Tema - Triple Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Tema(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TEMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Trange - True Range

Input = High, Low, Close

Output = double

*/
func Trange(high, low, close []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_TRANGE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*TriMa - Triangular Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func TriMa(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TRIMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Trix - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Trix(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TRIX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Tsf - Time Series Forecast

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Tsf(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_TSF(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*TypPrice - Typical Price

Input = High, Low, Close

Output = double

*/
func TypPrice(high, low, close []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_TYPPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*UltOsc - Ultimate Oscillator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod1:(From 1 to 100000)

Number of bars for 1st period.

optInTimePeriod2:(From 1 to 100000)

Number of bars fro 2nd period

optInTimePeriod3:(From 1 to 100000)

Number of bars for 3rd period

*/
func UltOsc(high, low, close []float64, timePeriod1, timePeriod2, timePeriod3 int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_ULTOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod1), C.int(timePeriod2), C.int(timePeriod3), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Var - Variance

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)

Nb of deviations

*/
func Var(real []float64, timePeriod int, nbDev float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_VAR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDev), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*WclPrice - Weighted Close Price

Input = High, Low, Close

Output = double

*/
func WclPrice(high, low, close []float64, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_WCLPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Willr - Williams' %R

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Willr(high, low, close []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(high))
	}
	C.TA_WILLR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}

/*Wma - Weighted Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Wma(real []float64, timePeriod int, outReal []float64) ([]float64, int) {
	var outBegIdx C.int
	var outNBElement C.int
	if outReal == nil {
		outReal = make([]float64, len(real))
	}
	C.TA_WMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal[:outNBElement], int(outBegIdx)
}
