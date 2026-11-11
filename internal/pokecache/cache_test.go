package pokecache

import (
	"testing"
)

type testCase struct {
	name          string
	input         string
	inputKey      string
	expected      []byte
	expectedFound bool
}

func TestCache(t *testing.T) {
	var nilExpected []byte
	addCases := []testCase{
		{
			name:          "Add Correctly",
			input:         "https://example.com",
			inputKey:      "Add Correctly",
			expected:      []byte("https://example.com"),
			expectedFound: false,
		},
		{
			name:          "Add Second",
			input:         "https://google.com",
			inputKey:      "Add Second",
			expected:      []byte("https://google.com"),
			expectedFound: false,
		},
		{
			name:          "Update Correctly",
			input:         "https://google.com",
			inputKey:      "Add Correctly",
			expected:      []byte("https://google.com"),
			expectedFound: true,
		},
	}
	getCases := []testCase{
		{
			name:          "Get Existing",
			input:         "https://example.com",
			inputKey:      "Add Correctly",
			expected:      []byte("https://google.com"),
			expectedFound: true,
		},
		{
			name:          "Not existing",
			input:         "https://notexample.com",
			inputKey:      "failed get",
			expected:      nilExpected,
			expectedFound: false,
		},
	}
	cache := NewCache()
	for _, current := range addCases {
		bytes := []byte(current.input)
		cache.Add(current.inputKey, bytes)
		val, ok := cache.Get(current.inputKey)
		if !ok {
			t.Errorf("%s: Not found in cache", current.name)
		}
		if val == nil {
			t.Errorf("%s: Get returned nil", current.name)
		}
		if string(val) != string(bytes) {
			t.Errorf("%s: Bytes input did not match bytes output", current.name)
		}
	}
	for _, current := range getCases {
		val, ok := cache.Get(current.inputKey)
		if ok != current.expectedFound {
			t.Errorf("Key not found: %s", current.inputKey)
		}
		if current.expectedFound == false && val != nil {
			t.Errorf("%s: Expected nil, got %s", current.name, val)
		}
		if current.expectedFound == true && val == nil {
			t.Errorf("%s: Get returned nil", current.name)
			continue
		}
		if string(val) != string(current.expected) {
			t.Errorf("Found key does not match expected value")
			continue
		}

	}

}
