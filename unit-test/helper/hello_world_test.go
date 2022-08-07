package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T){
	t.Run("Apnan", func(t *testing.T){
		result := HelloWorld("Apnan")
		assert.Equal(t, "Hello Apnan", result, "result should be 'Hello Apnan'")
	})
	t.Run("Juanda", func(t *testing.T){
		result := HelloWorld("Juanda")
		assert.Equal(t, "Hi Juanda", result, "result should be 'Hello Juanda'")
	})
	//to execute only subtest
	//go test -v -run /Juanda
}

func TestMain(t *testing.M){
	//BEFORE
	fmt.Println("BEFORE UNIT TEST")
	
	t.Run()

	//AFTER
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Can not run on Windows OS")
	}
	fmt.Println("Cek SKIP")
	result := HelloWorld("Apnan")
	require.Equal(t, "Hello Apnan", result, "result should be 'Hello Apnan'")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Apnan")
	assert.Equal(t, "Hello Apnan", result, "Result must be 'Hello Apnan'")
	fmt.Println("Test HelloWorld with assert done") //assert execute fail
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Apnan")
	require.Equal(t, "Hello Apnan", result, "Result must be 'Hello Apnan'")
	fmt.Println("Test HelloWorld with require done") //require execute failNow
}

func TestHelloWorldApnan(t *testing.T) {
	result := HelloWorld("Apnan")
	if result != "Hello Apnan" {
		//unit test failed
		t.Error("Result must be 'Hello Apnan'")
	}

	fmt.Println("TestHelloWorldApnan done") //this is stay to execute
}

func TestHelloWorldJuanda(t *testing.T) {
	result := HelloWorld("Juanda")
	if result != "Hello Juanda" {
		//unit test failed
		t.Fatal("Result must be 'Hello Juanda'")
	}
	fmt.Println("TestHelloWorldJuanda done") //this is not to execute
}

//to execution test
/*
	go test
	go test -v
	go test -v -run TestHelloWorld
	go test -v -run=TestHelloWorld
	cd ..
	go test -v ./...
*/
