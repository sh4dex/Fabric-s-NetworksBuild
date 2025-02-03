#!/bin/bash

# Set the name of the channel
export CHANNEL_NAME=marketplace

# Set the chaincode name
export CHAINCODE_NAME=foodcontrol

# Set the chaincode version
export CHAINCODE_VERSION=1

# Define the chaincode runtime language (Go in this case)
export CC_RUNTIME_LANGUAGE=golang

# Define the path to the chaincode source code
export CC_SOURCE_PATH="/opt/gopath/src/github.com/chaincode/$CHAINCODE_NAME/"

# Set the path to the Orderer's TLS certificate
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/acme.com/msp/tlscacerts/tlsca.acme.com-cert.pem

# Print environment variables for verification
echo "Environment variables have been set successfully:"
echo "CHANNEL_NAME: $CHANNEL_NAME"
echo "CHAINCODE_NAME: $CHAINCODE_NAME"
echo "CHAINCODE_VERSION: $CHAINCODE_VERSION"
echo "CC_RUNTIME_LANGUAGE: $CC_RUNTIME_LANGUAGE"
echo "CC_SOURCE_PATH: $CC_SOURCE_PATH"
echo "ORDERER_CA: $ORDERER_CA"


#exec:
#peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label ${CHAINCODE_NAME}_${CHAINCODE_VERSION} &>log.txt

