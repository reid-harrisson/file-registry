// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Store {
    mapping(string => string) private fileStore;

    // Event to log when a file is saved
    event FileSaved(string indexed filePath, string cid);

    // Save file information in the registry
    function save(string memory filePath, string memory cid) public {
        require(bytes(filePath).length > 0, "File path cannot be empty");
        require(bytes(cid).length > 0, "CID cannot be empty");
        
        fileStore[filePath] = cid;
        emit FileSaved(filePath, cid);
    }

    // Retrieve the CID using the file path
    function get(string memory filePath) public view returns (string memory) {
        return fileStore[filePath];
    }
}
