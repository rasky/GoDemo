/*
  GoDemo implements a simple TCP demon that decode
  an ans.1 DER format, and output the decoded  data on
  stdout.
*/

package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"reflect"
)

type Result struct {
	Modulus         *big.Int
	PrivateExponent *big.Int
	PublicExponent  int
}

// connHander it is the called
// goroutine that serve the user requests.
func connHandler(cn net.Conn) {

	// close connection on exit, to notify client
	defer cn.Close()

	// read all data
	data, err := ioutil.ReadAll(cn)
	if err != nil {
		log.Println(err)
		return
	}

	// structured data to store decoded data
	var decData Result

	// compute unmarshal function
	_, err = asn1.Unmarshal(data, &decData)

	if err != nil {
		log.Println(err)
		return
	}

	// render decoded data on stdout
	fmt.Printf("%X\n", decData)
}

func main() {

	fmt.Println("Der decoder service")

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
