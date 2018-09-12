[![Build Status](https://travis-ci.org/mchirico/gog.svg?branch=develop)](https://travis-ci.org/mchirico/gog) [![Maintainability](https://api.codeclimate.com/v1/badges/1558bc5ede187bd55266/maintainability)](https://codeclimate.com/github/mchirico/gog/maintainability)

# gog
Go Gorilla Mux -- My Template


## Install

```bash
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/rpc
go get -u github.com/levigross/grequests

go get -u github.com/mchirico/gog

```

## Upload File

```go
package main

import (
	"github.com/levigross/grequests"
	"log"
)

func main() {

	fd, err := grequests.FileUploadFromDisk("./books.json")
	resp, err := grequests.Post("http://localhost:8080/upload",
		&grequests.RequestOptions{
			Files: fd,
			Data:  map[string]string{"One": "Two"},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	log.Println(resp.String())
}


```




## Example

Here's an example of a very simple RPC server. You'll need
to download [books.json](https://raw.githubusercontent.com/mchirico/mchirico.github.io/master/p/books.json)

[Reference: Building RESTFul Web Services with Go](http://my.safaribooksonline.com/book/web-development/9781788294287)

## Server

```go
package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Args holds arguments passed to JSON RPC service
type Args struct {
	Id string
}

// Book struct holds Book JSON structure
type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}
type JSONServer struct{}

// GiveBookDetail
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON file and load data
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each book to find the given book
	for _, book := range books {
		if book.Id == args.Id {
			// If book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}
func main() {
	// Create a new RPC server
	s := rpc.NewServer() // Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}


```

## Client

```bash
curl -X POST \
   http://localhost:1234/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "Id": "1234"
   }],
   "id": "1"
}'
```


