#!/bin/bash

curl --location --request POST "https://api.airbrake.io/api/v4/projects/""$1""/deploys?key=""$2""" \
--header 'Content-Type: application/json' \
--data-raw '{
    "environment": "'"$3"'",
    "username": "deploy-script",
    "repository": "https://github.com/leonidasdeim/chasse",
    "revision": "'"$4"'"
}'