package main

import (
  "encoding/json"
  "fmt"

  // Hyperledger Fabric contract API
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing food assets
type SmartContract struct {
  contractapi.Contract
}

// Food describes basic details about a food asset
type Food struct {
  Farmer  string `json:"farmer"`
  Variety string `json:"variety"`
}

// Set creates a new food asset in the ledger
func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, foodId string, farmer string, variety string) error {
  // Check if the food already exists
  existingFood, err := s.Query(ctx, foodId)
  if existingFood != nil {
    return fmt.Errorf("Food ID %s already exists", foodId)
  }

  // Create new food asset
  food := Food{
    Farmer:  farmer,
    Variety: variety,
  }

  // Convert struct to JSON
  foodAsBytes, err := json.Marshal(food)
  if err != nil {
    return fmt.Errorf("Error marshaling JSON: %s", err.Error())
  }

  // Save to ledger
  return ctx.GetStub().PutState(foodId, foodAsBytes)
}

// Query retrieves a food asset from the ledger by its ID
func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, foodId string) (*Food, error) {
  foodAsBytes, err := ctx.GetStub().GetState(foodId)
  if err != nil {
    return nil, fmt.Errorf("Error reading from world state: %s", err.Error())
  }

  if foodAsBytes == nil {
    return nil, fmt.Errorf("Food ID %s does not exist!", foodId)
  }

  var food Food
  err = json.Unmarshal(foodAsBytes, &food)
  if err != nil {
    return nil, fmt.Errorf("Error unmarshaling JSON: %s", err.Error())
  }

  return &food, nil
}

// Main function
func main() {
  chaincode, err := contractapi.NewChaincode(new(SmartContract))
  if err != nil {
    fmt.Printf("Error creating foodcontrol smart contract: %s", err.Error())
    return
  }

  if err := chaincode.Start(); err != nil {
    fmt.Printf("Error starting foodcontrol smart contract: %s", err.Error())
  }
}

