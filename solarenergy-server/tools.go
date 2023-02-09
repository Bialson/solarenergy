package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ExtractJSONData(response *http.Response) (energy []ResponseElement, error error) {
	//Struct for extracing "data" field from JSON response
	data := EnergyData{}
	resData, err := ioutil.ReadAll(response.Body) //Reading response body
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}
	//Unmarshalling data field to EnergyData struct
	err = json.Unmarshal([]byte(resData), &data)
	if err != nil { //Checking for errors
		log.Fatalf("Could not unmarshal data: %v", err)
		return data.Energy, err
	}
	return data.Energy, nil //Returning data.Energy arr
}

// Method for filtering data by region
func FilterByRegion(region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

// Method for filtering data by character
func FilterByCharacter(character string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

// Method for filtering data by region and character
func FilterByCharacterAndRegion(character, region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character && Regions[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

// QuickSort implementation for sorting data by region in descending order
func partition(arr []ResponseElement, left, right int) ([]ResponseElement, int) {
	//Comparing each request is based on the decoded region from Regions array (variables.go)
	pivot := Regions[int(arr[right].IdPozycja1)]
	i := left
	for j := left; j < right; j++ {
		if Regions[int(arr[j].IdPozycja1)] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func QuickSortByRegion(arr []ResponseElement, left, right int) []ResponseElement {
	if left < right {
		var p int
		arr, p = partition(EnergyDataArr, left, right)
		arr = QuickSortByRegion(arr, left, p-1)
		arr = QuickSortByRegion(arr, p+1, right)
	}
	return arr
}

// Method for filtering data by type of energy source
func FilterByTypeOfSource(typeOfEnergy string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Types[int(EnergyDataArr[i].IdPozycja2)] == typeOfEnergy {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByEnergyUnit(energyUnit string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Units[int(EnergyDataArr[i].IdSposobPrezentacjiMiara)] == energyUnit {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByTypeAndUnit(typeOfEnergy, energyUnit string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Types[int(EnergyDataArr[i].IdPozycja2)] == typeOfEnergy && Units[int(EnergyDataArr[i].IdSposobPrezentacjiMiara)] == energyUnit {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}
