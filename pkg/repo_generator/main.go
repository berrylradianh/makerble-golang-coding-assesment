/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func main() {
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

	// get struct name from each model
	for _, f := range list {
		file, err := os.ReadFile(dir + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		fileString := string(file)
		var cutPre string
		var clean string
		splits := strings.Split(fileString, "type ")
		if len(splits) > 1 {
			cutPre = splits[1]
		}
		splits = strings.Split(cutPre, " struct")
		if len(splits) > 1 {
			clean = splits[0]
		}
		if clean == "" || clean == "Migration" {
			continue
		}

		replacer(clean)
	}

}

func replacer(structName string) {
	replace := structName
	replaceLow := strings.ToLower(replace[0:1]) + replace[1:]
	replaceUp := replace
	replaceTitle := matchFirstCap.ReplaceAllString(replace, "${1}_${2}")
	replaceTitle = matchAllCap.ReplaceAllString(replaceTitle, "${1}_${2}")
	replaceTitle = strings.ToLower(replaceTitle)

	repoFileName := "pkg/repo_generator/template_repo.gotmp"
	serviceFileName := "pkg/repo_generator/template_service.gotmp"
	err := generateFile(replaceLow, replaceUp, replaceTitle, repoFileName, serviceFileName)
	if err != nil {
		fmt.Println(err)
	}
}

func generateFile(replaceLow, replaceUp, title, repo, service string) error {
	fileRepo, err := os.Open(repo)
	if err != nil {
		fmt.Println(err)
	}
	defer fileRepo.Close()

	resFileRepo, err := os.Create("library/repository/" + title + "_repository.go")
	if err != nil {
		fmt.Println(errors.New("failed creating file, " + err.Error()))
	}
	defer resFileRepo.Close()

	var lineText string
	scanner := bufio.NewScanner(fileRepo)
	for scanner.Scan() {
		lineText = scanner.Text()
		lineText = strings.Replace(lineText, "xxx", replaceUp, -1)
		_, err = resFileRepo.WriteString(lineText + "\n")
		if err != nil {
			fmt.Println(errors.New("failed writing file, " + err.Error()))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fileService, err := os.Open(service)
	if err != nil {
		fmt.Println(err)
	}
	defer fileService.Close()
	_, err = os.ReadFile("library/service/" + title + "_service.go")
	if err == nil {
		fmt.Println("file exists, skipping...")
		return nil
	}
	resFileService, err := os.Create("library/service/" + title + "_service.go")
	if err != nil {
		fmt.Println(errors.New("failed creating file, " + err.Error()))
	}
	defer resFileService.Close()
	scannerService := bufio.NewScanner(fileService)
	for scannerService.Scan() {
		lineText = scannerService.Text()
		lineText = strings.Replace(lineText, "xxx", replaceLow, -1)
		lineText = strings.Replace(lineText, "yyy", replaceUp, -1)
		_, err = resFileService.WriteString(lineText + "\n")
		if err != nil {
			fmt.Println(errors.New("failed writing file, " + err.Error()))
		}
	}

	if err := scannerService.Err(); err != nil {
		fmt.Println(err)
	}
	return nil
}
