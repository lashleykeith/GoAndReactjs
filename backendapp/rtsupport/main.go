package main

import "net/http"
import "log"
import r "gopkg.in/gorethink/gorethink.v4"

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink: "name"`
}

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}



//http://jsbin.com/nayuseqemi/edit?js,console

/*
func addChannel(data interface{}) (error){
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil{
		return err
	}
	channel.Id = "1"
	//fmt.Printf("%#v\n", channel)
	fmt.Println("added channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn){
	// TODO: rethinkDB Query / changefeed
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add",
			Channel{"1", "Software Support"}}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}


*/