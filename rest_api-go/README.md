# REST API in Go Without And Frameworks

I have used the standard library `net/http` in _Go (Golang)_ which is used for web development, to develop a basic library amanagement system.

In this I'll be creating basic REST APIs using purely Go, without any frameworks, and have used the in-memory storage to store the data.


## API

Following basic APIs will be build to manage a library (list) of books.

- `GET /books` : Get all books summary
- `GET /books/{id}` : Get details of a book by id
- `POST /books` : Add a new book
- `POST /reserve/{id}` : Reserve a book by it's id
- `POST /return/{id}` : Return a book by it's id


## API calls

```bash
# Get summary of available books
curl http://127.0.0.1:8080/books

# Get detail of a book
curl http://127.0.0.1:8080/books/1

# Add a book
curl -X POST http://127.0.0.1:8080/books -H "Content-Type: application/json" -d '{"title": "How to Train your CAT", "author": "Garfield the Orange"}'

curl -X POST http://127.0.0.1:8080/books -H "Content-Type: application/json" -d '{"title": "Strategic Approach to finding Waldo", "author": "Kenny C. Looper"}'

# Reserve a book
curl -X POST http://127.0.0.1:8080/reserve/1

# Return a book
curl -X POST http://127.0.0.1:8080/return/1
```


## Development Notes

1. json.marshal/json.unmarshal
    - In Go for correctly encode (marshal) or decode (unmarshal) data from a struct field, the field must be exported.
    - An exported field in Go is the name which begins with an uppercase letter.
    - If a field has tag but not exported (starts with a lowercase letter) will be ignored when encoding.

```go
type MyStruct struct {
	// This field will be ignored by json.Marshal/Unmarshal
	unexportedField string `json:"unexported_field"`
	// This field will be correctly processed
	ExportedField string `json:"exported_field"`
}
```