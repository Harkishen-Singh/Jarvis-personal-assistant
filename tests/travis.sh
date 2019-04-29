which eslint
eslint ../app/*.js ../app/scripts/*.js

echo "building Jarvis-Desktop package"

cd ..
npm run build
echo "successfully build Jarvis-Desktop"

echo "cleaning ... "
rm -R Jarvis-Desktop-linux-x64
