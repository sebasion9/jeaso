package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"jesao++/core"
)


func getActions() []string {
	return []string{"sort", "reverse"}
}

func help() {
	fmt.Println("usage: jesao++ <file.json> <action> [action_arg]")
	fmt.Println("<file.json>\t-\tinput json file to process")
	fmt.Println("<action>\t-\tsort|reverse")
	fmt.Println("[action_arg]\t-\toptional argument for action")
}

func read_fi(path string) (error, []byte) {
	var err error = nil
	fi, err := os.Open(path)
	
	// close the file on exit
	defer func() {
		err = fi.Close()
	}()

	var out_str string = ""
	buf := make([]byte, 1024)
	for {
		_, err = fi.Read(buf)
		if err == io.EOF {
			break
		}
		out_str += string(buf)
	}
	out_str = strings.Trim(out_str, "\x00")
	if err != nil && err == io.EOF {
		return nil, []byte(out_str)
	}
	return err, []byte(out_str)
}


// reads json, does stuff, returns json

func main() {
	// parse args
	args := os.Args
	if len(args) < 3 {
		help()
		return
	}
	path := args[1]
	action := args[2]
	// action_args will be ...args, for example for sort (target, key, asc|desc)
	action_arg := ""
	if !slices.Contains(getActions(), action) {
		fmt.Printf("%s is invalid action\n", action)
		help()
		return
	}
	if len(args) == 4 {
		action_arg = args[3]
	}

	// read json

	err, body := read_fi(path)
	if err != nil {
		fmt.Printf("failed to read file at %s\n", path)
		fmt.Printf("err: %v\n", err)
		return
	}

	// deserialize top level array or object

	var jcore core.JSONCore

	jcore.UnmarshalJSON(body)
	if err != nil {
		fmt.Println("failed to deserialize json")
		fmt.Printf("err: %v\n", err)
		return
	}

	// do stuff with json

	switch action {
	case "sort":
		jcore.Sort(action_arg, "", "")
	case "reverse":
		//core.Reverse()
	}

	// serialize

	data, err := jcore.MarshalJSON()
	if err != nil {
		fmt.Println("failed to serialize json")
	}
	fmt.Println(string(data))
}


