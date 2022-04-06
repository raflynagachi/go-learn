package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("=====BEFORE TEST=====")
	m.Run()
	fmt.Println("=====AFTER TEST=====")
}

func BenchmarkTableTest(b *testing.B) {
	test_cases := []struct {
		name    string
		request string
	}{
		{
			name:    "Nagachi",
			request: "Nagachi",
		},
		{
			name:    "Rafly",
			request: "Rafly",
		},
		{
			name:    "Rafly Rigan Nagachi",
			request: "Rafly Rigan Nagachi",
		},
	}
	for _, val := range test_cases {
		b.Run(val.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(val.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Nagachi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Nagachi")
		}
	})
	b.Run("Rafly", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Rafly")
		}
	})
}

func BenchmarkHelloNagachi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Nagachi")
	}
}
func BenchmarkHelloRafly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Rafly Rigan Nagachi")
	}
}

func TestTableHello(t *testing.T) {
	var testCases = []struct {
		name     string
		request  string
		expected string
		message  string
	}{
		{
			name:     "Hello(Rafly)",
			request:  "Rafly",
			expected: "Hello, Rafly",
			message:  "Must be equal to 'Hello, Rafly'",
		},
		{
			name:     "Hello(Rigan)",
			request:  "Rigan",
			expected: "Hello, Rigan",
			message:  "Must be equal to 'Hello, Rafly'",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			assert.Equal(t, test.expected, result, test.message)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Rigan", func(t *testing.T) {
		result := Hello("Rigan")
		assert.Equal(t, "Hello, Rigan", result, "FAIL: Must be equal to 'Hello, Rigan'")
	})
	t.Run("Nagachi", func(t *testing.T) {
		result := Hello("Nagachi")
		assert.Equal(t, "Hello, Nagachi", result, "FAIL: Must be equal to 'Hello, Nagachi'")
	})
}
func TestSkipFunction(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("SKIP: Cannot run function on linux")
	}
}
func TestHelloNagachiAssertion(t *testing.T) {
	result := Hello("Nagachi")
	assert.Equal(t, "Hello, Nagachi", result, "must be equal to 'Hello, Nagachi'")
	fmt.Println("Test done: TestHelloNagachiAssertion")
}

func TestHelloRaflyAssertion(t *testing.T) {
	result := Hello("Rafly")
	require.Equal(t, "Hello, Rafly", result, "must be equal to 'Hello, Rafly'")
	fmt.Println("Test done: TestHelloRaflyAssertion")
}

func TestHelloNagachi(t *testing.T) {
	result := Hello("Nagachi")
	if result != "Hello, Nagachi" {
		// error
		t.Error("FAIL: must be \"Hello, Nagachi\"")
	}
	fmt.Println("Test done: TestHelloNagachi")
}

func TestHelloRafly(t *testing.T) {
	result := Hello("Rafly")
	if result != "Hello, Rafly" {
		// error
		t.Fatal("FAIL: must be \"Hello, Rafly\"")
	}
	// doesnt printed cause FailNow()
	fmt.Println("Test done: TestHelloRafly")
}
