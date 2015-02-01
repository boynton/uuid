package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("usage:\n")
	fmt.Printf("    uuid -h            show this help message\n")
	fmt.Printf("    uuid               generate and show a Version 1 UUID (time based) for the current time\n")
	fmt.Printf("    uuid <url>         generate and show a Version 3 UUID in the URL namespace for the given URL\n")
	fmt.Printf("    uuid <ns> <name>   show Version 3 UUID for namespace/name. The namespace can be a uuid or URL\n")
	os.Exit(1)
}

func main() {
	var u uuid.UUID
	switch len(os.Args) {
	case 1:
		u = uuid.NewUUID()
	case 2:
		if os.Args[1] == "-h" {
			usage()
		}
		u = uuid.NewMD5(uuid.NameSpace_URL, []byte(os.Args[1]))
	case 3:
		ns := uuid.Parse(os.Args[1])
		if ns == nil {
			ns = uuid.NewMD5(uuid.NameSpace_URL, []byte(os.Args[1]))
		}
		u = uuid.NewMD5(ns, []byte(os.Args[2]))
	default:
		usage()
	}
	fmt.Printf("%v\n", u)
}
