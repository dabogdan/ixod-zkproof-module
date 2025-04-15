# ixo Blockchain

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/ixofoundation/ixo-blockchain?color=white&label=release&style=flat-square)](https://github.com/ixofoundation/ixo-blockchain/releases/latest) ![GitHub Release Date](https://img.shields.io/github/release-date/ixofoundation/ixo-blockchain?label=date&color=white&style=flat-square) [![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/ixofoundation/ixo-blockchain?color=00d2ff&include_prereleases&label=candidate&style=flat-square)](https://github.com/ixofoundation/ixo-blockchain/releases/)

[![GitHub license](https://img.shields.io/github/license/ixofoundation/ixo-blockchain?color=lightgrey&style=flat-square)](https://github.com/ixofoundation/ixo-blockchain/blob/main/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ixofoundation/ixo-blockchain?color=lightgrey&style=flat-square) ![GitHub repo size](https://img.shields.io/github/repo-size/ixofoundation/ixo-blockchain?color=lightgrey&style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/ixofoundation/ixo-blockchain)](https://goreportcard.com/report/github.com/ixofoundation/ixo-blockchain)

[![Discord](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/invite/ixo) [![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/ixonetwork)
[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/ixoworld)
[![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/ixo-blog)

<p align="center">
  <img src="./.github/assets/readme_banner.png" />
</p>
<br />

# zkProof Module for Cosmos SDK

The `zkproof` module enables privacy-preserving commitments using zero-knowledge proofs in Cosmos SDK applications.

Instead of submitting sensitive data, users submit a **hash of a zero-knowledge proof** along with a `claim_type` and optional `public_input` that is stored on-chain. The **actual proof generation and verification are done off-chain**.

---

## Module Features

- Accepts transactions (`MsgSubmitProof`) that:
  - Store the **creator address**
  - Store a **claim type** (e.g., `age-verification`)
  - Store the **ZK proof hash** (`proof_hash`)
  - Store a **public input** (e.g., `"age>18"`)

- Provides **query commands** to retrieve these stored values.

- Does not verify ZK proofs on-chain — assumes off-chain proof generation/verification.

- 💡 Designed for integration with off-chain ZK servers and potential CosmWasm-based ZK verifiers in future.

---

## Requirements

| Dependency       | Version     |
|------------------|-------------|
| Go               | `v1.20+`    |
| Cosmos SDK       | `v0.50+`    |
| CometBFT         | Compatible  |
| Protobuf         | Installed (`protoc`) with Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`) |
| Docker           | Optional for CI or isolated dev builds |

---

## Build

```bash
make build
```

My zkProof Module:
* Accepts MsgSubmitProof transactions
* Stores creator, claim_type, proof_hash and public_input
* Lets users query those stored values
* Does not verify or generate the ZK proof — it assumes that happens off-chain
* Assumes that verifiers trust the off-chain prover

## Procedure: Run blockchain node locally without Docker

### 1. Clear previous state
```bash
rm -rf ~/.ixod
```
### 2. Initialize the chain with chain ID and token denom
```bash
ixod init test-node --chain-id=ixo-local --default-denom=uixo
```
### 3. Create the 'alice' key (test keyring backend is non-persistent, great for dev) 
```bash
ixod keys add alice --keyring-backend test
```
### 4. Add genesis account for Alice with a large balance
```bash
ixod genesis add-genesis-account $(ixod keys show alice -a --keyring-backend test) 1000000000000uixo
```
### 5. Generate gentx for Alice (this sets Alice as a validator)
```bash
ixod genesis gentx alice 500000000000uixo --chain-id=ixo-local --amount=500000000000uixo --keyring-backend=test
```
### 6. Combine all gentxs (only needed for multi-validator setups, but required)
```bash
ixod genesis collect-gentxs
```
### 7. Start the blockchain node
```bash
ixod start
```
### 8. Run tests
```bash
unit tests: go test ./x/zkproof/test/...
```

## Module use case: 
I'm verified that I am over 18 years old, but won’t show my ID. The proof is generated offline, some data of it is stored online for everybody to be able to query it.

## Command to test xkProof module:

1) SUBMIT PROOF
```bash
 get balances: ixod query bank balances $(ixod keys show alice -a --keyring-backend test)
```

(alice should have a balance as per instructions above)
```bash
ixod tx zkproof submit-proof age-verification zk-proof-hash "age>18" \
  --from alice \
  --keyring-backend test \
  --chain-id ixo-local \
  --gas auto \
  --gas-adjustment 1.5 \
  --fees 2000uixo \
  --broadcast-mode sync \
  -y
```

to check the transaction: 
```bash 
ixod query tx <txhash>
```

2) QUERY PROOF
```bash
ixod query zkproof proof age-verification --creator $(ixod keys show alice -a --keyring-backend test)
```

## Module Folder Structure
```bash
x/zkproof/
├── client/
│   └── cli/               # CLI commands: submit-proof, query proof
├── keeper/
│   ├── keeper.go          # Core logic for storing/retrieving proofs
│   ├── msg_server.go      # Handles MsgSubmitProof
│   └── query_server.go    # Handles gRPC proof queries
├── types/
│   ├── messages.go        # MsgSubmitProof definition
│   ├── query.proto        # gRPC query definitions
│   ├── keeper.proto       # Data types
├── test/
│   ├── keeper_test.go     # Unit test for storage
│   ├── msg_server_test.go # Unit test for message handling
│   └── query_server_test.go # Unit test for querying
```