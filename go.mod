module github.com/ly123525/study-notes

go 1.17

require logger v0.0.1

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
)

replace logger => ./go/logger
