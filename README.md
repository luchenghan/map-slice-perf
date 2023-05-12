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

### 結論
- slice 寫入資料較快；map則較慢需要做hash index。
- slice 適合存放較小資料集，當資料足夠大時，取資料會隨著O(N)變慢。
- map 適合存放較大資料集，不會因為資料變大而變慢，都會維持在O(1)情況。

### Reference
[link](https://betterprogramming.pub/performance-impact-of-maps-compared-to-slices-in-go-1-18-15352fbd6010)