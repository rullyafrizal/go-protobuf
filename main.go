package main

import (
	"log"

	"github.com/rullyafrizal/go-protobuf/user"

	"google.golang.org/protobuf/proto"
)

func main() {
	var u *user.Person = &user.Person{
		Name: "Rully Afrizal",
		Age:  20,
		Address: &user.Address{
			StreetName:   "Jl. Raya",
			StreetNumber: 103,
			City:         "Bandung",
			State:        "Jawa Barat",
			Country:      "Indonesia",
			PostCode:     "40132",
		},
	}

	// Encode (from user to protobuf)
	data, err := proto.Marshal(u)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Println(data)

	// Decode (from protobuf to user)
	var decodedUser *user.Person = &user.Person{}

	err = proto.Unmarshal(data, decodedUser)

	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Println(decodedUser.GetName())
	log.Println(decodedUser.GetAge())
	log.Println(decodedUser.GetAddress().GetStreetName())
	log.Println(decodedUser.GetAddress().GetStreetNumber())
	log.Println(decodedUser.GetAddress().GetPostCode())
	log.Println(decodedUser.GetAddress().GetCity())
	log.Println(decodedUser.GetAddress().GetState())
	log.Println(decodedUser.GetAddress().GetCountry())
}
