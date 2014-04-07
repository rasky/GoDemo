/*
  GoDemo implements a simple TCP demon that decode
  an ans.1 DER formt, and output the decoded  data on
  stdout.
*/

package main


import(
	"fmt"
	"log"
	"net"
	"bytes"
	"math/big"
	"encoding/asn1"
)


type Result struct {
	Modulus *big.Int
	Exponent int
}

// connHander it is the called
// goroutine that serve the user requests.
func connHandler(cn net.Conn) {

	// alloc buffer for incoming data on socket
	encodeData := new(bytes.Buffer)

	// data is a simple slice to collect partial chunks
	data := make([]byte, 512)

	for {
		n, err := cn.Read(data)

		if n > 0 {
			encodeData.Write(data[:n])
		}

		if err != nil {
			break;
		}

	}

	// structred data to store decoded data
	var decData Result

	// compute unmarchal function
	_, err := asn1.Unmarshal(encodeData.Bytes(), &decData)

	if err != nil {
		fmt.Println(err)
		return
	}

	// render decoded data on stdout

	fmt.Print("Decoded\nINTEGER\t\t:")
	// print modulus
	for _, v := range decData.Modulus.Bytes() {
		fmt.Printf("%X", v)
	}

	// render exponent
	fmt.Print("\nINTEGER\t\t")
	fmt.Printf(":%X\n", decData.Exponent)

	fmt.Println();

	// close connection to notify client
	cn.Close();
}

func main() {

	fmt.Println("Der decoder service");

	ln, err := net.Listen("tcp", ":4000")

	if err != nil {
		fmt.Println("Busy port.")
	}

	defer ln.Close()

	for {
		// wait for incaming connection
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}

		// serve request
		go connHandler(conn)
	}
}
