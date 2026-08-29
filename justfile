# ASSEMBLY: go build -gcflags="-S" ./pkg/collections > cmd.S 2>&1
set windows-shell := ["powershell.exe", "-c"]

export GOEXPERIMENT := 'simd'

default:
    just --list

documentation:
    go doc -all -u -http

[group('checks')]
build:
    go build ./...
    go build -tags simd_none ./...
    go build -tags simd_detect ./...
    go build -tags simd_avx256 ./...
    go build -tags simd_avx512 ./...

[group('checks')]
test:
    go test -v ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

show-coverage-report:
    go tool cover -html=reports/coverage.out

coverage-report: test show-coverage-report

[group('checks')]
generate:
    go generate ./...

[group('checks')]
lint:
    go tool golangci-lint run -v --fix

[group('checks')]
format:
    go fmt ./...

[group('checks')]
fix:
    go fix ./...
