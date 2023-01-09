#!/bin/bash

STR=$(curl -s "$2")
if [[ ${STR} == *$1* ]]; then
    echo "Check successful"
    exit 0 
fi

echo "Version is not up to date"
exit 1