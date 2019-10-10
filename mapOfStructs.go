package random

import "math"

type ValuesSimple struct {
	Items []float64
}

type Values struct {
	Items []float64
	Sum   float64
	Max   float64
	Min   float64
}

func initMapOfStructsSimple() map[string]ValuesSimple {
	numKeys := 1000
	numItems := 4800
	out := make(map[string]ValuesSimple)

	for i := 0; i < numKeys; i++ {
		out[string(i)] = ValuesSimple{Items: make([]float64, numItems)}
	}

	return out
}

func initMapOfStructsSimplePtr() map[string]*ValuesSimple {
	numKeys := 1000
	numItems := 4800
	out := make(map[string]*ValuesSimple)

	for i := 0; i < numKeys; i++ {
		out[string(i)] = &ValuesSimple{Items: make([]float64, numItems)}
	}

	return out
}

func fillMapOfStructsSimple(m map[string]ValuesSimple) {
	for _, v := range m {
		for i := 0; i < cap(v.Items); i++ {
			v.Items[i] = float64(i)
		}
	}
}

func fillMapOfStructsSimpleReferenceKey(m map[string]ValuesSimple) {
	for k := range m {
		for i := 0; i < cap(m[k].Items); i++ {
			m[k].Items[i] = float64(i)
		}
	}
}

func fillMapOfStructsSimplePtr(m map[string]*ValuesSimple) {
	for _, v := range m {
		for i := 0; i < cap(v.Items); i++ {
			v.Items[i] = float64(i)
		}
	}
}

func fillMapOfStructsSimplePtrReferenceKey(m map[string]*ValuesSimple) {
	for k := range m {
		for i := 0; i < cap(m[k].Items); i++ {
			m[k].Items[i] = float64(i)
		}
	}
}

func initMapOfStructs() map[string]Values {
	numKeys := 1000
	numItems := 4800
	out := make(map[string]Values)

	for i := 0; i < numKeys; i++ {
		out[string(i)] = Values{Items: make([]float64, numItems)}
	}

	return out
}

func initMapOfStructsPtr() map[string]*Values {
	numKeys := 1000
	numItems := 4800
	out := make(map[string]*Values)

	for i := 0; i < numKeys; i++ {
		out[string(i)] = &Values{Items: make([]float64, numItems)}
	}

	return out
}

func fillMapOfStructs(m map[string]Values) {
	for k, v := range m {
		for i := 0; i < cap(v.Items); i++ {
			v.Items[i] = float64(i)
			v.Sum += float64(i)
			v.Max = math.Max(v.Max, float64(i))
			v.Min = math.Max(v.Min, float64(i))
			m[k] = v
		}
	}
}

func fillMapOfStructsReferenceKey(m map[string]Values) {
	for k := range m {
		for i := 0; i < cap(m[k].Items); i++ {
			m[k].Items[i] = float64(i)
		}
	}
}

func fillMapOfStructsPtr(m map[string]*Values) {
	for _, v := range m {
		for i := 0; i < cap(v.Items); i++ {
			v.Items[i] = float64(i)
			v.Sum += float64(i)
			v.Max = math.Max(v.Max, float64(i))
			v.Min = math.Min(v.Min, float64(i))
		}
	}
}

func fillMapOfStructsPtrReferenceKey(m map[string]*Values) {
	for k := range m {
		for i := 0; i < cap(m[k].Items); i++ {
			m[k].Items[i] = float64(i)
			m[k].Sum += float64(i)
			m[k].Max = math.Max(m[k].Max, float64(i))
			m[k].Min = math.Max(m[k].Min, float64(i))
		}
	}
}
