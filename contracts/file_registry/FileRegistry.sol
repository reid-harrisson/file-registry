// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

contract FileRegistry {
    mapping(string => string) cidForPath;

    constructor() {}

    // save cid for file path
    function save(string calldata filePath, string calldata cid) public {
        require(bytes(filePath).length != 0, "Invalid file path!");
        require(bytes(cid).length != 0, "Invalid cid!");
        cidForPath[filePath] = cid;
    }

    // get cid for file path
    function get(string calldata filePath) public view returns (string memory) {
        require(bytes(filePath).length != 0, "Invalid file path!");
        return cidForPath[filePath];
    }
}
