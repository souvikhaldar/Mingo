package main

import (
	"bytes"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

func main() {
	port := flag.String("port", "/dev/ttyACM0", "The port at which device is connected. For linux it is /dev/ttyACM0")
	c := &serial.Config{
		Name: *port,
		Baud: 9600,
	}
	device, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal("Unable to connect to the device: ", err)
	}
	defer device.Close()
	for {
		buffer := make([]byte, 128)
		n, err := device.Read(buffer)
		if err != nil {
			log.Println("Unable to read: ", err)
		}

		if len(bytes.Trim(buffer[:n], "\r")) < 1 {
			continue
		}
		log.Println("Buffer: ", string(buffer))
		log.Println("Position of first null char: ", bytes.IndexAny(buffer, "\r"))
		log.Println("Position of first null char till n: ", bytes.IndexAny(buffer[:n], "\r"))
		log.Println("Length till EOF: ", len(bytes.Trim(buffer[:n], "\r")))
		//log.Println("First 3 bytes are: ", string(buffer[:3]))

		log.Println("Test Distance: ", string(bytes.Trim(buffer[:n], "\r")))

		distance, err := strconv.Atoi(string(bytes.Trim(buffer[:n], "\r")))
		if err != nil {
			log.Println("Unable to convert to int: ", err)
			continue
		}
		log.Println("Distance: ", distance)
		// if distance < 5 {
		// 	log.Println("Object is within boundary, capturing image")
		// 	out, err := exec.Command("/bin/sh", "/home/pi/Development/arduino/captureImageAndSend.sh").Output()
		// 	if err != nil {
		// 		log.Println("Unable to capture image: ", err)
		// 		continue
		// 	}
		// 	log.Println("Output of command: ", string(out))
		// }
		time.Sleep(5 * time.Second)
	}

}
