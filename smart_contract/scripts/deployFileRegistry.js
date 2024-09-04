const hre = require("hardhat");

async function deploy() {
  const FileRegistry = await hre.ethers.getContractFactory("FileRegistry");
  const fileRegistry = await FileRegistry.deploy();
  await fileRegistry.waitForDeployment();
  console.log("FileRegistry deployed at:", fileRegistry.target);
}

deploy()
  .then(() => console.log("Deployment completed!"))
  .catch((error) => console.error("Error deploying contract:", error));
