package connection

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

var Connection *ftp.ServerConn
func Connect() {
	var err error
	Connection, err = ftp.Dial("streetram.net:21066", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("connection stablished.")
	// login to the ftp server
	err = Connection.Login("fahad", "@Dfklpt699")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("LoggedIn as Fahad.")
}