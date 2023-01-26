#!/bin/bash

STR=$(curl -I "$2" 2>/dev/null | head -n 1 | cut -d$' ' -f2)
if [[ ${STR} == *$1* ]]; then
    echo "Check successful"
    exit 0 
fi

echo "Version is not up to date"
exit 1