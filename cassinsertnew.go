package main

import (
    "log"
	"fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
	"github.com/gocassa"
	

)

type User struct {
  FirstName  string
  SecondName  string
  ID string
}


func main() {
     fmt.Println("Starting http server")
      http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))

    })
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        storeUser(r)

    })


    // Continue to process new requests until an error occurs
    log.Fatal(http.ListenAndServe(":8003", nil))

}

func storeUser(r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
      panic(err) // dont panic - it will stop the server
  }
  log.Println(string(body)) //1. got our json data upon curl
  var user User  //you can also do this - user := User
  err = json.Unmarshal(body, &user) //giving the ref to the object so that unmarshall can update the data from body -> user object.
  if err != nil {
      panic(err)
  }
  //Now after unmarshalling, you can access all the data from User object
 	 log.Println(user.FirstName)
 	 log.Println(user.SecondName)
	log.Println(user.ID)
  //now that we have our user object wiuth data, we can use cassandra driver to pass our object which will store the data
	 // connect to the cluster
	
        keySpace, err := gocassa.ConnectToKeySpace("demo", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}
	PatientsTable := keySpace.Table("Patient", &User{}, gocassa.Keys{
		PartitionKeys:[]string{"ID"},
	})
	// Create the table - we ignore error intentionally
	PatientsTable.Create()

	// We insert the first record into our table - yay!
	err = PatientsTable.Set(User{
		FirstName : user.FirstName,
		SecondName :user.SecondName,
		ID : user.ID,
	}).Run()
	if err != nil {
		panic(err)
	}

	result := User{}
	if err := PatientsTable.Where(gocassa.Eq("ID", user.ID)).ReadOne(&result).Run(); err != nil {
		panic(err)
	}
	fmt.Println(result)

}


