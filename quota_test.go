package random

import (
	"testing"
	"time"
)

func TestQuotaInvalidUser(t *testing.T) {
	m := newManager()
	m.run()

	ok, err := m.useToken("user1")
	if err == nil {
		t.Fatalf("Should have returned an error for an invalid user")
	}
	if ok {
		t.Fatalf("Should not have allowed a token use for user1")
	}
}

func TestQuotaValidUser(t *testing.T) {
	m := newManager()
	m.run()

	user := "user1"
	m.addRule(user, newRule(1, 5*time.Second))

	ok, err := m.useToken(user)
	if err != nil {
		t.Fatalf("Did not expect an error on valid user, %v", err)
	}
	if ok {
		t.Fatalf("Should not have allowed token use since the rule was just created %+v", *m.rules[user])
	}
}

func TestQuotaCountMax(t *testing.T) {
	m := newManager()
	m.run()

	m.addRule("user1", newRule(1, 5*time.Second))
	m.addRule("user2", newRule(1, 1*time.Second))
	m.addRule("user3", newRule(4, 2*time.Second))

	// wait till all tokens are given out
	time.Sleep(10 * time.Second)

	m.mu.Lock()
	for k, r := range m.rules {
		if r.count != r.maxQueries {
			t.Fatalf("Expected %d tokens available but got %d, for %s", r.maxQueries, r.count, k)
		}
	}
	m.mu.Unlock()
}

func BenchmarkQuotaUpdateMillionKeys(b *testing.B) {
	m := newManager()

	for i := 0; i < 10000; i++ {
		m.addRule(string(i), newRule(1, 30*time.Second))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.addTokens()
	}
}

func BenchmarkQuotaUseMillionKeys(b *testing.B) {
	m := newManager()
	m.run()

	numKeys := 10000
	for i := 0; i < numKeys; i++ {
		m.addRule(string(i), newRule(1, 5*time.Second))
	}
	time.Sleep(6 * time.Second)

	b.ResetTimer()
	var numOk int
	for n := 0; n < b.N; n++ {
		if ok, err := m.useToken(string(n % numKeys)); err == nil && ok {
			numOk++
		}
	}
	b.Logf("Got %d ok out of %d", numOk, b.N)
}
