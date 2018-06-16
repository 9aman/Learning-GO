# BET
 - Benchmark
 - Example
 - Test

 ```
 BenchmarkCat(b *testing.B)
 ExampleCat()
 TestCat(t *testing.T)
 ```

 # Commands

 ```

 godoc -http=:8080

 go test
 go test -bench .
 go test -cover
 go test -coverprofile cover.out
 go tool cover -html=cover.out
 ```