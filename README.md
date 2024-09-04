# File Registry

## Summary

- Create a file registry on-chain. Upload the files to the IPFS and then store the CID in a smart contract against an upload ID.

## Requirements

- Implement a File Registry based on the smart contract that allows users to store file information on-chain. Contract should support following read and write operations:

    1. function save(string filePath, string cid) - save the data in the registry   
    2. function get(string filePath) public view returns (string) - gets the cid of the file using the filePath

- Implement an API service with the following routes:

    1. POST /v1/files
        - Argument: file and filePath
        - Uploads it to the IPFS node using one of the client libraries. IPFS client will return the cid for the uploaded file.
        - Saves the data in the File Registry.

    2. GET /v1/files?filePath=/dumy.txt
        - Argument: filePath as queryString
        - Get the data from the File Registry and return the CID to the client.

## Run on local environment
1. Install go.

2. Download required packages.
    ``` bash
    go mod download
    ```
3. Install swagger.
    ``` bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
4. Generate swagger documentation.
    ``` bash
    swag init -g cmd/main.go
    ```
5. Copy `.env.example` and rename it to `.env`.

6. Run application.
    ``` bash
    go run ./cmd
    ```
7. You can see swagger documentation at <http://localhost:8000/docs/index.html>.

solc --abi contracts/FileRegistry.sol -o contracts

abigen --abi contracts/FileRegistry.abi --pkg main --type Storage --out contracts/FileRegistory.go

## Guidelines

1. We are interested in your code management and development skills, so please solve the problem keeping that in mind.
2. The API service needs to be developed in Golang.
3. You are free to use any Ethereum test network, or a local environment (eg. ganache)
4. Please ensure that the coding conventions, directory structure and build approach to your project
5. The codebase should contain a docker-compose file that sets up the API and IPFS node in a docker environment.

## Resources

- Deploying Smart Contracts
    - <https://ethereumdev.io/deploying-your-first-smart-contract/>
    - <https://ethereum.org/en/developers/docs/smart-contracts/deploying/>
    - <https://docs.openzeppelin.com/learn/deploying-and-interacting>

- Running an IPFS Node
    - <https://docs.ipfs.tech/install/run-ipfs-inside-docker/#set-up>

- Go IPFS Client
    - <https://github.com/ipfs/kubo/tree/master/client/rpchttps://github.com/ipfs/go-ipfs-api>

# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.js
```
