#!/bin/bash

# Run your test files in the desired order
go test -v ./pkg/test_apis/add_test.go
go test -v ./pkg/test_apis/get_test.go
go test -v ./pkg/test_apis/update_test.go
go test -v ./pkg/test_apis/delete_test.go

# Add more test files in the desired order
