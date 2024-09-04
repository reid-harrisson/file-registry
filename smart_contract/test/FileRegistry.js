const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FileRegistry", function () {
  let addr;

  beforeEach(async function () {
    // Deploy the contract
    FileRegistry = await ethers.getContractFactory("FileRegistry");
    [addr] = await ethers.getSigners();
    fileRegistry = await FileRegistry.deploy();
  });

  it("Should return correct cid for file path", async function () {
    const filePath = "1.txt";
    const cid = "abcdef";
    await fileRegistry.connect(addr).save(filePath, cid);
    const cidFromPath = await fileRegistry.get(filePath);
    expect(cidFromPath).to.equal(cid);
  });

  it("Should update cid for the same file path", async function () {
    const filePath = "1.txt";
    const cid = "abcdef";
    const updatedCid = "bcdef";
    await fileRegistry.connect(addr).save(filePath, cid);
    await fileRegistry.connect(addr).save(filePath, updatedCid);
    const cidFromPath = await fileRegistry.get(filePath);
    expect(cidFromPath).to.equal(updatedCid);
  });

  it("Should return empty cid for non registered file path", async function () {
    const filePath = "1.txt";
    const cid = "abcdef";
    const nonRegisteredFilePath = "2.txt";
    const emptyString = "";
    await fileRegistry.connect(addr).save(filePath, cid);
    const cidFromPath = await fileRegistry.get(nonRegisteredFilePath);
    expect(cidFromPath).to.equal(emptyString);
  });

  it("Should fail to register an invalid cid", async function () {
    const filePath = "1.txt";
    const invalidCid = "";
    await expect(
      fileRegistry.connect(addr).save(filePath, invalidCid)
    ).to.be.revertedWith("Invalid cid!");
  });
});
