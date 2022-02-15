package grpc

func ConvertArrayFloat32to64(in []float32) []float64 {
	var out = make([]float64, 0)

	for _, v := range in {
		out = append(out, float64(v))
	}

	return out
}

func ConvertArrayFloat64to32(in []float64) []float32 {
	var out = make([]float32, 0)

	for _, v := range in {
		out = append(out, float32(v))
	}

	return out
}
