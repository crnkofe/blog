# Fuzzin around

This is a short demo of fuzzin(using Go 1.16.5).

At the moment fuzzing is still in Beta. To get it execute:
```bash
go get golang.org/dl/gotip
gotip download dev.fuzz
```

Note: `gotip` execution will take a while.

To build and run the code:
```bash
go build bad_code.go
./bad_code
```

To run unit tests:
```bash
go test .
```

To run fuzzer:
```bash
gotip test .
```
