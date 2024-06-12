#!/bin/sh

echo "================================================"
echo "/ping"
curl --request GET \
  --url http://localhost:12021/ping

echo -e "\n================================================"
echo "/login"
curl --request POST \
  --url http://localhost:12021/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asdx@asd.com",
	"password": "testtest"
}'

echo -e "\n================================================"
echo "/register"
curl --request POST \
  --url http://localhost:12021/register \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asdx@asd.com",
	"password": "testtest"
}'

echo -e "\n================================================"
echo "/login"
curl --request POST \
  --url http://localhost:12021/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asdx@asd.com",
	"password": "testtest"
}'

echo -e "\n================================================"
echo "/login"
curl --request POST \
  --url http://localhost:12021/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asd@asd.com",
	"password": "testtesttest"
}'

echo -e "\n================================================"
echo "/register"
curl --request POST \
  --url http://localhost:12021/register \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asd@asd.com",
	"password": "testtesttest"
}'

echo -e "\n================================================"
echo "/login"
curl --request POST \
  --url http://localhost:12021/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"asd@asd.com",
	"password": "testtesttest"
}'

echo -e "\n================================================"
