# genpasswd

[![Build Status](https://travis-ci.com/mingcheng/genpasswd.go.svg?branch=master)](https://travis-ci.com/mingcheng/genpasswd.go)

This is a simple password generator written in golang. This tool is based and inspired from https://github.com/sethvargo/go-password , thanks [Vargo](https://github.com/sethvargo).

## Some improvements

* A simple command interface for users
* Clipboard support for Linux(Xorg/Wayland), macOS, Windows etc.

## Build and Usage

Just run `make build` to build genpasswd(require go version > 1.11). The `genpasswd` has many params, including:

```
  -allow-repeat
    	Allows characters to repeat (default true)
  -copies int
    	Generate copies of password (default 1)
  -digits int
    	Number of digits to include in the password (default 4)
  -length int
    	Specify the password length (default 16)
  -no-clipboard
    	Do not copy to clipboard
  -no-upper
    	Excludes uppercase letters from the results (default true)
  -symbols int
    	Number of digits to include in the password (default 4)
  -version
    	Print version and exit
```

### Running in Docker

Run `make docker_image` to build Docker image for gopasswd, if u wanna running this tool in contrainer.

Then, type  `docker run --rm genpasswd:0.0.1` to run genpasswd after build is completed.

That's it, have fun!

--

## README.md from original `go-password` library.

[![GoDoc](https://godoc.org/github.com/sethvargo/go-password?status.svg)](https://godoc.org/github.com/sethvargo/go-password)

This library implements generation of random passwords with provided
requirements as described by  [AgileBits
1Password](https://discussions.agilebits.com/discussion/23842/how-random-are-the-generated-passwords)
in pure Golang. The algorithm is commonly used when generating website
passwords.

The library uses crypto/rand for added randomness.

Sample example passwords this library may generate:

```text
0N[k9PhDqmmfaO`p_XHjVv`HTq|zsH4XiH8umjg9JAGJ#\Qm6lZ,28XF4{X?3sHj
7@90|0H7!4p\,c<!32:)0.9N
UlYuRtgqyWEivlXnLeBpZvIQ
Q795Im1VR5h363s48oZGaLDa
wpvbxlsc
```

## Installation

```sh
$ go get -u github.com/sethvargo/go-password/password
```

## Usage

```golang
package main

import (
  "log"

  "github.com/sethvargo/go-password/password"
)

func main() {
  // Generate a password that is 64 characters long with 10 digits, 10 symbols,
  // allowing upper and lower case letters, disallowing repeat characters.
  res, err := password.Generate(64, 10, 10, false, false)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf(res)
}
```

See the [GoDoc](https://godoc.org/github.com/sethvargo/go-password) for more information.

## License

This code is licensed under the MIT license.
