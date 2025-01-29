package main

import (
  "encoding/json"
  "fmt"
  "strconv"

  //from Hyperledger
  "github.com/hyperledger/fabric-contract-api-go"
)

//Declaring the structure for the contract to have control over the food
type SmartContract struct{
  contractapi.Contract
}

//Main Function 
func main(){
  chaincode, err := contractapi.NewChanincode(new(SmartContract))

  //If an error occurs while creating the contrat
  if err != nil{
    fmt.Printf("Error create foodcontrol smartcontract: %s", err.Error())
    return
  }

  //If an error occurs while Starting the contract
  if err := chaincode.Start(); err != nil{
    fmt.Printf("Error create foodcontrol smartcontract: %s", err.Error())
  }

}
