// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FileRegistry {
    mapping(string => string) private fileStore;

    // Save file information in the registry
    function save(string memory filePath, string memory cid) public {
        fileStore[filePath] = cid;
    }

    // Retrieve the CID using the file path
    function get(string memory filePath) public view returns (string memory) {
        return fileStore[filePath];
    }
}
