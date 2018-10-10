package main

import "fmt"

func main() {
	/*
		elements := map[string]string{
			"H":  "Hydrogen",
			"He": "Helium",
			"Li": "Lithium",
			"Be": "Beryllium",
			"B":  "Boron",
			"C":  "Carbon",
			"N":  "Nitrogen",
			"O":  "Oxygen",
			"F":  "Fluorine",
			"Ne": "Neon",
		}*/

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	/*
		elements := make(map[string]string)

		elements["H"] = "Hydrogen"
		elements["He"] = "Helium"
		elements["Li"] = "Lithium"
		elements["Be"] = "Beryllium"
		elements["B"] = "Boron"
		elements["C"] = "Carbon"
		elements["N"] = "Nitrogen"
		elements["O"] = "Oxygen"
		elements["F"] = "Fluorine"
		elements["Ne"] = "Neon"

		name, ok := elements["Li"]
		fmt.Println(name, ok)

		if name, ok = elements["Un"]; ok {
			fmt.Println(name, ok)
		}
	*/

}
