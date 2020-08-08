package random

import (
	"math/rand"
	"testing"
)

func BenchmarkJumpConsistentHash(b *testing.B) {
	numBuckets := 10
	key := rand.Uint64()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e6; j++ {
			JumpConsistentHash(key, numBuckets)
		}
	}
}

func TestJumpConsistentHash(t *testing.T) {
	numKeys := 3
	for numBuckets := 1; numBuckets < 15; numBuckets++ {
		rand.Seed(86)
		bucketMapping := make([]int32, numKeys)
		for i := 0; i < numKeys; i++ {
			bucketMapping[i] = JumpConsistentHash(rand.Uint64(), numBuckets)
		}
		t.Logf("Buckets: %d, %d\n", numBuckets, bucketMapping)
	}
}

func TestJumpConsistentHashDist(t *testing.T) {
	numKeys := 1000000
	numBuckets := 256

	ba := make([][]uint64, numBuckets)

	rand.Seed(86)
	for i := 0; i < numKeys; i++ {
		key := rand.Uint64()
		bucket := JumpConsistentHash(key, numBuckets)
		ba[bucket] = append(ba[bucket], key)
	}

	for b, keys := range ba {
		t.Logf("Bucket: %d, Keys: %d\n", b, len(keys))
	}
}
