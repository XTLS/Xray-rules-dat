#!/usr/bin/env bash

echo "Downloading Iran white lists..."
curl -L -o ./temp/pre-ir-redundant https://raw.githubusercontent.com/Chocolate4U/Iran-v2ray-rules/refs/heads/redundant/redundant-domains.txt
curl -L -o ./temp/pre-iran-hosted-domains https://raw.githubusercontent.com/bootmortis/iran-hosted-domains/refs/heads/release/domains.txt
cat ./temp/pre-iran-hosted-domains | grep -Ev ".+\.ir$" | sed '1 a\ir\nxn--mgba3a4f16a' | LC_ALL=C sort -u > ./temp/iran-hosted-domains
comm -23 ./temp/iran-hosted-domains ./temp/pre-ir-redundant > ./temp/iran-hosted-domains-lite
cp ./temp/iran-hosted-domains-lite ./data/std/
cp ./temp/iran-hosted-domains-lite ./data/exp/
cp ./temp/iran-hosted-domains ./data/exp/
echo "Downloading Iran black lists..."
curl -L -o ./temp/dynx-antiban-domains https://raw.githubusercontent.com/MrDevAnony/DynX-AntiBan-Domains/refs/heads/main/DynX-AntiBan-list.lst
cp ./temp/dynx-antiban-domains ./data/std/
cp ./temp/dynx-antiban-domains ./data/exp/
curl -L -o ./temp/iran-sanctioned-domains https://raw.githubusercontent.com/armanlife/iran-sanctioned-domains/refs/heads/main/data/ir-sanctioned-domain
cp ./temp/iran-sanctioned-domains ./data/std/
cp ./temp/iran-sanctioned-domains ./data/exp/
curl -L -o ./temp/pre-iran-blocked-domains https://raw.githubusercontent.com/armanlife/iran-blocked-domains/refs/heads/main/data/ir-blocked
./sanitize --inputpath ./temp/pre-iran-blocked-domains --outputname iran-blocked-domains -outputdir ./data/exp/

