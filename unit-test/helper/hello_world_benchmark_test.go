package helper

import "testing"

//table Benchmark
func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Apnan",
			request: "Apnan",
		},
		{
			name: "Juanda",
			request: "Juanda",
		},
	}

	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++{
				HelloWorld(benchmark.request)
			}
		})
	}
}

//sub Benchmark
func BenchmarkSub(b *testing.B){
	b.Run("Apnan", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Apnan")
		}
	})

	b.Run("Juanda", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Juanda")
		}
	})
}

/*
	to execute sub benchmark
	go test -v -run=TestTidakAda -bench=BenchmarkSub/Apnan
*/

func BenchmarkHelloWorldApnan(b *testing.B){
	for i := 0; i < b.N; i++{
		HelloWorld("Apnan")
	} 
}

func BenchmarkHelloWorldJuanda(b *testing.B){
	for i := 0; i < b.N; i++{
		HelloWorld("Juanda")
	} 
}

/*
	to execute bencmark
	go test -v -bench=
	go test -v -run=TestTidakAda -bench=
*/