#!/bin/bash
env GOOS=linux GOARCH=arm GOARM=5 go build -o read
scp read pi@192.168.0.200:~/Development/arduino