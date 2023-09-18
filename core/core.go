package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

var assocs map[string]Association

type Association interface {
	CanAdd() bool
}

func init() {
	paths := Files("../plugins")
	assocs = make(map[string]Association)

	var (
		plug *plugin.Plugin
		err  error
	)

	for _, p := range paths {
		fmt.Println(p)
		// load module
		// 1. open the so file to load the symbols
		plug, err = plugin.Open("/home/alok/workspace/github.com/aloksinhanov/go-hello-plugin/plugins/" + p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Association
		symAssoc, err := plug.Lookup("Association")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Greeter (defined above)
		association, ok := symAssoc.(Association)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}
		idx := strings.LastIndex(p, "/")
		key := p[idx+1 : len(p)-3]
		fmt.Printf("\nkey: %v", key)
		assocs[key] = association
	}
}

func main() {
	ch := "NA"
	for ch != "" {
		fmt.Println("\nEnter relation to check if can be added: ")
		fmt.Scanln(&ch)
		if assoc, found := assocs[ch]; found {
			fmt.Println(assoc.CanAdd())
		} else {
			fmt.Println("Unknown relation.")
		}
	}

	return
}

func Files(path string) (paths []string) {
	fsys := os.DirFS(path)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".so" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
