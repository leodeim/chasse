#!/bin/bash

if ! [ -f "$1"  ] ;
then
	echo "cant find binary file to deploy"
	exit 0
fi

basename "$1"
file="$(basename -- "$1")"

systemctl stop chasse-api.service

if [ -f go/chasse.default.json  ] ; then jq -c --indent 4 --arg ver "$file" '.version = $ver' go/chasse.default.json > tmp.$$.json && mv tmp.$$.json go/chasse.default.json ; fi
if [ -f go/chasse.json  ] ; then jq -c --indent 4 --arg ver "$file" '.version = $ver' go/chasse.json > tmp.$$.json && mv tmp.$$.json go/chasse.json ; fi

ln -sf ../"$1" go/chasse-api
systemctl start chasse-api.service
