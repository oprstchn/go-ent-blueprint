package main

import (
	"os"
	"strings"
)

const protoDir = "./"

func main() {
	files, err := os.ReadDir(protoDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, "_service.go") {
			fileName := protoDir + name
			fb, err := os.ReadFile(fileName)
			if err != nil {
				panic(err)
			}

			replacedFile := strings.ReplaceAll(string(fb), "pageToken := string(bytes)", "pageToken, err := xid.FromBytes(bytes)\n		if err != nil {\n    		return nil, status.Errorf(codes.InvalidArgument, \"page token is invalid\")\n		}")
			if err := os.WriteFile(fileName, []byte(replacedFile), os.FileMode(0755)); err != nil {
				panic(err)
			}
		}
	}
}
