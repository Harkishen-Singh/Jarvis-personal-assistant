install:
	cd service/ && go get -v -u -d ./...

install-all || update:
	cd service/ && go get -v -u -d ./...
	cd view/desktop/ && npm install

run:
	go run service/main.go

views:
	cd view/desktop/ && npm start

build:
	mkdir bin
	cd service/ && go build main.go
	mv service/main bin/service

clean:
	rm -R bin

fix:
	go fmt ./...
