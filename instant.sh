instantiateChaincode() {
    ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/grooo.com/orderers/orderer.grooo.com/msp/tlscacerts/tlsca.grooo.com-cert.pem

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer chaincode instantiate -o orderer.grooo.com:7050 -C generalchannel -n mycc -l go -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "AND ('Grooo1MSP.peer','Grooo2MSP.peer')"
    res=$?
    set +x
  else
    set -x
    peer chaincode instantiate -o orderer.grooo.com:7050 --tls true --cafile $ORDERER_CA -C generalchannel -n mycc -l go -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "AND ('Grooo1MSP.peer','Grooo2MSP.peer')"
    res=$?
    set +x
  fi
  echo
}