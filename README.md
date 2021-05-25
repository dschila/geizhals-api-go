# Geizhals API Go
Using geizhals.de via a RESTFul interface to search and display products. 

## Motivation
As part of a private project to learn the programming language Go, I needed the possibility to collect availability of articles from Geizhals.de without calling the website directly. 

## Run
`# go run main.go`

or

`# go build -o geizhals-api-go`

`# ./geizhals-api-go`
## Docker
`# docker build -t geizhals-api-go:1.0 .`

`# docker run -d --name geizhals-api-go -p 8080:8080 geizhals-api-go:1.0`

## How to use

Open your Browser and access the API over `localhost:8080/api/search/AMD`

Available Endpoints:

```http
GET /api/search/:query
GET /api/search/category/:category/:query
GET /api/article/:identifier
```

## Contribute
Maybe someone will find the project helpful or can give me hints/improvements. 