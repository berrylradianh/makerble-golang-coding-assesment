/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// generate models
	u := flag.String("u", "user", "user")
	p := flag.String("p", "", "password")
	d := flag.String("d", "clinic-portal-makerble-golang-test", "database")
	t := flag.String("t", "yourtables", "tables comma separated")
	db := flag.String("db", "localhost", "database ip")

	flag.Parse()

	cmd := exec.Command(
		"gentool",
		"-dsn",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", *u, *p, *db, *d),
		"-db="+*db,
		"-outPath=./library/model/",
		"-modelPkgName=model",
		"-onlyModel",
		"-fieldNullable",
		"-tables="+*t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// get all filenames in a folder
	dir := "./library/model/"
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range list {
		file, err := os.ReadFile(dir + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		fileString := string(file)
		var clean string
		splits := strings.Split(fileString, "package model")
		if len(splits) > 1 {
			clean = "package model" + splits[1]
		}

		clean = strings.ReplaceAll(clean, "int32", "int")

		err = os.WriteFile(dir+f.Name(), []byte(clean), 0)
		if err != nil {
			fmt.Println(err)
		}

		name := strings.TrimSuffix(f.Name(), "gen.go")
		if name == f.Name() { // no suffix so no need to rename
			continue
		}
		err = os.Rename(dir+f.Name(), dir+name+"go")
		if err != nil {
			fmt.Println(err)
		}
	}
}
