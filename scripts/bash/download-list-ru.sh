#!/usr/bin/env bash

echo "Downloading Domain lists only available in Russia..."
curl -L -o ./temp/available-only-in-ru https://raw.githubusercontent.com/runetfreedom/russia-domains-list/refs/heads/main/ru-available-only-inside
cp ./temp/available-only-in-ru ./data/std/
cp ./temp/available-only-in-ru ./data/exp/
echo "Downloading Re:filter list..."
curl -L -o ./temp/refilter https://raw.githubusercontent.com/1andrevich/Re-filter-lists/refs/heads/main/domains_all.lst
cp ./temp/refilter ./data/std/
cp ./temp/refilter ./data/exp/
echo "Downloading Antifilter list community edition..."
curl -L -o ./temp/antifilter-list-community https://community.antifilter.download/list/domains.lst
cp ./temp/antifilter-list-community ./data/std/
cp ./temp/antifilter-list-community ./data/exp/
echo "Downloading Antifilter list..."
curl -L -o ./temp/pre-antifilter-list-full https://antifilter.download/list/domains.lst
./sanitize --inputpath ./temp/pre-antifilter-list-full --outputname antifilter-list-full --outputdir ./data/exp/
