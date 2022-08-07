package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Apnan cek",
			request: "Apnan cek",
			expected: "Hello Apnan cek",
		},
		{
			name: "Juanda",
			request: "Juanda",
			expected: "Hello Juanda",
		},
	}
	
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "result must be " + test.expected)
		})
	}
}