package main

func FilterByRegion(region string) []EnergyElement {
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByCharacter(character string) []EnergyElement {
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja2)] == character {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func FilterByCharacterAndRegion(character, region string) []EnergyElement {
	result := []EnergyElement{}
	for i := range EnergyDataArr {
		if Variables[int(EnergyDataArr[i].IdPozycja2)] == character && Variables[int(EnergyDataArr[i].IdPozycja1)] == region {
			result = append(result, EnergyDataArr[i])
		}
	}
	return result
}

func partition(arr []EnergyElement, left, right int) ([]EnergyElement, int) {
	pivot := Variables[int(arr[right].IdPozycja1)]
	i := left
	for j := left; j < right; j++ {
		if Variables[int(arr[j].IdPozycja1)] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func QuickSortByRegion(arr []EnergyElement, left, right int) []EnergyElement {
	if left < right {
		var p int
		arr, p = partition(EnergyDataArr, left, right)
		arr = QuickSortByRegion(arr, left, p-1)
		arr = QuickSortByRegion(arr, p+1, right)
	}
	return arr
}