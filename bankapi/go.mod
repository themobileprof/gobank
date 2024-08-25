module bankapi

go 1.22.5

require (
	github.com/themobileprof/bank v0.0.1
	github.com/themobileprof/db v0.0.1
	golang.org/x/exp v0.0.0-20240808152545-0cdaa3abc0fa
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)

replace github.com/themobileprof/bank => ../bankcore

replace github.com/themobileprof/db => ../db
