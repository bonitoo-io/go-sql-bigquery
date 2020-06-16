#  BigQuery SQL Driver for Golang
This is an implementation of the BigQuery Client as a `database/sql/driver` for easy integration and usage.

# Goals of project

This project is meant to be a basic `database/sql` driver implementation for Golang so that developers can easily use 
`*sql.DB` functions with Google's BigQuery database.

Unlike the original project, this driver
- does __not__ contain ORM extension
- provides explicit authentication support in connection string
- contains no logging

# Authentication

As this is using the Google Cloud Go SDK, you will need to have your credentials available
via the GOOGLE_APPLICATION_CREDENTIALS environment variable point to your credential JSON file.

Alternatively, you can specify `apiKey` connection string parameter with API key value,
or `credentials` parameter with base-64 encoded service account or refresh token JSON credentials as the value.  
Connection string examples:  
```js
"bigquery://projectid/location/dataset?apiKey=AIzaSyB6XK8IO5AzKZXoioQOVNTFYzbDBjY5hy4"
"bigquery://projectid/location/dataset?credentials=eyJ0eXBlIjoiYXV0..."
```

## Usage

```go
package main

import "github.com/bonitoo-io/go-sql-bigquery"

func main() {
    db, err := sql.Open("bigquery", "bigquery://projectid/location/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    ...
}
```

# Contribution

Contributions are welcome.  

# Current Support

* [x] `driver.Conn` implemented
* [x] `driver.Querier` implemented
* [x] `driver.Pinger` implemented
* [x] `driver.DriverContext` implemented
* [x] `driver.QueryerContext` implemented
* [x] `driver.ExecerContext` implemented
* [x] Prepared Statements - supported via a quick hack
* [ ] Parameterized Queries
