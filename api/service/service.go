package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Bialson/solarenergy/assets"
	api "github.com/Bialson/solarenergy/proto"
	"github.com/buger/jsonparser"
)

var API App = &Data{}

type App interface {
	RequestDBWData(year, cat, section int64) *http.Response
	ExtractJSONData(response *http.Response) error
	SendStreamData(stream interface{}) error
	FilterByRegion(region string) []ResponseElement
	FilterByCharacter(character string) []ResponseElement
	FilterByType(energyType string) []ResponseElement
	FilterByUnit(unit string) []ResponseElement
	ApplyFilters(filters map[string]string, amount int64)
	SortByRegion(left, right int) []ResponseElement
}

func (dataArray *Data) RequestDBWData(year, cat, section int64) *http.Response {
	//Requesting data from DBW API
	url := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?sorts=id-pozycja-2&id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", cat, section, year, assets.PERIOD, 2)
	log.Printf("Requesting data from: %s", url)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Could not get data: %v", err)
		return nil
	}
	return response
}

func (dataArray *Data) ExtractJSONData(response *http.Response) error {
	resData, err := io.ReadAll(response.Body) //Reading response body
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
		return err
	}
	jsonData, _, _, err := jsonparser.Get(resData, "data")
	if err != nil { //Checking for errors
		log.Fatalf("Could not extract data: %v", err)
		return err
	}
	err = json.Unmarshal(jsonData, &dataArray.Energy)
	if err != nil { //Checking for errors
		log.Fatalf("Could not extract data: %v", err)
		return err
	}
	return nil
}

func (dataArray *Data) SendStreamData(stream interface{}) error {
	switch stream := stream.(type) {
	case api.SolarService_GetEnergyFromHomesByParamsServer:
		err := EnergyConsumptionStream(dataArray.Energy, stream)
		if err != nil {
			return err
		}
	case api.SolarService_GetEcoEnergyByParamsServer:
		return fmt.Errorf("unimplemented response")
	default:
		return fmt.Errorf("unsupported data type")
	}
	return nil
}

func (dataArray *Data) ApplyFilters(filters map[string]string, amount int64) {
	fmt.Println("Applying filters...")
	for filter, value := range filters {
		switch {
		case filter == "region" && value != "":
			fmt.Printf("Filtering by region... [%v]\n", value)
			dataArray.FilterByRegion(value)
		case filter == "character" && value != "":
			fmt.Printf("Filtering by character... [%v]\n", value)
			dataArray.FilterByCharacter(value)
		case filter == "type":
			fmt.Printf("Filtering by type... [%v]\n", value)
			dataArray.FilterByType(value)
		case filter == "unit":
			fmt.Printf("Filtering by unit... [%v]\n", value)
			dataArray.FilterByUnit(value)
		}
	}
	if amount > int64(len(dataArray.Energy)) {
		amount = int64(len(dataArray.Energy))
	}
	*dataArray = Data{Energy: dataArray.Energy[:amount]}
	log.Printf("Filtered data count: %v", len(dataArray.Energy))
}

// Method for filtering data by region
func (dataArray *Data) FilterByRegion(region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range dataArray.Energy {
		if assets.Regions[int(dataArray.Energy[i].IdPozycja1)] == region {
			result = append(result, dataArray.Energy[i])
		}
	}
	if len(result) > 0 {
		*dataArray = Data{Energy: result}
	}
	return result
}

// Method for filtering data by character
func (dataArray *Data) FilterByCharacter(character string) []ResponseElement {
	result := []ResponseElement{}
	for i := range dataArray.Energy {
		if assets.Regions[int(dataArray.Energy[i].IdPozycja2)] == character {
			result = append(result, dataArray.Energy[i])
		}
	}
	if len(result) > 0 {
		*dataArray = Data{Energy: result}
	}
	return result
}

// // Method for filtering data by type of energy source
func (dataArray *Data) FilterByType(typeOfEnergy string) []ResponseElement {
	result := []ResponseElement{}
	for i := range dataArray.Energy {
		if assets.EnergyTypes[int(dataArray.Energy[i].IdPozycja2)] == typeOfEnergy {
			result = append(result, dataArray.Energy[i])
		}
	}
	if len(result) > 0 {
		*dataArray = Data{Energy: result}
	}
	return result
}

func (dataArray *Data) FilterByUnit(energyUnit string) []ResponseElement {
	result := []ResponseElement{}
	for i := range dataArray.Energy {
		if assets.Units[int(dataArray.Energy[i].IdSposobPrezentacjiMiara)] == energyUnit {
			result = append(result, dataArray.Energy[i])
		}
	}
	if len(result) > 0 {
		*dataArray = Data{Energy: result}
	}
	return result
}

// QuickSort implementation for sorting data by region in descending order
func partition(dataArray []ResponseElement, left, right int) ([]ResponseElement, int) {
	//Comparing each request is based on the decoded region from Regions dataArrayay (variables.go)
	pivot := assets.Regions[int(dataArray[right].IdPozycja1)]
	i := left
	for j := left; j < right; j++ {
		if assets.Regions[int(dataArray[j].IdPozycja1)] <= pivot {
			dataArray[i], dataArray[j] = dataArray[j], dataArray[i]
			i++
		}
	}
	dataArray[i], dataArray[right] = dataArray[right], dataArray[i]
	return dataArray, i
}

func (dataArray *Data) SortByRegion(left, right int) []ResponseElement {
	if left < right {
		var p int
		dataArray.Energy, p = partition(dataArray.Energy, left, right)
		dataArray.Energy = dataArray.SortByRegion(left, p-1)
		dataArray.Energy = dataArray.SortByRegion(p+1, right)
	}
	return dataArray.Energy
}

// func FilterByTypeAndUnit(typeOfEnergy, energyUnit string) []ResponseElement {
// 	result := []ResponseElement{}
// 	for i := range DatadataArray {
// 		if Types[int(DatadataArray[i].IdPozycja2)] == typeOfEnergy && Units[int(DatadataArray[i].IdSposobPrezentacjiMiara)] == energyUnit {
// 			result = append(result, DatadataArray[i])
// 		}
// 	}
// 	return result
// }
