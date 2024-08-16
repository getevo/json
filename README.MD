# Advanced JSON

JSON Advanced: A Fork of `encoding/json` with Enhanced Functionality

**JSON Advanced** is a fork of the standard Go `encoding/json` package, providing additional functionality to marshal and unmarshal JSON. This library is fully compatible with the original `encoding/json` library and includes two new struct tags: `omit_encode` and `omit_decode`.

## Features

- Fully compatible with Go's standard `encoding/json` library.
- Supports two new struct tags:
    - `omit_encode`: Excludes a field from being encoded into JSON.
    - `omit_decode`: Excludes a field from being decoded from JSON.

## Installation

To install the package, run:

```sh
go get github.com/getevo/json
```

Then, import it in your Go files:
```golang
import "github.com/yourusername/json-advanced"
```

## Usage
The JSON Advanced package can be used in the same way as the standard `encoding/json` package, with the addition of the omit_encode and omit_decode struct json tags.

### Example
Here's a basic example that demonstrates the use of the new struct tags:
```golang 
package main

import (
    "fmt"
    "log"
    "github.com/getevo/json"
)

type User struct {
    Name        string `json:"name"`
    Password    string `json:"password,omit_encode"`
    SecretToken string `json:"secret_token,omit_decode"`
}

func main() {
    user := User{
        Name:        "John Doe",
        Password:    "password123",
        SecretToken: "supersecrettoken",
    }

    // Marshal the struct to JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData)) // Output: {"name":"John Doe","secret_token":"supersecrettoken"}

    // Unmarshal JSON back into the struct
    input := `{"name":"Jane Doe","password":"newpassword","secret_token":"newtoken"}`
    err = json.Unmarshal([]byte(input), &user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", user)
    // Output: {Name:Jane Doe Password:password123 SecretToken:supersecrettoken}
}

```


## API Reference

The API of **JSON Advanced** is fully compatible with the standard `encoding/json` package, with added support for the `omit_encode` and `omit_decode` struct tags.

### Functions

- **`func Marshal(v interface{}) ([]byte, error)`**

  Marshals the provided data into JSON format. Fields tagged with `omit_encode` will be excluded from the JSON output.

- **`func Unmarshal(data []byte, v interface{}) error`**

  Unmarshals the provided JSON data into the specified struct. Fields tagged with `omit_decode` will not be populated during this process.

- **`func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)`**

  Similar to `Marshal`, but formats the output with indentation. Fields tagged with `omit_encode` will be omitted from the JSON output.

- **`func NewDecoder(r io.Reader) *json.Decoder`**

  Creates a new JSON decoder. Fields tagged with `omit_decode` will be ignored when decoding JSON into structs.

- **`func NewEncoder(w io.Writer) *json.Encoder`**

  Creates a new JSON encoder. Fields tagged with `omit_encode` will be omitted when encoding structs to JSON.

### Types

- **`type Decoder`**

  The `Decoder` type functions similarly to `encoding/json.Decoder`, with the additional capability to respect the `omit_decode` tag when decoding JSON.

- **`type Encoder`**

  The `Encoder` type functions similarly to `encoding/json.Encoder`, with the additional capability to respect the `omit_encode` tag when encoding to JSON.

## License

This project is licensed under the MIT License. Please see the [LICENSE](LICENSE) file for more details.


## Acknowledgments

This library is a fork of the Go standard library's `encoding/json` package, enhanced with additional functionality inspired by the needs of the Go development community.
