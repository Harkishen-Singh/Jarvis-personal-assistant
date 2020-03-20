# run the service as a background process
cd service/ && go run main.go &
# cd ..

# run js linting checks
eslint view/web/app-jarvis.js
eslint view/desktop/app/*.js view/desktop/app/scripts/*.js

# run html lint checks
# htmlhint view/web/*.html view/web/components/*.html
htmlhint view/desktop/app/templates/components/*.html view/desktop/app/templates/*.html

# go unit tests
# cd service/ && go test ./...
# cd ..

# application build
cd service/ && go build *.go
cd ..
mv service/main jarvis
# cd view/desktop/ && npm run build
# cd ../..
