#!/usr/bin/env bash
set -euo pipefail

# Make sure to run local_node.sh at the repository root before running this script

RPC_URL="${RPC_URL:-http://127.0.0.1:8545}"
REST_URL="${REST_URL:-http://127.0.0.1:1317}"
PRECOMPILE="${PRECOMPILE:-0x0000000000000000000000000000000000001111}"
AMOUNT="${AMOUNT:-1000000}"

ACCOUNT="${ACCOUNT:-$(curl -s -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' "$RPC_URL" | jq -r '.result[0] // empty')}"
DENOM="${DENOM:-$(curl -s "$REST_URL/cosmos/bank/v1beta1/denoms_metadata" | jq -r '.metadatas[0].base // empty')}"
[[ -n "$ACCOUNT" && -n "$DENOM" ]] || { echo "set ACCOUNT and/or DENOM"; exit 1; }

before_hex="$(curl -s -H 'Content-Type: application/json' --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"$ACCOUNT\",\"latest\"],\"id\":1}" "$RPC_URL" | jq -r '.result')"
before="$(cast to-dec "$before_hex")"

data="$(cast calldata "mint(address,string,uint256)" "$ACCOUNT" "$DENOM" "$AMOUNT")"
tx="$(curl -s -H 'Content-Type: application/json' --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_sendTransaction\",\"params\":[{\"from\":\"$ACCOUNT\",\"to\":\"$PRECOMPILE\",\"gas\":\"0x493e0\",\"data\":\"$data\"}],\"id\":1}" "$RPC_URL")"
hash="$(echo "$tx" | jq -r '.result // empty')"
[[ -n "$hash" ]] || { echo "$tx" | jq .; exit 1; }

for _ in $(seq 1 60); do
  receipt="$(curl -s -H 'Content-Type: application/json' --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$hash\"],\"id\":1}" "$RPC_URL")"
  status="$(echo "$receipt" | jq -r '.result.status // empty')"
  [[ -n "$status" ]] && break
  sleep 1
done
[[ "${status:-}" == "0x1" ]] || { echo "$receipt" | jq .; exit 1; }

after_hex="$(curl -s -H 'Content-Type: application/json' --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"$ACCOUNT\",\"latest\"],\"id\":1}" "$RPC_URL" | jq -r '.result')"
after="$(cast to-dec "$after_hex")"

delta="$(echo "$after - $before" | bc)"

echo "account=$ACCOUNT denom=$DENOM amount=$AMOUNT tx=$hash"
echo "before=$before after=$after delta=$delta"
