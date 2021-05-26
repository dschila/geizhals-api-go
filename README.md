# Geizhals API Go
Using geizhals.de via a RESTFul interface to search and display products. 

## Motivation

As part of a private project to learn the programming language Go, I needed the possibility to collect availability of articles from Geizhals.de without calling the website directly. 

## Run

    # go run main.go
or

    # go build -o geizhals-api-go
    # ./geizhals-api-go   

## Docker

    # docker build -t geizhals-api-go:1.0 .
    # docker run -d --name geizhals-api-go -p 8080:8080 geizhals-api-go:1.0

# How to use

## Search article

### Request

`GET /api/search/:query`

    curl -i -H 'Accept: application/json' http://localhost:8080/api/search/AMD

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Wed, 26 May 2021 07:39:06 GMT
    Transfer-Encoding: chunked

    [
        {
            "Name": "AMD Ryzen 5 3600, 6C/12T, 3.60-4.20GHz, boxed",
            "LowestPrice": 187,
            "OfferCount": 70,
            "URL": "amd-ryzen-5-3600-100-100000031box-a2064574.html",
            "ImageURL": "https://gzhls.at/i/45/74/2064574-r0.jpg",
            "Availability": "lagernd beim Händler"
        },
    ...
    ]
    
## Search article in a specific category

### Request

`GET /api/search/:query?category=:categoryId`

    curl -i -H 'Accept: application/json' http://localhost:8080/api/search/LG?category=4

Category | ID
--- | ---
Hardware | 1
Software | 2
Video, Photo & TV | 4
Phone | 5
Audio & HiFi | 6
Movies | 7
Household | 8
Sport & Leisure | 9
Drugstore | 10
Car & Motorcycle | 11
Toys & Model Making | 12
Hardware store & Garden | 13
Office & School | 14

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Wed, 26 May 2021 07:39:06 GMT
    Transfer-Encoding: chunked

    [
        {
            "Name": "LG OLED 65CX9LA",
            "LowestPrice": 1675,
            "OfferCount": 24,
            "URL": "lg-oled-65cx9la-a2250846.html",
            "ImageURL": "https://gzhls.at/i/08/46/2250846-r0.jpg",
            "Availability": "lagernd beim Händler"
        },
    ...
    ]
## Get article details

### Request

`GET /api/article/:identifier`

    curl -i -H 'Accept: application/json' http://localhost:8080/api/article/amd-ryzen-5-3600-100-100000031box-a2064574.html

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Wed, 26 May 2021 07:39:06 GMT
    Transfer-Encoding: chunked

    {
        "article": {
            "Name": "AMD Ryzen 5 3600, 6C/12T, 3.60-4.20GHz, boxed (100-100000031BOX)",
            "PriceFrom": 187",
            "PriceTo": 388.6,
            "ImageURL": "https://gzhls.at/i/45/74/2064574-m0.jpg"
        },
        "provider": [
            {
                "Name": "Mindfactory",
                "OfferURL": "https://geizhals.de//redir.cgi?h=mindfactory\u0026loc=https%3A%2F%2Fwww.mindfactory.de%2Fproduct_info.php%2Finfo%2Fp1313643%2Fpid%2Fgeizhals\u0026ghaID=2064574\u0026key=7167ae0337fc34de7a73d669596f3ab8",
                "Price": 187,
                "Availability": "Zentrallager: 5 Stück lagernd, Lieferung 1-3 WerktageFiliale Wilhelmshaven: 5 Stück lagerndStand: 26.05.2021, 09:55",
                "Shipping": "Vorkasse, PayPal, Giropay, Finanzierung € 8,99.Expressversand möglich.Lieferung nur innerhalb Deutschlands."
            },
        ...
        ]
    }

# Contribute
Maybe someone will find the project helpful or can give me hints/improvements. 