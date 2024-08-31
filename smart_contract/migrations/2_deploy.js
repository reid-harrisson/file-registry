const FileRegistry = artifacts.require('FileRegistry')

module.exports = async function (deployer) {
    await deployer.deploy(FileRegistry)
}