#!/bin/bash

PWD=$(pwd)

cd ../../hlf-network || exit 1

./minifab cleanup

cd ..

rm -rf hlf-network

cd "$PWD" || exit 1