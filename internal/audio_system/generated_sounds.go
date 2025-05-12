package audio_system

import "math"

/**/
func Soundwave_CreateSound(SampleRate, BaseFreq, q, refFreq int, decayAmp, decayX []float32) []byte {
	// const refFreq = 110
	dd := 5    //5
	ee := 12.0 //12
	length := dd * SampleRate * BaseFreq / refFreq
	refData := make([]float32, length)
	for i := 0; i < length; i++ {
		refData[i] = Soundwave_NoiseAt(SampleRate, BaseFreq, i, float32(refFreq), 5.0, decayAmp, decayX)
	}

	freq := float64(BaseFreq) * math.Exp2(float64(q-1)/ee) //12.0

	// Calculate the wave data for the freq.
	length02 := dd * SampleRate * BaseFreq / int(freq)
	l := make([]float32, length02)
	r := make([]float32, length02)
	for i := 0; i < length02; i++ {
		idx := int(float64(i) * freq / float64(refFreq))
		if len(refData) <= idx {
			break
		}
		l[i] = refData[idx]
	}
	copy(r, l)
	n := Soundwave_ToBytes(l, r)
	return n
}

/*This is a copy of  the ebiten examples "PianoAt" function;*/
func Soundwave_NoiseAt(sRate, BaseFreq, i int, freq, divBy float32, amp, decayX []float32) float32 {
	var v float32
	for j := 0; j < len(amp); j++ {
		// Decay
		a := amp[j] * float32(math.Exp(float64(-5*float32(i)*freq/float32(BaseFreq)/(decayX[j]*float32(sRate)))))
		v += a * float32(math.Sin(2.0*math.Pi*float64(i)*float64(freq)*float64(j+1)/float64(sRate)))
	}
	return v / divBy
}

func Soundwave_ToBytes(l, r []float32) []byte {
	if len(l) != len(r) {
		panic("len(l) must equal to len(r)")
	}
	b := make([]byte, len(l)*8)
	for i := range l {
		lv := math.Float32bits(l[i])
		rv := math.Float32bits(r[i])
		b[8*i] = byte(lv)
		b[8*i+1] = byte(lv >> 8)
		b[8*i+2] = byte(lv >> 16)
		b[8*i+3] = byte(lv >> 24)
		b[8*i+4] = byte(rv)
		b[8*i+5] = byte(rv >> 8)
		b[8*i+6] = byte(rv >> 16)
		b[8*i+7] = byte(rv >> 24)
	}
	return b
}
