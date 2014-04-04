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


func connHandler(cn net.Conn) {

	encodeData := new(bytes.Buffer)

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

	var decData Result

	_, err := asn1.Unmarshal(encodeData.Bytes(), &decData)

	if err != nil {
		fmt.Println(err)
	}


	fmt.Print("Decoded\nINTEGER\t\t:")
	for _, v := range decData.Modulus.Bytes() {
		fmt.Printf("%X", v)
	}

	fmt.Print("\nINTEGER\t\t")
	fmt.Printf(":%X\n", decData.Exponent)

	fmt.Println();
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

		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go connHandler(conn)
	}
}
