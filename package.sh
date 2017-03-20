#!/bin/bash
go build

cd ui
npm run build:aot:prod
cd ..
mkdir dist
cd dist
mkdir malhora_info
cp ../malhora_info malhora_info/
cp -rf ../ui/dist malhora_info/public
tar cvjf malhora_info.dist.tar.bz2 malhora_info
cd ..
