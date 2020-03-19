install:
	cd service/ && go get -v -u -d ./...

install-all || update:
	cd service/ && go get -v -u -d ./...
	cd view/desktop/ && npm install

run:
	go run service/maintut.go

views:
	cd view/desktop/ && npm start

build:
	cd service/ && go build maintut.go
	mv service/maintut bin/service
