package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
	genpasswd "github.com/mingcheng/genpasswd.go"
)

const banner = `
                                           o
                                           |
o--o o-o o-o  o-o   oo o-o o-o o   o   o o-O
|  | |-' |  | |  | | |  \   \   \ / \ / |  |
o--O o-o o  o O-o  o-o-o-o o-o   o   o   o-o
   |          |
o--o          o

Simple password generator, version %s, build %s
`

var (
	length       = flag.Int("length", 16, "Specify the password length")
	numDigits    = flag.Int("digits", 4, "Number of digits to include in the password")
	numSymbols   = flag.Int("symbols", 4, "Number of digits to include in the password")
	noUpper      = flag.Bool("no-upper", true, "Excludes uppercase letters from the results")
	allowRepeat  = flag.Bool("allow-repeat", true, "Allows characters to repeat")
	loops        = flag.Int("copies", 1, "Generate copies of password")
	noClipboard  = flag.Bool("no-clipboard", false, "Do not copy to clipboard")
	printVersion = flag.Bool("version", false, "Print version and exit")
	version      string
	buildTime    string
)

func main() {
	var passwords []string

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(banner, version, buildTime))
		flag.PrintDefaults()
	}

	// parse command line
	flag.Parse()

	// print version and exit
	if *printVersion {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(banner, version, buildTime))
		os.Exit(0)
	}

	// generate passwords by loop times
	for i := 0; i < *loops; i++ {
		res, err := genpasswd.Generate(*length, *numDigits, *numSymbols, *noUpper, *allowRepeat)
		if err != nil {
			log.Fatal(err)
		}

		passwords = append(passwords, res)
	}

	// @see https://stackoverflow.com/questions/28799110/how-to-join-a-slice-into-a-string
	result := strings.Join(passwords, "\n")

	// copy passwords to clipboard
	isInLinuxTerminal := runtime.GOOS == "linux" && os.Getenv("DISPLAY") == ""
	if !isInLinuxTerminal && *noClipboard != true {
		if err := clipboard.WriteAll(result); err != nil {
			log.Fatal(err)
		}
	}

	// print generated passwords at last
	fmt.Println(result)
}
