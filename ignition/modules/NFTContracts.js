const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("NFTContractsModule", (m) => {
    // 部署MainNFT合约
    const mainNFT = m.contract("MainNFT", []);

    // 部署ChildNFT合约
    const childNFT = m.contract("ChildNFT", []);

    // 设置合约之间的关联 - 先设置MainNFT中的ChildNFT地址
    m.call(mainNFT, "setChildNFTContract", [childNFT], {
        after: [mainNFT, childNFT]
    });

    // 然后设置ChildNFT中的MainNFT地址
    m.call(childNFT, "setMainNFTContract", [mainNFT], {
        after: [mainNFT, childNFT]
    });

    return { mainNFT, childNFT };
}); 