package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jasonlvhit/gocron"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

func main() {
	var flag_mode string
	flag.StringVar(&flag_mode, "mode", "server", "tcp server")

	flag.Parse()

	if flag_mode == "client" {
		fmt.Println("Client mode")

		gocron.Every(10).Second().Do(clientMode)

	} else {
		fmt.Printf("Listen TCP: %v:%v\r\n", HOST, PORT)

		listen, err := net.Listen(TYPE, HOST+":"+PORT)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// close listener
		defer listen.Close()
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			// go func() {
			// 	handleIncomingRequest(conn)

			// 	sendHexData(conn, "7E 80 01 00 05 4E 5D E6 6F BA 22 48 BB 00 55 01 02 00 23 7E")
			// }()

			// go handleConnection(conn)

			go handleIncomingRequest(conn)
		}
	}

	// Start all the pending jobs
	<-gocron.Start()
}

// handleIncomingRequest
func handleIncomingRequest(conn net.Conn) {

	data, err := bufio.NewReader(conn).ReadString('\n')
	// data, err := bufio.NewReader(conn).Read()
	if err != nil {
		log.Fatal("get client data error: ", err)
	}

	fmt.Printf("%#v\r\n", data)

	// write to log
	writeLogFile("log_test_tcp.txt", data)

	// close conn
	defer conn.Close()
}

//writeLogFile
func writeLogFile(name string, data_buf string) {
	// write the whole body at once
	err := ioutil.WriteFile(name, []byte(data_buf), 0644)
	if err != nil {
		panic(err)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			m := strings.Fields(line)[0]
			fmt.Println("Methods", m)
		}
		if line == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := ``

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func testReceivedHexData() (string, error) {
	test_data := string("~ N]�o��ODYxNjUyMDUwMDY4NDQwbmV3U25SdWxl�~")
	hx := hex.EncodeToString([]byte(test_data))
	fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(test_data + " ==> " + hx)

	return hx, nil
}

func sendHexData(conn net.Conn, string_hex string) error {
	string_hex = strings.Join(strings.Fields(string_hex), "")
	fmt.Println(string_hex)

	data, err := hex.DecodeString(string_hex)
	if err != nil {
		return err
	}
	// fmt.Println(data)
	fmt.Fprint(conn, data)

	_, err = conn.Write(data)
	if err != nil {
		println("Write to server failed:", err.Error())
		return err
	}

	return nil

}

func clientMode() {
	strEcho, _ := testReceivedHexData()
	servAddr := "gpsdev.tracksolidpro.com:21122"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	// go handleIncomingRequest(conn)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	// println("reply from server=", string(reply))

	defer conn.Close()
}
