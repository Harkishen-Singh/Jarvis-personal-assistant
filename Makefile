install:
	cd service/ && go get -v -u -d ./...

install-all || update:
	cd service/ && go get -v -u -d ./...
	cd view/desktop/ && npm install

run:
	node src/service/main.js

app:
	cd src/view/desktop/ && npm start

clean:
	rm -R bin

fix:
	go fmt ./...
