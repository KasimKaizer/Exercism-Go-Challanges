// Package grep contains solution for Grep exercise on Exercism.
package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search takes a pattern and a list of files and returns all the lines that have a match for the
// pattern in all those files.
// it also takes a list of flags, and modifies the search method based on them.
func Search(pattern string, flags, files []string) []string {

	// store all the input flags in form of a map.
	flagMap := make(map[string]bool)
	for _, flag := range flags {
		flagMap[flag] = true
	}
	if len(files) > 1 {
		flagMap["-m"] = true // if there are more then 1 files, add -m (multiple files) flag
	}
	output := make([]string, 0)
	for _, file := range files {
		// send the file to searchFile function and append whatever output we get to output slice.
		output = append(output, searchFile(flagMap, pattern, file)...)
	}
	return output
}

// searchFile takes a pattern and a file and searches the file for all the instances of that pattern.
// it also takes a map which contains the flags, and modifies the search method based on them.
func searchFile(flagMap map[string]bool, pattern, file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	output := make([]string, 0)

	if _, ok := flagMap["-i"]; ok {
		// if we have '-i' flag then convert pattern to all lower case.
		pattern = strings.ToLower(pattern)
	}

	// scan thorough the file, line by line.
	for lineNum := 1; scanner.Scan(); lineNum++ {
		curLine := scanner.Text() // get the text from the current line.
		toCheck, isMatch := curLine, false

		// if we have '-i' flag then convert line to lowercase.
		if _, ok := flagMap["-i"]; ok {
			toCheck = strings.ToLower(toCheck)
		}

		// if we have '-x' flag then check if whole line is equal to the pattern.
		if _, ok := flagMap["-x"]; ok {
			isMatch = (pattern == toCheck)
		} else { // if '-x' is not present then check if pattern is present in the line.
			isMatch = strings.Contains(toCheck, pattern)
		}

		// if '-v' flag is present then just invert the isMatch, as we match for lines where
		// pattern is not present.
		if _, ok := flagMap["-v"]; ok {
			isMatch = !isMatch
		}

		// if its not a match then move on to next line.
		if !isMatch {
			continue
		}

		// checks for matched cases from here on.

		// if '-l' is present then just return the file name.
		if _, ok := flagMap["-l"]; ok {
			return []string{file}
		}

		// if "-n" flag is present then add the line number to the current line.
		if _, ok := flagMap["-n"]; ok {
			curLine = fmt.Sprintf("%d:%s", lineNum, curLine)
		}

		// if '-m' flag is present then add the file name to the current line.
		if _, ok := flagMap["-m"]; ok {
			curLine = fmt.Sprintf("%s:%s", file, curLine)
		}

		// append the current line to output.
		output = append(output, curLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}
