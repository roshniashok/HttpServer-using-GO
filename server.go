package main

import (
    "log"
	"fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"

)

type User struct {
  FirstName   string
  SecondName  string
  Age         int
  Records     *Records

}

type Records struct {
  PatientId   int
  Condition   bool
  HeartData   *HeartData
  LungData    *LungData //what ever categories
}

type HeartData struct {
  Beats   int
  Condition bool
}

type LungData struct {
  tarcontent int
  condition bool
}


func main() {
     fmt.Println("Starting http server")
    // All URLs will be handled by this function
    // http.HandleFunc uses the DefaultServeMux
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))

    })
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        storeUser(r)

    })


    // Continue to process new requests until an error occurs
    log.Fatal(http.ListenAndServe(":8000", nil))

}

func storeUser(r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
      panic(err) // dont panic - it will stop the server
  }
  log.Println(string(body)) //1. got our json data upon curl
  var user User //you can also do this - user := User{}
  err = json.Unmarshal(body, &user) //giving the ref to the object so that unmarshall can update the data from body -> user object.
  if err != nil {
      panic(err)
  }
  //Now after unmarshalling, you can access all the data from User object
  log.Println(user.FirstName)
  log.Println(user.SecondName)
  log.Println(user.Records.HeartData.Condition) //its nested. You can understand or derive a json from struct or vice versa.

  //now that we have our user object wiuth data, we can use cassandra driver to pass our object which will store the data
}
