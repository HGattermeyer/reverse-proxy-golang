# Description
Simple Reverse-proxy in Golang.
For each new incoming request, we want to redirect it to just one backend server at a time according to a customizable URL mapping. However, the same incoming URL at a different time can be redirected to another server (you should decide according to which policy; think about the best approach and the pros and cons of it). 
E.g.,
Time T0
From: GET /api/v1.1/car/ferrari
Redirect to: GET hostname.com/api/v1.1/car/ferrari
Time T1
From: GET /api/v1.1/car/ferrari
Redirect to: GET secondary.com/api/v1.1/car/ferrari

# How to run

Run the run_docker.ps1 to run Docker-Compose.

The file [Proxy.postman_collection.json](Proxy.postman_collection.json) has the endpoints and body read to use it on Postman.

The file [upload-file-model_test.txt](upload-file-model_test.txt) contains the information used to upload a file to the app.
