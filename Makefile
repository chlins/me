
build:
	go build -o ./bin/me ./me.go

clean:
	rm -rf ./bin/*

.PHONY: build clean

