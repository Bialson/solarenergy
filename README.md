
# solarenergy

This tool provides interaction with energy consumption data from Poland's Central Statistical Office (GUS) provided by API DBW, created also by GUS. 




## Features

- Decoding data returned from API DBW
- Sorting results by region name
- Filtering results with region name and character
- Displaying results in custom CLI




## Roadmap

- More filtering parameters

- Analyzing module for percentage power consumption from regions acording to total Poland's power consumption

- Present data on charts



## Tech Stack

**Client CLI:** `Go + urfave cli v2`

**Server:** `Go + gRPC`

<br>

## Installation

Make sure you have installed Go. Next configure Go paths in your profile files.

### Linux `~/.profile`
```
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
  export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin
```

Verify path by following command (if path contains above paths everything is fine)
```bash
  $GOPATH
```

### Windows 

Check your enviromental variable
```powershell
  $env:GOPATH
```

If everything is ok you should get the following path in you console
```powershell
  C:\Users\[username]\go
```

Add the workspace’s `bin` subdirectory to your `$PATH`. You can do this using the `setx` command in PowerShell

```powershell
  setx PATH "$($env:path);$GOPATH\bin"
```

### Docker 

Change directory to main folder `solarenergy` and then build image
```bash
  docker build --pull --rm -f "Dockerfile" -t solarenergy:latest "." 
```
<br>

### CLI's installation

Clone repository to your local folder

```bash
  git clone https://github.com/Bialson/solarenergy.git
```

Your folder structure will be looking like this

```
.
└── solarenergy/
    ├── proto/
    │   ├── energy_grpc.pb.go
    │   ├── energy.pb.go
    │   └── energy.proto
    ├── solarenergy-cli/
    │   ├── bin/
    │   │   └── solarenergy-cli
    │   ├── main.go
    │   └── service.go
    ├── solarenergy-server/
    │   ├── bin/
    │   │   └── solarenergy-server
    │   ├── main.go
    │   ├── service_test.go
    │   ├── service.go
    │   ├── sort.go
    │   ├── test_server.go
    │   └── variables.go
    ├── .gitignore
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── Makefile
    ├── README.md
    └── run.sh
```

To install client and server CLI change directory to `solarenergy-cli` and `solarenergy-server` folders and run `go install` command in each folder

```bash
  cd solarenergy-cli
  go install
  cd ../solarenergy-server
  go install
```

> **Note**
> Package provides builded files ready to install, if you want to add some changes and then rebuild the preffered CLI, run `go build -o bin/[file_name]` command in it's folder.

<br>
    
## Run Locally

#### Start server
```bash
  solarenergy-server
```

On your you will be receiving log's from server, request, response status.

#### Run client

```bash
  solarenergy-cli gp -y YEAR_VALUE -r REGION_NAME -ch REGION_CHARACTER -a AMOUNT_VALUE
```

If you run `solarenergy-cli` command without params, you should see help page with all params supported by this command.

#### Run tests
```bash
  cd solarenergy-server
  go test
```

#### Run docker image
```bash
  docker run --rm -it  solarenergy:latest 
```
<br>

## Client CLI Reference

> **Note**
> If you don't provide any optional params, data request will be executed with default values

| Parameter | Type     | Value            |
| :-------- | :------- | :------------------------- |
| `Year` | `int` | 2020 |
| `Response amount` | `int` | 0 |
| `Region` | `string` | "" |
| `Region character` | `string` | "" |

<br>

#### Get all data
```bash
  solarenergy-cli gp -y YEAR_VALUE
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `get-power, gp` | `string` | **Required**. Get power consumption command |
| `--year, -y` | `int` | **Optional**. Get power consumption from given year |

<br>

#### Get data from region

```bash
  solarenergy-cli gp -y YEAR_VALUE -r REGION_NAME
```

| Parameter | Type     | Description                       | 
| :-------- | :------- | :-------------------------------- |
| `--year, -y` | `int` | **Optional**. Get power consumption from given year |
|  `--region, -r`  | `string` | **Optional**.  Precise power consumption by region |

<br>

#### Get data by region character

```bash
  solarenergy-cli gp -y YEAR_VALUE -r REGION_NAME -ch REGION_CHARACTER
```

| Parameter | Type     | Description                       | 
| :-------- | :------- | :-------------------------------- |
| `--year, -y` | `int` | **Optional**. Get power consumption from given year |
| `--region, -r`      | `string` | **Optional**.  Precise power consumption by region |
| `--character, -ch`      | `string` | **Optional**.  Precise power consumption by region character |

<br>

#### Get data and precise their amount

```bash
  solarenergy-cli gp -y YEAR_VALUE -r REGION_NAME -ch REGION_CHARACTER -a AMOUNT_VALUE
```  

| Parameter | Type     | Description                       | 
| :-------- | :------- | :-------------------------------- |
| `--year, -y` | `string` | **Optional**. Get power consumption from given year |
| `--region, -r`      | `string` | **Optional**.  Precise power consumption by region |
| `--character, -ch`      | `string` | **Optional**.  Precise power consumption by region character |
| `--amount, -a`      | `int` | **Optional**.  Limit power consumption records |

<br>

## Params reference

| Parameter | Type     | Value            |
| :-------- | :------- | :------------------------- |
| `Year` | `int` | <2000; 2020> |
| `Response amount` | `int` | <0; 204> |
| `Region` | `string` | 	POLSKA <br> MAŁOPOLSKIE <br>  ŚLĄSKIE <br> LUBUSKIE <br> WIELKOPOLSKIE <br> ZACHODNIOPOMORSKIE <br> DOLNOŚLĄSKIE <br> OPOLSKIE <br> KUJAWSKO-POMORSKIE <br> POMORSKIE <br> WARMIŃSKO-MAZURSKIE <br> ŁÓDZKIE <br> ŚWIĘTOKRZYSKIE <br> LUBELSKIE <br> PODKARPACKIE <br> PODLASKIE <br> MAZOWIECKIE|  
| `Region character` | `string` | Ogółem <br> Miasto <br> Wieś |
