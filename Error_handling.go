package main

// Description of error handling

// Errors are the unintended result of a program. Suppose we want to have an API call to an external service.
//  This API call can be successful or unsuccessful. Application errors in Go language 
//  can be detected in cases where the type of error is present. Consider the following example:

resp, err: = http.Get ("http://example.com/")

// In the above code, calling the API to the Error object can be successful or unsuccessful. 
// We can check if there is an error or not, and manage the response accordingly:

import (
  "fmt"
  "net / http"
)

func main () {
  resp, err: = http.Get ("http://example.com/")
  if err! = nil {
    fmt.Println (err)
    return
  }
  fmt.Println (resp)
}


// Panic is something that is not managed and occurs suddenly during program execution.
//  Go language is not an ideal way to manage exceptions in a program. It is recommended to use an error object
//   instead. When a panic occurs, the execution of the program is suspended and then defer is executed.

// Defer is what is executed at the end of each function.
// There are several channels that a function is waiting for. to this