#This is a script with the commands for automate the process to:
# - Create the channel.
# - Include all the organizations into the created channel.
# - Updating the channels with the anchor peers cinfiguration for every organization.

export CHANNEL_NAME=marketplace # replace marketplace with [Channel name to create]

# CREATE THE CHANNEL:
# Ensure to use the /msp/tlscacerts/tlsca.acem.com-cert.pem to avoid erros
# Do not confuse with the /ca/ cert

channel create \
  -o orderer.acme.com:7050 \
  -c $CHANNEL_NAME \
  -f ./channel-artifacts/channel.tx \
  --tls true \
  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/acme.com/msp/tlscacerts/tlsca.acme.com-cert.pem

peer channel join -b marketplace.block #[channel.block to join org to the network]

#For 2nd and N organization:

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.acme.com/users/Admin@org2.acme.com/msp \
CORE_PEER_ADDRESS=peer0.org2.acme.com:7051 \
CORE_PEER_LOCALMSPID="Org2MSP" \
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt \
peer channel join -b marketplace.block


#Configuration for the anchor peers:
#(Edit the Organization name)

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.acme.com/users/Admin@org3.acme.com/msp \
CORE_PEER_ADDRESS=peer0.org3.acme.com:7051 \
CORE_PEER_LOCALMSPID="Org3MSP" \
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.acme.com/tls/ca.crt \
peer channel update \
-o orderer.acme.com:7050 \
-c $CHANNEL_NAME \
-f ./channel-artifacts/Org3MSPanchors.tx \
--tls \
--cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/acme.com/orderers/orderer.acme.com/msp/tlscacerts/tlsca.acme.com-cert.pem

