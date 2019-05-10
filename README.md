# objectia-go
[![Build Status](https://travis-ci.org/objectia/objectia-go.svg?branch=master)](https://travis-ci.org/objectia/objectia-go) 
<!-- [![codecov](https://codecov.io/gh/objectia/objectia-go/branch/master/graph/badge.svg)](https://codecov.io/gh/objectia/objectia-go) -->

Go client for [Objectia API](https://objectia.com)&reg; 
 
## Documentation

See the [Go API docs](https://docs.objectia.com/guide/go.html).

## Installation

You don't need this source code unless you want to modify the package. If you just
want to use the package, just run:

```bash
$ go get -u github.com/objectia/objectia-go
```    

### Requirements

* Go 1.8 or later


### Development:

```bash
$ go get -u github.com/stretchr/testify/assert
```

## Usage

The library needs to be configured with your account's API key. Get your own API key by signing up for a free [Objectia account](https://objectia.com).

```go
package main

import (
    "fmt"
    "github.com/objectia/objectia-go"
)

func main() {
    apiKey := "<your API key>"
    client, err := objectia.NewClient(apiKey, nil)
    if err != nil {
        panic(err)
    }

    result, err := client.GeoLocation.Get("8.8.8.8", nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Location: %+v\n", result)
}
```

Look in the [examples folder](./examples) for more code examples.


## API

#### type Client

``` go
type GeoLocation struct {
  // contains filtered or unexported fields
}

type Client struct {
  // contains filtered or unexported fields
  GeoLocation *GeoLocation
}
```

Client is the API client that performs all operations against a Objectia API server.


#### func NewClient 

``` go
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) 
```

NewClient initializes a new API client. You may use your own http client, or pass nil to use the default configuration.

## GeoLocation API

#### func (*GeoLocation) Get

``` go
func (c *GeoLocation) Get(ip string, options *GetLocationOptions) (*IPLocation, error)
```

Returns geoip location data for the specified IP address (or domain name).


#### func (*GeoLocation) GetCurrent

``` go
func (c *GeoLocation) GetCurrent(options *GetLocationOptions) (*IPLocation, error)
```

Returns geoip location data for the requester's IP address.


#### func (*GeoLocation) GetBulk

``` go
func (c *GeoLocation) GetBulk(iplist []string, options *GetLocationOptions) ([]IPLocation, error)
```

Returns an array of geo location data for the specified IP addresses (or domain names).
You can add up to 50 IP addresses or domain names.


## License and Trademarks

Copyright (c) 2018-19 UAB Salesfly.

Licensed under the [MIT license](https://en.wikipedia.org/wiki/MIT_License). 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Objectia is a registered trademark of [UAB Salesfly](https://www.salesfly.com).