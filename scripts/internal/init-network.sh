#!/bin/bash

cleanup_init_network_sh() {
  cleanup_genesis_sh
}

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
source "$SCRIPT_DIR"/genesis.sh

# arg1: an existing local dir where validator accounts are created, mandatory
init_network() {
  local val_accounts_dir="$1"
  local validators="$2"
  local chain_id="$3"
  local native_currency="$4"
  local val_tokens="$5"
  local val_stake="$6"
  local genesis_accounts_spec="$7"
  local -r wasm_code_path="$8"
  local -r treasury_init_tokens_u128="$9"
  
  node_id_and_val_pubkeys="$(setup_validators "$validators")"
  local final_genesis_file;
  final_genesis_file=$(generate_genesis "$chain_id" "$native_currency" \
                                          "$val_tokens" "$val_stake" \
                                          "$genesis_accounts_spec" \
                                          "$wasm_code_path" "$treasury_init_tokens_u128" \
                                          "$val_accounts_dir" "$node_id_and_val_pubkeys")
  propagate_genesis "$final_genesis_file" "$validators"
}
