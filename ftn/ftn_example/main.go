package main

import (
	"fmt"
	"os"

	"github.com/mrmikhailv/ftdi/ftn"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	dl, err := ftn.FindDevices(0x0403, 0x6001)
	checkErr(err)
	fmt.Println("Found", len(dl), "devices:")
	for i, d := range dl {
		c, err := d.Connect()
		checkErr(err)
		desc, err := c.Description()
		checkErr(err)
		serial, err := c.Serial()
		checkErr(err)
		fmt.Printf("%d: desc='%s' serial='%s'\n", i, desc, serial)
	}
}
