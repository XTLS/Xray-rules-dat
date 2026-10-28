#!/usr/bin/env bash

echo "Downloading GFWList..."
curl -L -o ./temp/pre-gfwlist-cn https://raw.githubusercontent.com/gfwlist/gfwlist/refs/heads/master/list.txt
./extract-abp --inputpath ./temp/pre-gfwlist-cn --outputname gfwlist-cn --outputdir ./data/std/
./extract-abp --inputpath ./temp/pre-gfwlist-cn --outputname gfwlist-cn --outputdir ./data/exp/
echo "Downloading dnsmasq-china-list..."
curl -L -o ./temp/pre-acc-cn https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/refs/heads/master/accelerated-domains.china.conf
./extract-dnsmasq --inputpath ./temp/pre-acc-cn --outputname dnsmasq-china-list --outputdir ./data/exp/
echo "Downloading GreatFire.org List..."
curl -L -o ./temp/greatfire-list-cn https://raw.githubusercontent.com/Loyalsoldier/cn-blocked-domain/refs/heads/release/domains.txt
cp ./temp/greatfire-list-cn ./data/exp/
