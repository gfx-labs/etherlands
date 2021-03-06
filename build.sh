#!/bin/sh

flatc --go ./flatbuffer/etherlands.fbs
mkdir -p db
mv Etherlands/* proto
rm -r Etherlands
abigen --abi ./district.abi --pkg main --type DistrictContract --out district_contract.go

go clean
go build .
