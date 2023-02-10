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
	FilterByRegion(region string) (energy []ResponseElement)
	ApplyFilters(filters map[string]string)
}

func (arr *EnergyData) ApplyFilters(filters map[string]string) {
	fmt.Println("Applying filters...")
	for filter, value := range filters {
		switch {
		case filter == "region" && value != "":
			fmt.Println("Filtering by region...")
		case filter == "character" && value != "":
			fmt.Println("Filtering by character...")
		}
	}
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
	//Struct for extracing "data" field from JSON response
	// data := EnergyData{}
	resData, err := io.ReadAll(response.Body) //Reading response body
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}
	//Unmarshalling data field to EnergyData struct
	err = json.Unmarshal([]byte(resData), &arr)
	if err != nil { //Checking for errors
		log.Fatalf("Could extract data: %v", err)
		return nil
	}
	*arr = EnergyData{Energy: arr.Energy} //Saving data.Energy to EnergyDataArr
	return arr.Energy                     //Returning data.Energy arr
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
		arr.Energy = result
	}
	return result
}

// // Method for filtering data by character
// func FilterByCharacter(character string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }

// // Method for filtering data by region and character
// func FilterByCharacterAndRegion(character, region string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character && Regions[int(EnergyDataArr[i].IdPozycja1)] == region {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }

// // QuickSort implementation for sorting data by region in descending order
// func partition(arr []ResponseElement, left, right int) ([]ResponseElement, int) {
// 	//Comparing each request is based on the decoded region from Regions array (variables.go)
// 	pivot := Regions[int(arr[right].IdPozycja1)]
// 	i := left
// 	for j := left; j < right; j++ {
// 		if Regions[int(arr[j].IdPozycja1)] <= pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[right] = arr[right], arr[i]
// 	return arr, i
// }

// func QuickSortByRegion(arr []ResponseElement, left, right int) []ResponseElement {
// 	if left < right {
// 		var p int
// 		arr, p = partition(EnergyDataArr, left, right)
// 		arr = QuickSortByRegion(arr, left, p-1)
// 		arr = QuickSortByRegion(arr, p+1, right)
// 	}
// 	return arr
// }

// // Method for filtering data by type of energy source
// func FilterByTypeOfSource(typeOfEnergy string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Types[int(EnergyDataArr[i].IdPozycja2)] == typeOfEnergy {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }

// func FilterByEnergyUnit(energyUnit string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Units[int(EnergyDataArr[i].IdSposobPrezentacjiMiara)] == energyUnit {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }

// func FilterByTypeAndUnit(typeOfEnergy, energyUnit string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range EnergyDataArr {
// 		if Types[int(EnergyDataArr[i].IdPozycja2)] == typeOfEnergy && Units[int(EnergyDataArr[i].IdSposobPrezentacjiMiara)] == energyUnit {
// 			result = append(result, EnergyDataArr[i])
// 		}
// 	}
// 	return result
// }
