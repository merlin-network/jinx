# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Blackchain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

apk add bash jq

KEY1="seed0"
CHAINID="brickchain-666"
MONIKER1="seed-0"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
HOMEDIR="/root/.jinxd"
TRACE=""
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json


jinxd init $MONIKER1 -o --chain-id $CHAINID --home "$HOMEDIR"

jq '.app_state["staking"]["params"]["bond_denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["crisis"]["constant_fee"]["denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["evm"]["params"]["evm_denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["mint"]["params"]["mint_denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";

jinxd config set client keyring-backend $KEYRING --home "$HOMEDIR"

jinxd keys add $KEY1 --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"

jinxd genesis add-genesis-account $KEY1 100000000000000000000000000ablack --keyring-backend $KEYRING --home "$HOMEDIR"

jinxd genesis gentx $KEY1 1000000000000000000000ablack --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
