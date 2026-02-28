#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 || $# -gt 2 ]]; then
  echo "usage: sh ./scripts/foundry/abigen.sh ./path/to/precompile.sol [optional interface]" >&2
  exit 1
fi

SOL_PATH="$1"
if [[ ! -f "$SOL_PATH" ]]; then
  echo "error: file not found: $SOL_PATH" >&2
  exit 1
fi

CONTRACT_NAME="${2:-$(basename "$SOL_PATH" .sol)}"
OUT_FILE="$(dirname "$SOL_PATH")/abi.json"
CONTRACT_REF="${SOL_PATH}:${CONTRACT_NAME}"

if ! command -v forge >/dev/null 2>&1; then
  echo "error: forge is not installed or not in PATH" >&2
  exit 1
fi

if ! command -v jq >/dev/null 2>&1; then
  echo "error: jq is not installed or not in PATH" >&2
  exit 1
fi

TMP_OUT_DIR="$(mktemp -d)"
TMP_CACHE_DIR="$(mktemp -d)"
cleanup() {
  rm -rf "$TMP_OUT_DIR" "$TMP_CACHE_DIR"
}
trap cleanup EXIT

forge inspect --root . --out "$TMP_OUT_DIR" --cache-path "$TMP_CACHE_DIR" "$CONTRACT_REF" abi --json 2>/dev/null | jq . >"$OUT_FILE"

echo "wrote $OUT_FILE from $CONTRACT_REF"
