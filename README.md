# Event Decoder Library

## Overview
The `event-decoder` library is a Go package designed to facilitate the decoding of Ethereum blockchain events. Utilizing the power of the go-ethereum ABI package, this library provides an easy-to-use interface for extracting and interpreting event data from Ethereum smart contracts.

## Features
- Decoding of Ethereum event parameters.
- Support for various Ethereum data types.
- Integration with the go-ethereum ABI package for robust handling of Ethereum data structures.
- Error handling specifically tailored for event decoding.

## Installation
To use the `event-decoder` library in your Go project, you can install it using `go get`:

```bash
go get github.com/paribu/event-decoder
```

## Supported Types

|               | Single | Array[] | Matrix[]-[] | Fixed Array[n] | Fixed Inner Matrix[n]-[] | Fixed Outer Matrix[]-[n] | Fixed Matrix[n]-[n] |
| ------------- |:------:|:-------:|:-----------:|:--------------:|:------------------------:|:------------------------:|:-------------------:|
| bool          |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| string        |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| bigint        |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| address       |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| uint8         |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| uint          |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| int           |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| bytesN        |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| bytes         |   ✅   |    ✅   |      ✅     |       ✅       |            ✅            |            ✅            |          ✅         |
| SimpleStruct  |   ✅   |    ❌   |      ❌     |       ❌       |            ❌            |            ❌            |          ❌         |
| NestedStruct  |   ✅   |    ❌   |      ❌     |       ❌       |            ❌            |            ❌            |          ❌         |
| ComplexStruct |   ❌   |    ❌   |      ❌     |       ❌       |            ❌            |            ❌            |          ❌         |

We are planning to add missing types as soon as possible, but contributions are welcome.

## Usage
The primary function of this library is Decode, which takes an event.Event and a contractABI *abi.ABI as input and returns a slice of *event.Parameter and an error.

### Basic Example
Here is a basic example of how to use the event-decoder library:

```go
package main

import (
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/paribu/event-decoder"
    "github.com/paribu/event-decoder/event"
)

func main() {
    // Assuming 'e' is your event data and 'contractABI' is the ABI of the contract
    e := &event.Event{ /*...*/ }
    contractABI := &abi.ABI{ /*...*/ }

    parameters, err := event_decoder.Decode(e, contractABI)
    if err != nil {
        // Handle error
    }

    // Use the decoded parameters
    for _, p := range parameters {
        fmt.Println("Parameter:", p.Name, "Value:", p.Value, "Type:", p.Type)
    }
}
```

## Dependencies
- [go-ethereum](https://github.com/ethereum/go-ethereum): This library relies on the go-ethereum package, specifically its ABI capabilities.

## Contributing
Contributions to the `event-decoder` library are welcome. Please feel free to submit issues and pull requests through the GitHub repository.

## License
This library is licensed under MIT License. Please see the LICENSE file for more details.
