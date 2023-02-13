package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type EnergyData struct {
	Energy []ResponseElement `json:"data"`
}

type Energy interface {
	RequestDBWData(year, cat, section int64) *http.Response
	ExtractJSONData(response *http.Response) (energy []ResponseElement)
	FilterByRegion(region string) []ResponseElement
	FilterByCharacter(character string) []ResponseElement
	FilterByType(energyType string) []ResponseElement
	FilterByUnit(unit string) []ResponseElement
	ApplyFilters(filters map[string]string, amount int64)
	SortByRegion(left, right int) []ResponseElement
}

func (arr *EnergyData) ApplyFilters(filters map[string]string, amount int64) {
	fmt.Println("Applying filters...")
	for filter, value := range filters {
		switch {
		case filter == "region" && value != "":
			fmt.Printf("Filtering by region... [%v]\n", value)
			EnergyService.FilterByRegion(value)
		case filter == "character" && value != "":
			fmt.Printf("Filtering by character... [%v]\n", value)
			EnergyService.FilterByCharacter(value)
		case filter == "type":
			fmt.Printf("Filtering by type... [%v]\n", value)
			EnergyService.FilterByType(value)
		case filter == "unit":
			fmt.Printf("Filtering by unit... [%v]\n", value)
			EnergyService.FilterByUnit(value)
		}
	}
	if amount > int64(len(arr.Energy)) {
		amount = int64(len(arr.Energy))
	}
	*arr = EnergyData{Energy: arr.Energy[:amount]}
	log.Printf("Filtered data count: %v", len(EnergyDataArr.Energy))
}

func (arr *EnergyData) RequestDBWData(year, cat, section int64) *http.Response {
	//Requesting data from DBW API
	url := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?sorts=id-pozycja-2&id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", cat, section, year, PERIOD, MAX_RESULTS)
	log.Printf("Requesting data from: %s", url)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Could not get data: %v", err)
		return nil
	}
	return response
}

func (arr *EnergyData) ExtractJSONData(response *http.Response) (energy []ResponseElement) {
	resData, err := io.ReadAll(response.Body) //Reading response body
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}
	err = json.Unmarshal([]byte(resData), &arr)
	if err != nil { //Checking for errors
		log.Fatalf("Could extract data: %v", err)
		return nil
	}
	*arr = EnergyData{Energy: arr.Energy}
	return arr.Energy
}

// Method for filtering data by region
func (arr *EnergyData) FilterByRegion(region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range arr.Energy {
		if Regions[int(arr.Energy[i].IdPozycja1)] == region {
			result = append(result, arr.Energy[i])
		}
	}
	if len(result) > 0 {
		*arr = EnergyData{Energy: result}
	}
	return result
}

// Method for filtering data by character
func (arr *EnergyData) FilterByCharacter(character string) []ResponseElement {
	result := []ResponseElement{}
	for i := range arr.Energy {
		if Regions[int(arr.Energy[i].IdPozycja2)] == character {
			result = append(result, arr.Energy[i])
		}
	}
	if len(result) > 0 {
		*arr = EnergyData{Energy: result}
	}
	return result
}

// // Method for filtering data by type of energy source
func (arr *EnergyData) FilterByType(typeOfEnergy string) []ResponseElement {
	result := []ResponseElement{}
	for i := range arr.Energy {
		if Types[int(arr.Energy[i].IdPozycja2)] == typeOfEnergy {
			result = append(result, arr.Energy[i])
		}
	}
	if len(result) > 0 {
		*arr = EnergyData{Energy: result}
	}
	return result
}

func (arr *EnergyData) FilterByUnit(energyUnit string) []ResponseElement {
	result := []ResponseElement{}
	for i := range arr.Energy {
		if Units[int(arr.Energy[i].IdSposobPrezentacjiMiara)] == energyUnit {
			result = append(result, arr.Energy[i])
		}
	}
	if len(result) > 0 {
		*arr = EnergyData{Energy: result}
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

func (arr *EnergyData) SortByRegion(left, right int) []ResponseElement {
	if left < right {
		var p int
		arr.Energy, p = partition(arr.Energy, left, right)
		arr.Energy = EnergyService.SortByRegion(left, p-1)
		arr.Energy = EnergyService.SortByRegion(p+1, right)
	}
	return arr.Energy
}

// func FilterByTypeAndUnit(typeOfEnergy, energyUnit string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Types[int(EnergyDataArr[i].IdPozycja2)] == typeOfEnergy && Units[int(EnergyDataArr[i].IdSposobPrezentacjiMiara)] == energyUnit {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }
