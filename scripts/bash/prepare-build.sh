#!/usr/bin/env bash

echo "Building GeoSite Encoder..."
go build -v -trimpath -o ./geosite ./main.go
echo "Building GeoSite Data Dumper..."
go build -v -trimpath -o ./datadump ./geositedumper/main.go
echo "Building GeoSite Extract-ABP..."
go build -v -trimpath -o ./extract-abp ./tool/abpextract/main.go
echo "Building GeoSite Extract-DNSMasq..."
go build -v -trimpath -o ./extract-dnsmasq ./tool/dnsmasqextract/main.go
echo "Building GeoSite Sanitizer..."
go build -v -trimpath -o ./sanitize ./tool/sanitizer/main.go

echo "Preparing directories..."
mv ./predata ./data
mkdir ./temp/
echo "Preparation completed."
