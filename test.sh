#! /bin/bash
rm database/test.db
go clean -testcache
go mod tidy
go test -v ./...