install:
	cd service/ && go get -v -u -d ./...

install-all || update:
	cd service/ && go get -v -u -d ./...
	cd view/desktop/ && npm install

run:
	cd service/ && go run maintut.go

views:
	cd view/desktop/ && npm start
