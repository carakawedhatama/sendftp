package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	// global variable
	var (
		ftpHost      string // your FTP Host (ip address or domain)
		ftpPort      string // your FTP port
		ftpUser      string // your FTP username
		ftpPassw     string // your FTP password
		ftpSharedDir string // your target folder on FTP Server
		ftpFileName  string // your file name
		destination  string // your destination folder and file on FTP Server
	)

	//initialize value
	ftpHost = "192.168.0.10" // or "yourDomain.com"
	ftpPort = "21"
	ftpUser = "YourUserName"
	ftpPassw = "YourPassword"
	ftpSharedDir = "/path/to/dir/ftpserver/"
	ftpFileName = "yourFile.extension"
	destination = "./" + ftpSharedDir + ftpFileName

	//connecting to FTP
	connect, err := ftp.Dial(ftpHost + ":" + ftpPort)
	if err != nil {
		log.Fatal(err)
	}

	//login to FTP Server
	err = connect.Login(ftpUser, ftpPassw)
	if err != nil {
		log.Fatal(err)
	}

	//prepare file to upload
	file, err := os.Open("./" + ftpFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//upload file to destination folder
	err = connect.Stor(destination, file)
	if err != nil {
		log.Fatal(err)
	}

	//close connection to FTP Server
	connect.Quit()

	uploadTime := time.Now()
	fmt.Println("Upload completed to " + destination + " at " + uploadTime.String())
}
