#!/bin/bash

for i in $(eval echo {1..$1})
do 
  curl --location --request POST 'http://127.0.0.1:9080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Bob", 
    "email": "bob@gmail11.com", 
    "password": "ilovealice"
}'  
done