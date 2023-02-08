package main

//Method for filtering data by region

func FilterByRegion(region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

//Method for filtering data by character

func FilterByCharacter(character string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

//Method for filtering data by region and character

func FilterByCharacterAndRegion(character, region string) []ResponseElement {
	result := []ResponseElement{}
	for i := range EnergyDataArr {
		if Regions[int(EnergyDataArr[i].IdPozycja2)] == character && Regions[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

//QuickSort implementation for sorting data by region in descending order

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
