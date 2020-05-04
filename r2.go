package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func initGolang(dirPath string, path string, name string) {
	fmt.Print(dirPath + "\n")
	fmt.Print(path + "\n")
	fmt.Print(name + "\n")
	args := []string{name}
	cmd := exec.Command("bash /mnt/c/Users/snobo/Documents/GitHub/r2/golang.sh", args...)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("%s\n", out)
	/*if runtime.GOOS == "windows" {
		exec.Command("cd " + path)
		exec.Command("go mod init " + path + name)
	}*/
}

func main() {
	pkgTypes := "golang|python|java|node.js"

	pkgType := flag.String("type", "", "Pakages Types: {"+pkgTypes+"};")
	pkgName := flag.String("name", "", "Name of the package to init: {example.com/app")
	dirName := flag.String("dir", "", "Name of Dir to create (Required)")
	dirPath := flag.String("path", "path", "Path {C:\\Windows\\Users\\example...}; (Required)")
	flag.Parse()

	if *dirName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	os.Mkdir(*dirPath+"\\"+*dirName, os.ModePerm)

	switch *pkgType {
	case "golang":
		initGolang(*dirPath, *dirPath+"\\"+*dirName, *pkgName)
	}
}
