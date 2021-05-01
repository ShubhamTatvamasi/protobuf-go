package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShubhamTatvamasi/protobuf-go/pkg/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {

	sm := doSimple()

	// readAndWriteDemo(sm)
	jsonDemo(sm)

}

func jsonDemo(sm proto.Message) {
	smAsJson := toJSON(sm)

	fmt.Println(smAsJson)

	sm2 := &simple.SimpleMessage{}
	fromJSON(smAsJson, sm2)
	fmt.Println("Created proto buf", sm2)

}

func toJSON(pb proto.Message) string {

	marshaler := jsonpb.Marshaler{}
	out, _ := marshaler.MarshalToString(pb)
	return out
}

func fromJSON(in string, pb proto.Message) {

	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatal(nil)
	}

}

func readAndWriteDemo(sm proto.Message) {

	writeToFile("simple.bin", sm)

	sm2 := &simple.SimpleMessage{}

	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, _ := proto.Marshal(pb)

	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file")
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, _ := os.ReadFile(fname)

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol struct", err)
		return err
	}

	return nil
}

func doSimple() *simple.SimpleMessage {

	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "New name"

	fmt.Println("The Name is:", sm.GetName())
	fmt.Println("The Id is:", sm.GetId())

	return &sm

}
