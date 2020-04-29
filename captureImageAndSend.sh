#!/bin/sh
cd ~/Pictures/
raspistill -o "$(date +"%Y_%m_%d_%I_%M_%p").jpg"
if [ $? -eq 0 ]
then 
		echo "Captured image "  "$(date +"%Y_%m_%d_%I_%M_%p").jpg"
			echo "Sending to Joey"
				
				rsync -avP "$(date +"%Y_%m_%d_%I_%M_%p").jpg" joey@192.168.0.107:~/Pictures/pi/
			else
					echo "Capture failed!"
				fi
