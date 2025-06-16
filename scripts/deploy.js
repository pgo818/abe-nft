const hre = require("hardhat");

async function main() {
    console.log("🚀 开始与NFT合约交互...\n");

    // 获取签名者
    const [deployer, user1, user2] = await hre.ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    console.log("用户1地址:", user1.address);
    console.log("用户2地址:", user2.address);
    console.log();

    // 部署合约
    console.log("📦 部署合约...");
    const MainNFT = await hre.ethers.getContractFactory("MainNFT");
    const mainNFT = await MainNFT.deploy();
    await mainNFT.waitForDeployment();
    console.log("MainNFT部署到:", await mainNFT.getAddress());

    const ChildNFT = await hre.ethers.getContractFactory("ChildNFT");
    const childNFT = await ChildNFT.deploy();
    await childNFT.waitForDeployment();
    console.log("ChildNFT部署到:", await childNFT.getAddress());

    // 设置合约关联
    console.log("\n🔗 设置合约关联...");
    await mainNFT.setChildNFTContract(await childNFT.getAddress());
    await childNFT.setMainNFTContract(await mainNFT.getAddress());
    console.log("✅ 合约关联设置完成");

    // 显示合约状态
    console.log("\n📈 合约状态统计:");
    const mainNftTotalSupply = await mainNFT.totalSupply();
    const childNftTotalSupply = await childNFT.totalSupply();
    const maxAmount = await mainNFT.MAX_AMOUNT();

    console.log(`   主NFT总供应量: ${mainNftTotalSupply}/${maxAmount}`);
    console.log(`   子NFT总供应量: ${childNftTotalSupply}`);


    console.log("\n🎉 交互演示完成!");
    console.log("\n📝 合约地址记录:");
    console.log(`MainNFT: ${await mainNFT.getAddress()}`);
    console.log(`ChildNFT: ${await childNFT.getAddress()}`);
}

main().catch((error) => {
    console.error("❌ 执行出错:", error);
    process.exitCode = 1;
}); 