package main

import(
	"fmt"
	r "gopkg.in/gorethink/gorethink.v4"

)

type User struct {
	Id string `gorethink:"id,omitempty"`
	Name string `gorethink:""`
}


func main(){
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "rtsupport",
		})
	if err != nil{
		fmt.Println(err)
		return
	}
	// user := User{
	// 	Name: "Keith Lashley",
	// }

	// response, _ := r.Table("user").
	// 	Get("b50a3e62-7f6a-4c59-8f2d-0ae1448245d0").
	// 	Delete().
	// 	RunWrite(session)

	// fmt.Printf("%#v\n", response)

	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse){
		fmt.Printf("%#v\n", changeResponse)
	}
}