package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "81.20.50.204:23")
	if err != nil {
		log.Println("tcp conn", err.Error())
	}
	defer conn.Close()

	time.Sleep(time.Second * 1)

	// Per standard, roll up row must start with:
	// Roll Up Type
	// Carriage Return
	// 'PAC' "to set indent and base row

	//To relocate the roll-up display immediately, send a PAC which implies the new base row. To continuepainting the current row, assuming the relocation occurred in the middle of a row, the PAC used shouldindent to the same column number where the cursor resided prior to the move or to the nearest indent to theleft of that column. The service provider should then either re-send enough characters to return the cursor toits original column, or use the appropriate Tab Offset command to move the cursor there.

	//

	conn.Write([]byte("\x03\x013 f1\x0D")) // Enter Pass Thru mode on the encoder
	conn.Write([]byte("\x14\x25\x14\x2d\x14\x60TEST CAPTION THAT IS LONG......!"))
	time.Sleep(time.Second)
	conn.Write([]byte("\x14\x25\x14\x2d\x14\x60\x10\x29another line in red \x10\x2e"))
	time.Sleep(time.Second)
	conn.Write([]byte("\x14\x25\x14\x2d\x14\x66a fourth line in cyan"))
	time.Sleep(time.Second)
	conn.Write([]byte("\x14\x25\x14\x2d\x14\x60This changes\x11\x2acolour\x11\x20to\x11\x22yellow"))
	time.Sleep(time.Second)
	// Special Characters
	conn.Write([]byte("\x14\x25\x14\x2d\x14\x60*| u\x12\x25\x1a\x25c\x12\x2b-\x12\x2a<\x12\x3e+\x13\x3c+\x13\x3d+\x13\x3e+\x13\x3f \x11\x37 \x11\x36 \x11\x32 \x11\x31 \x11\x35 \x7f"))
	time.Sleep(time.Second * 5)
	conn.Write([]byte("\x03"))

}
