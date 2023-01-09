#!/bin/bash

if ! [ -f "$1"  ] ;
then
        echo "cant find binary file to deploy"
        exit 0
fi

basename "$1"
file="$(basename -- "$1")"

rm -rf /var/www/html/*
cp "$1" /var/www/html/
cd /var/www || exit
tar -xf html/"$file" -C html/
mv html/build/* html/
rm -rf html/build/
rm html/"$file"