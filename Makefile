test:
	go test -v ./...

testrace:
	go test -race github.com/lazharichir/gocached -v -count=1 ./...