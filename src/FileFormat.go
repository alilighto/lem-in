package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//  fill a struct with file content
func HandleFormat(filename string) (*antcolonie, error) {
	antcolonie := &antcolonie{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, fmt.Errorf("invalid data format")
	}
	ants, err := strconv.Atoi(scanner.Text())
	if ants <= 0 || err != nil {
		return nil, fmt.Errorf("invalid data format")
	}
	antcolonie.Ants = ants
	for scanner.Scan() {
		Line := scanner.Text()
		if Line == "##start" {
			scanner.Scan()
			antcolonie.StartRoom = strings.Fields(scanner.Text())[0]
			continue
		}
		if Line == "##end" {
			scanner.Scan()
			antcolonie.StartRoom = strings.Fields(scanner.Text())[0]
			continue
		}
		if Line == "" || strings.HasPrefix(Line, "#") {
			continue
		}
		if strings.Contains(Line, "-") {
			antcolonie.Edges = append(antcolonie.Edges, Line)
		} else {
			antcolonie.Rooms = append(antcolonie.Rooms, strings.Fields(Line)[0])
		}
	}
	return antcolonie, nil
}
