package random

func JumpConsistentHash(key uint64, buckets int) int32 {
	var b, j int64
	b = -1
	j = 0

	if buckets < 1 {
		buckets = 1
	}

	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int32(b)
}
