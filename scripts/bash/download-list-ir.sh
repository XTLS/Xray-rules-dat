#!/usr/bin/env bash

echo "Downloading Iran related lists..."
curl -L -o ./temp/pre-ir-redundant https://raw.githubusercontent.com/Chocolate4U/Iran-v2ray-rules/refs/heads/redundant/redundant-domains.txt
curl -L -o ./temp/pre-expand-ir https://raw.githubusercontent.com/bootmortis/iran-hosted-domains/refs/heads/release/domains.txt
cat ./temp/pre-expand-ir | grep -Ev ".+\.ir$" | sed '1 a\ir\nxn--mgba3a4f16a' | LC_ALL=C sort -u > ./temp/expand-ir
comm -23 ./temp/expand-ir ./temp/pre-ir-redundant > ./temp/mainly-ir
cp ./temp/mainly-ir ./data/std/
cp ./temp/mainly-ir ./data/exp/
cp ./temp/expand-ir ./data/exp/
