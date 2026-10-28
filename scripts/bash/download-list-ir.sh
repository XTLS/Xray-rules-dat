#!/usr/bin/env bash

echo "Downloading Iran related lists..."
curl -L -o ./temp/pre-ir-redundant https://raw.githubusercontent.com/Chocolate4U/Iran-v2ray-rules/refs/heads/redundant/redundant-domains.txt
curl -L -o ./temp/pre-iran-hosted-domains https://raw.githubusercontent.com/bootmortis/iran-hosted-domains/refs/heads/release/domains.txt
cat ./temp/pre-iran-hosted-domains | grep -Ev ".+\.ir$" | sed '1 a\ir\nxn--mgba3a4f16a' | LC_ALL=C sort -u > ./temp/iran-hosted-domains
comm -23 ./temp/iran-hosted-domains ./temp/pre-ir-redundant > ./temp/iran-hosted-domains-lite
cp ./temp/iran-hosted-domains-lite ./data/std/
cp ./temp/iran-hosted-domains-lite ./data/exp/
cp ./temp/iran-hosted-domains ./data/exp/
