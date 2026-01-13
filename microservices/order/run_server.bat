@echo off
set "DATA_SOURCE_URL=appuser:apppass@tcp(localhost:3307)/orders_db?charset=utf8mb4&parseTime=True&loc=Local"
set APPLICATION_PORT=3000
set ENV=development
go run cmd/main.go
