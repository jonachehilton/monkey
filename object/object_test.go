package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{"Hello"}
	hello2 := &String{"Hello"}
	diff1 := &String{"Diff"}
	diff2 := &String{"Diff"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
