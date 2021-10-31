.PHONY: exec
exec:
	go run main.go

.PHONY: vendor
vendor:
	go mod tidy & go mod vendor
