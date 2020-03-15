package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var errorColor = "\033[1;31m%s\033[0m\n"

func main() {
	inputFile := flag.String("input", "", "The CDS package file")
	outputFile := flag.String("output", "", "The tarball file to create")
	flag.Parse()
	if *inputFile == "" || *outputFile == "" {
		fmt.Println("CDSTool Command Options:")
		flag.PrintDefaults()
		return
	}

	chaincode, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf(errorColor, err)
		return
	}

	depSpec := &pb.ChaincodeDeploymentSpec{}
	err = proto.Unmarshal(chaincode, depSpec)
	if err != nil {
		fmt.Printf(errorColor, err)
		return
	}
	fmt.Printf("Chaincode Information: =%+v\n", depSpec.ChaincodeSpec)

	payload := depSpec.CodePackage
	err = ioutil.WriteFile(*outputFile, payload, 0644)
	if err != nil {
		fmt.Printf(errorColor, err)
		return
	}
}
