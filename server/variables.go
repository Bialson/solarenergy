package main

const (
	DATA_CAT    = 1002
	PERIOD      = 282
	SECTION     = 156
	MAX_RESULTS = 500
)

var Variables = map[int]string{
	1002:    "Energia elektryczna",
	282:     "Rok - dane roczne",
	156:     "Polska, województwa; Charakter miejscowości",
	186:     "[MWh]",
	187:     "[kWh]",
	188:     "[kWh] - na 1 mieszkańca",
	189:     "[kwh] - na 1 odbiorcę",
	6655092: "Ogółem",
	6655093: "Miasto",
	6655153: "Wieś",
	33617:   "POLSKA",
	33619:   "MAMŁOPOLSKIE",
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

type EnergyElement struct {
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

var EnergyDataArr []EnergyElement
var EnergyDataArrFiltered []EnergyElement
