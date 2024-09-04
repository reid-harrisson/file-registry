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

## Run Application
1. Install docker.
2. Copy `backend/.env.example` and rename it to `.env`.
3. Start docker container.

    ```bash
    make up
    ```

    If it doesn't work, use this command.

    ```bash
    docker compose up
    ```

4. Stop docker container.

    ```bash
    make down
    ```

    If it doesn't work, use this comand

    ```bash
    docker compose down
    ```
5. You can see swagger documentation at this url.

    <http://127.0.0.1:8000/docs/index.html>

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