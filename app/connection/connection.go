package connection

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

var Connection *ftp.ServerConn

func Connect() {
	var err error
	addr := "streetram.net:21066"
	user := "fahad"
	pass := "@Dfklpt699"
	Connection, err = ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("Error when trying to connect with ftp server.")
		fmt.Println("Failed to connect with FTP server(" + addr + ").")
		panic(err)
	}
	fmt.Println("connection stablished.")

	// login to the ftp server
	err = Connection.Login(user, pass)
	if err != nil {
		fmt.Println("Error when trying to login on ftp server.")
		fmt.Println("Failed to login on FTP server(" + addr + ") as (" + user + ").")
		panic(err)
	}
	fmt.Println("LoggedIn as: "+user)
}
