package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/mingcheng/genpasswd.go"
)

var passwords []string

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
	// parse command line
	flag.Parse()

	// print version and exit
	if *printVersion {
		fmt.Printf("Version %s, build %s\n", version, buildTime)
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
