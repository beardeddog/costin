package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"costin/config"
	"costin/stamps"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("collecting").C("stamps")
	// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := resturant{}
	// err = c.Find(bson.M{"borough": "Manhattan"}).One(&result)

	configs := []config.Config{}
	err = c.Find(bson.M{"type": "config"}).All(&configs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("configs: %v\n\n", configs)

	stamps := []stamps.StampData{}
	err = c.Find(bson.M{"type": "stamp"}).All(&stamps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stamps: %v\n", stamps)

}
