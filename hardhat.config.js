require("@nomicfoundation/hardhat-toolbox");
//npx hardhat run scripts/deploy.js --network ganache
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: {
    hardhat: {
      chainId: 1337
    },
    localhost: {
      url: "http://127.0.0.1:8545",
      chainId: 1337
    },
    ganache: {
      url: "http://127.0.0.1:7545",
      chainId: 1337,
      accounts: {
        count: 10,
        initialIndex: 0,
        mnemonic: "version coast leader private blossom business tunnel emerge photo seek donor mask",
        path: "m/44'/60'/0'/0",
      },
    }
  }
};
