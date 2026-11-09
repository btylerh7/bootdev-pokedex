package pokecache

import (
	"fmt"
	"testing"
)

type testCase struct {
	name          string
	input         string
	inputKey      string
	expected      []byte
	expectedFound bool
}

func TestCacheAdd(t *testing.T) {
	fmt.Println("Testing adding to cache")
	failCount := 0
	passCount := 0
	var nilExpected []byte
	addCases := []testCase{
		{
			name:          "Add Correctly",
			input:         "https://example.com",
			inputKey:      "successful test",
			expected:      []byte("https://example.com"),
			expectedFound: false,
		},
	}
	getCases := []testCase{
		{
			name:          "Get Existing",
			input:         "https://example.com",
			inputKey:      "successful test",
			expected:      []byte("https://example.com"),
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
			failCount++
			t.Errorf("%s: Not found in cache", current.name)
		}
		if string(val) != string(bytes) {
			failCount++
			t.Errorf("%s: Bytes input did not match bytes output", current.name)
		}
		passCount++
	}
	for _, current := range getCases {
		val, ok := cache.Get(current.inputKey)
		if ok != current.expectedFound {
			failCount++
			t.Errorf("Key not found: %s", current.inputKey)
			continue
		}
		if string(val) != string(current.expected) {
			failCount++
			t.Errorf("Found key does not match expected value")
			continue
		}
		passCount++

	}
	t.Logf("Passed: %d\nFailed: %d", passCount, failCount)

}
