#!/bin/sh
set -eu

DATADIR=/data/geth
GENESIS=/data/genesis.json
ETHERBASE=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

PRIVKEY_HEX=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
KEYPASS=devpass

# Initialize only once
if [ ! -d "$DATADIR/geth/chaindata" ]; then
  echo "Initializing chain from genesis..."
  geth init --datadir "$DATADIR" "$GENESIS"
fi

# Import signer key once
if [ ! -d "$DATADIR/keystore" ] || [ -z "$(ls -A "$DATADIR/keystore" 2>/dev/null || true)" ]; then
  echo "Importing dev private key into geth keystore..."
  echo -n "$KEYPASS" > /tmp/password.txt
  echo -n "$PRIVKEY_HEX" > /tmp/key.txt
  geth account import --datadir "$DATADIR" --password /tmp/password.txt /tmp/key.txt
  rm -f /tmp/password.txt /tmp/key.txt
fi

# Need a password file for unlock at runtime
echo -n "$KEYPASS" > /tmp/password.txt

echo "Starting geth (Clique PoA)..."
exec geth \
  --datadir "$DATADIR" \
  --networkid 1337 \
  --http --http.addr 0.0.0.0 --http.port 8545 \
  --http.api eth,net,web3,txpool,debug,personal \
  --http.corsdomain "*" \
  --http.vhosts "*" \
  --ws --ws.addr 0.0.0.0 --ws.port 8546 \
  --ws.api eth,net,web3 \
  --mine \
  --miner.etherbase "$ETHERBASE" \
  --unlock "$ETHERBASE" \
  --password /tmp/password.txt \
  --allow-insecure-unlock \
  --nodiscover \
  --ipcdisable \
  --verbosity 3
