package isPalindrome

import "testing"

type exampleSentences struct {
	sentence       string
	expectedResult bool
}

func TestPalindrome(t *testing.T) {
	sentences := []exampleSentences{
		{
			sentence:       "A Santa at Nasa.",
			expectedResult: true,
		},
		{
			sentence:       "Are we not drawn onward to new era?",
			expectedResult: true,
		},
		{
			sentence:       "Lager, sir, is regal.",
			expectedResult: true,
		},
		{
			sentence:       "Olson is in Oslo.",
			expectedResult: true,
		},
		{
			sentence:       "This is obviously not a palindrome, is it?",
			expectedResult: false,
		},
	}

	for _, sentence := range sentences {
		t.Logf("Checking sentence: %s", sentence.sentence)
		result := Check(sentence.sentence)
		if result != sentence.expectedResult {
			t.Errorf("Expected value of %t, but it was %t", sentence.expectedResult, result)
		}
	}
}
