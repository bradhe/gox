clean:
	# Not implemented

build: clean
	go generate ./...

test: clean
	go test -v ./...

release: 
	# Not implemented
