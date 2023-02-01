package main

import (
	"sort"
)

func SortByRegions(i, j int) bool {
	return Variables[int(EnergyDataArr[i].IdPozycja1)] > Variables[int(EnergyDataArr[j].IdPozycja1)]
}

func FilterByRegion(region string) []EnergyElement {
	sort.Slice(EnergyDataArr, SortByRegions)
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByCharacter(character string) []EnergyElement {
	sort.Slice(EnergyDataArr, SortByRegions)
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja2)] == character {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByCharacterAndRegion(character, region string) []EnergyElement {
	sort.Slice(EnergyDataArr, SortByRegions)
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja2)] == character && Variables[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}
