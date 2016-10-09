package randomizer

import "testing"

func Test_should_create_random_ints(t *testing.T) {
	randomInts := Ints(2, 10)

	if len(randomInts) != 2 {
		t.Errorf("length should be %d but is %d", 2, len(randomInts))
	}

	for i := range randomInts {
		if randomInts[i] >= 10 {
			t.Errorf("should be less than %d but is %d", 10, randomInts[i])
		}
	}
}

func Test_should_create_random_strings(t *testing.T) {
	randomStrings := Strings(2, 5)

	if len(randomStrings) != 2 {
		t.Errorf("length should be %d but is %d", 2, len(randomStrings))
	}

	for i := range randomStrings {
		if len(randomStrings[i]) > 5 {
			t.Errorf("length should be %d but is %d", 5, len(randomStrings[i]))
		}
	}
}
