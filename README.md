# Golang testing

### Tool
- benchstat
  - go install golang.org/x/perf/cmd/benchstat@latest
  - go install 前需要 go mod init module name

### Benchmarks
- 產生壓力測試資料
```
go test -bench="BenchmarkCallMapValues" -run=^# -count=5 | tee map.txt

go test -bench="BenchmarkCallSliceValues" -run=^# -count=5 | tee slice.txt
```

- 統計壓力測試資料
```
benchstat *.txt
```

### Reference
[link](https://betterprogramming.pub/performance-impact-of-maps-compared-to-slices-in-go-1-18-15352fbd6010)