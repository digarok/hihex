package main

import "flag"
import "fmt"
import "os"
import "io/ioutil"

func check(e error) {
	if e != nil {
		//panic(e)
		fmt.Printf("Could not find file or other critical error.\nCheck options and try again.\n\n")
		os.Exit(9)
	}
}

func isAscii(b byte) bool {
	return b >= 32 && b < 127
}

func usage() {
	fmt.Println(`
NAME
     hihex -- show hexdump of a file

SYNOPSIS
     hihex [options] [file]

DESCRIPTION
     The hihex utility works like the 'hexdump' command with a special
     addition.  It allows you to view the character output of ASCII 
     values that have the high bit set.

     This is particularly useful for examining binary files from older 
     computers, such as the Apple II, where they would commonly use the 
     high bit versions of characters to display text in other modes such 
     as inverse. 

     The options are as follows: 

     -verbose         Show informational messages when running
     -lowercase       Use lowercase hex values ('0e' vs '0E')
     -hihex           Show high bit ASCII characters (defaults to true)
     -offset={offset} Starting offset address (24-bits, eg "0x002000")
     -chunksize={int} Number of characters to show per line (default 16)


  `)
}

func main() {
	var offset = flag.Int("offset", 0x000000, "24-bit hex address ofset to start at (eg 0x002000)")
	var chunksize = flag.Int("chunksize", 16, "number of hex bytes to display per line")
	var verbose = flag.Bool("verbose", false, "use -verbose to turn on informational messages")
	var lowercase = flag.Bool("lowercase", false, "use -lowercase to print out lower case hex alpha chars")
	var hihex = flag.Bool("hihex", true, "show 'high bit' ASCII characters (kind of the whole point)")
	var help = flag.Bool("?", false, "show usage info")
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" || *help {
		usage()
		os.Exit(1)
	}

	var data, err = ioutil.ReadFile(filename)
	check(err)
	if *verbose == true {
		fmt.Printf("hihex v0.0.0  by  Dagen Brock\n\n")
		fmt.Println("Now processing filename:     ", filename)
		fmt.Println("File Size (decimal):         ", len(data))
		fmt.Printf("File Size (hexadecimal):     %X\n", len(data))
		fmt.Printf("Beginning at offset:          0x%06X\n", *offset)
		fmt.Println("Chunksize (bytes per line):  ", *chunksize)
		fmt.Println("Show 'high bit' ASCII chars: ", *hihex)
		fmt.Println("Lower case hex on:           ", *lowercase)
	}

	i := 0
	for i < len(data) {
		// print offset
		if *lowercase {
			fmt.Printf("%06x ", *offset+i)
		} else {
			fmt.Printf("%06X ", *offset+i)
		}
		j := 0
		for j < *chunksize {
			if i+j < len(data) {
				if *lowercase {
					fmt.Printf(" %02x", data[i+j])
				} else {
					fmt.Printf(" %02X", data[i+j])
				}
			} else {
				fmt.Print(" __")
			}
			j++
		}
		fmt.Print("  ")
		j = 0
		for j < *chunksize {
			if i+j < len(data) {
				if isAscii(data[i+j]) {
					fmt.Print(string(data[i+j]))
				} else if isAscii(data[i+j]&0x7f) && *hihex {
					fmt.Print(string(data[i+j] & 0x7f))
				} else {
					fmt.Print("☐")
				}
			} else {
				fmt.Print(" ")
			}
			j++
		}
		i += *chunksize
		fmt.Println()
	}
}
