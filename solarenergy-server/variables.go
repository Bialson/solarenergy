package main

// Variables for API DBW request

const (
	DATA_CAT_1  = 1002
	DATA_CAT_2  = 232
	PERIOD      = 282
	SECTION_1   = 156
	SECTION_2   = 865
	MAX_RESULTS = 204
)

// var EnergyDataArrFiltered []ResponseElement

// Variables for decoding data from API DBW response

var Variables = map[int]string{
	156:  "Polska, województwa; Charakter miejscowości",
	232:  "Pozyskanie energii",
	237:  "Zużycie bezpośrednie energii",
	282:  "Rok - dane roczne",
	865:  "Polska, Nośniki energii odnawialnej",
	1002: "Zużycie energii elektrycznej w gospodarstwach domowych",
	1080: "Zużycie ogółem nośników energii",
}

var Units = map[int]string{
	158: "[TJ]",
	186: "[MWh]",
	187: "[kWh]",
	188: "[kWh] - na 1 mieszkańca",
	189: "[kWh] - na 1 odbiorcę",
	240: "[ktoe]",
	241: "[Mtoe]",
	242: "[GWh]",
}

var Regions = map[int]string{
	6655092: "Ogółem",
	6655093: "Miasto",
	6655153: "Wieś",
	33617:   "POLSKA",
	33619:   "MAŁOPOLSKIE",
	33929:   "ŚLĄSKIE",
	34187:   "LUBUSKIE",
	34353:   "WIELKOPOLSKIE",
	34815:   "ZACHODNIOPOMORSKIE",
	35067:   "DOLNOŚLĄSKIE",
	35390:   "OPOLSKIE",
	35542:   "KUJAWKO-POMORSKIE",
	35786:   "POMORSKIE",
	35976:   "WARMIŃSKO-MAZURSKIE",
	36185:   "ŁÓDZKIE",
	36450:   "ŚWIĘTOKRZYSKIE",
	36627:   "LUBELSKIE",
	36924:   "PODKARPACKIE",
	37185:   "PODLASKIE",
	37380:   "MAZOWIECKIE",
}

var Types = map[int]string{
	1273014: "Energia słoneczna",
	1273015: "Pompa ciepła",
	7065497: "Energia odnawialna i biopaliwa",
	7065503: "Energia wodna",
	7065509: "Energia wiatrowa",
	7065591: "Energia z ogniw fotowoltaicznych",
	7065599: "Energia geotermalna",
	7065605: "Pierwotne biopaliwa stałe",
	7065612: "Węgiel drzewny",
	7065618: "Biogaz",
	7065626: "Odnawialne odpady komunalne",
	7072409: "Biopaliwa ciekłe - bioetanol",
	7072413: "Domieszka bioetanolu w benzynie",
	7072600: "Biopaliwa ciekłe - biodiesel",
	7072606: "Domieszka biodieselu w oleju napędowym",
	7072617: "Bio nafta do silników odrzutowych",
	7072623: "Domieszka biopaliw w nafcie lotniczej",
	7072632: "Biopłyny",
	7072658: "Bioenergia",
	7072668: "Energia z kolektorów słonecznych",
}

// var Tags = map[int]string{}

//Struct for decoding data array from API DBW response

//Struct for decoding data record from API DBW response
type ResponseElement struct {
	Rownumber                int64   `json:"rownumber"`
	IdZmienna                int64   `json:"id-zmienna"`
	IdPrzekroj               int64   `json:"id-przekroj"`
	IdWymiar1                int64   `json:"id-wymiar-1"`
	IdPozycja1               int64   `json:"id-pozycja-1"`
	IdWymiar2                int64   `json:"id-wymiar-2"`
	IdPozycja2               int64   `json:"id-pozycja-2"`
	IdOkres                  int64   `json:"id-okres"`
	IdSposobPrezentacjiMiara int64   `json:"id-sposob-prezentacji-miara"`
	IdDaty                   int64   `json:"id-daty"`
	IdBrakWartosci           int64   `json:"id-brak-wartosci"`
	IdTajnosci               int64   `json:"id-tajnosci"`
	IdFlaga                  int64   `json:"id-flaga"`
	Wartosc                  float32 `json:"wartosc"`
	Precyzja                 int64   `json:"precyzja"`
}

//Arrays for storing decoded data from JSON response and filtered data
// type EnergyData []ResponseElement
