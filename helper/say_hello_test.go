package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {
	result := SayHelloBro("Fauzan")

	if result != "hello Fauzan" {
		// t.Fail()
		t.Error("gagal")
		// t.FailNow()
		// t.Fatal()

	}

}

func TestSayHelloAssert(t *testing.T) {
	result := SayHelloBro("fauzana")
	assert.Equal(t, "hello fauzan", result, "tidak sesuai")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("before")
	m.Run()
	//after
	fmt.Println("after")
}

func TestSubTest(t *testing.T) {
	t.Run("Fauzan", func(t *testing.T) {
		result := SayHelloBro("fauzan")
		assert.Equal(t, "hello fauzan", result, "tidak sesuai")
	})

	t.Run("fakhri", func(t *testing.T) {
		result := SayHelloBro("Fakhri")
		assert.Equal(t, "hello Fakhri", result, "tidak sesuai")
	})
}

func TestTableSayHelloBro(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "fauzan",
			request:  "fauzan",
			expected: "hello fauzan",
		},
		{
			name:     "fakhri",
			request:  "fakhri",
			expected: "hello fakhri",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHelloBro(test.request)
			require.Equal(t, test.expected, result)

		})
	}
}
