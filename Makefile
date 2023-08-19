generate:
	go mod init github.com/mochammadshenna/clean-architecture
	gocto generate
	mockery --all
	go mod tidy
mock:
	mockery --all
test:
	go test ./...
