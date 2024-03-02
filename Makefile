build:
	go build ./gateway/cmd/main.go
	go build ./comments/cmd/main.go
	go build ./censors/cmd/main.go
	go build ./aggregator/cmd/server/main.go

run: build
