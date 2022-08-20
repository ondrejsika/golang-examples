package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//go:embed data/names-male.json
var NamesMaleJsonFile []byte

//go:embed data/names-female.json
var NamesFemaleJsonFile []byte

//go:embed data/names-surnames.json
var NamesSurnamesJsonFile []byte

type Json struct {
	Data []string `json:"data"`
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(randomEmailPrefix())
}

func randomEmailPrefix() string {
	var j Json
	var maleNames []string
	var femaleNames []string
	var surnames []string

	_ = json.Unmarshal(NamesMaleJsonFile, &j)
	maleNames = j.Data
	_ = json.Unmarshal(NamesFemaleJsonFile, &j)
	femaleNames = j.Data
	_ = json.Unmarshal(NamesSurnamesJsonFile, &j)
	surnames = j.Data

	name := randomChoice(append(maleNames, femaleNames...))
	nameFirst := name[:1]
	middleName := randomChoice(append(maleNames, femaleNames...))
	middleNameFirts := middleName[0:1]
	surname := randomChoice(surnames)
	surnameFirst := surname[:1]
	number := strconv.Itoa(rand.Intn(8) + 1)

	results := []string{
		// john
		name,
		// john1
		fmt.Sprintf("%s%s", name, number),
		// doe
		surname,
		// doe1
		fmt.Sprintf("%s%s", surname, number),

		// johndoe
		fmt.Sprintf("%s%s", name, surname),
		// johndoe1
		fmt.Sprintf("%s%s%s", name, surname, number),
		// john.doe
		fmt.Sprintf("%s.%s", name, surname),
		// john.doe1
		fmt.Sprintf("%s.%s%s", name, surname, number),
		// john.doe.1
		fmt.Sprintf("%s.%s.%s", name, surname, number),

		//doejohn
		fmt.Sprintf("%s%s", surname, name),
		//doejohn1
		fmt.Sprintf("%s%s%s", surname, name, number),
		//doe.john
		fmt.Sprintf("%s.%s", surname, name),

		// jdoe
		fmt.Sprintf("%s%s", nameFirst, surname),
		// jdoe1
		fmt.Sprintf("%s%s%s", nameFirst, surname, number),

		// jdoe
		fmt.Sprintf("%s%s", name, surnameFirst),
		// jdoe1
		fmt.Sprintf("%s.%s", name, surnameFirst),

		// j.doe
		fmt.Sprintf("%s.%s", nameFirst, surname),
		// j.doe1
		fmt.Sprintf("%s.%s%s", nameFirst, surname, number),
		// j.doe.1
		fmt.Sprintf("%s.%s.%s", nameFirst, surname, number),

		// doej
		fmt.Sprintf("%s%s", surname, nameFirst),
		// doej1
		fmt.Sprintf("%s%s%s", surname, nameFirst, number),
		// doe.j
		fmt.Sprintf("%s.%s", surname, nameFirst),

		// j.doe
		fmt.Sprintf("%s.%s", nameFirst, surname),
		// j.doe1
		fmt.Sprintf("%s.%s%s", nameFirst, surname, number),
		// j.doe.1
		fmt.Sprintf("%s.%s.%s", nameFirst, surname, number),

		// jxdoe
		fmt.Sprintf("%s%s%s", nameFirst, middleNameFirts, surname),
		// jxdoe
		fmt.Sprintf("%s.%s.%s", nameFirst, middleNameFirts, surname),
		// jxdoe1
		fmt.Sprintf("%s%s%s%s", nameFirst, middleNameFirts, surname, number),
	}
	return strings.ToLower(randomChoice(results))
}

func randomChoice(data []string) string {
	return data[rand.Intn(len(data))]
}
