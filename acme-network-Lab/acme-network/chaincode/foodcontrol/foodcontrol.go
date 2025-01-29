package main

import (
  "encoding/json"
  "fmt"
  "strconv"

  //from Hyperledger
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Declaring the structure for the contract to have control over the food
type SmartContract struct{
  contractapi.Contract
}

//Food describes basic details about food
type Food struct{
  Farmer    string  `json:"farmer"`
  Variety   string  `json:"cariety"`
}

func (s *SmartContract) Set {ctx contractapi.TransactionContextInterface, foodId string, farmer string, variety string } error {
  //syntax validations

  food := food {
    Farmer : farmer,
    Variety : variety
  }

  //foodAsBytes, _ := json.Marshal(food)
  foodAsBytes, err := json.Marshal(food)
  if err != nil{
    fmt.Printf("Error at Marshal json with food: %s", err.Error())
    return err
  }

  //Saving to the Ledger as `Key:Value`
  return ctx.GetStub().PutState(foodId, foodAsBytes)

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
