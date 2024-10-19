.PHONY : build
build :
	go build -o ./bin/guru ./cmd/guru/*.go

.PHONY : run
run : build
	./bin/guru

.PHONY : test
test :
	go test ./...

.PHONY : install
install : 
	cp ./bin/guru /usr/local/bin

