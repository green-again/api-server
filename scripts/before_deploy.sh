#!/bin/bash
REPOSITORY=/home/ubuntu
cd $REPOSITORY

docker run -i -t --name ga-api green-again:latest
