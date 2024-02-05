#!/bin/bash

cd ../web

npm install
npm run build

rm -rf ../dist/web/static/*
cp -r dist/*  ../dist/web/static

cd ../console

go build main.go
yes | cp main ../dist

cd ../build