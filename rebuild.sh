#!/bin/bash

rm -rf ./fdk-aac
mkdir ./fdk-aac
curl -L https://github.com/mstorsjo/fdk-aac/archive/refs/tags/v2.0.3.tar.gz | tar -xz --strip-components=1 -C ./fdk-aac

cd ./fdk-aac

rm -rf ./documentation/

./autogen.sh
./configure
make -j$(nproc) CFLAGS="-g -O2 -Werror" CXXFLAGS="-g -O2 -Werror"
mv ./.libs ./libs