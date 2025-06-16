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

    // 铸造主NFT
    console.log("\n🎨 铸造主NFT...");
    const mainNftUri = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";

    // 用户1铸造一个主NFT
    await mainNFT.connect(user1).mint(mainNftUri);
    console.log(`✅ 用户1铸造了主NFT #0`);

    // 验证主NFT信息
    const owner0 = await mainNFT.ownerOf(0);
    const tokenUri0 = await mainNFT.tokenURI(0);
    console.log(`   所有者: ${owner0}`);
    console.log(`   URI: ${tokenUri0}`);

    // 用HTTP网关方式铸造另一个主NFT
    console.log("\n🌐 使用HTTP网关铸造主NFT...");
    const ipfsUri = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";
    await mainNFT.connect(user2).mintWithHttpGateway(ipfsUri);
    console.log(`✅ 用户2铸造了主NFT #1`);

    const tokenUri1 = await mainNFT.tokenURI(1);
    console.log(`   转换后的URI: ${tokenUri1}`);

    // 创建子NFT
    console.log("\n👶 创建子NFT...");

    // 用户1为用户2创建一个子NFT
    const childUri1 = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";
    await mainNFT.connect(user1).createChildNFTWithURI(0, user2.address, childUri1);
    console.log(`✅ 用户1为用户2创建了子NFT #0`);

    // 验证子NFT信息
    const childOwner0 = await childNFT.ownerOf(0);
    const childCreator0 = await childNFT.getChildCreator(0);
    const parentTokenId0 = await childNFT.getParentTokenId(0);
    console.log(`   子NFT所有者: ${childOwner0}`);
    console.log(`   子NFT创建者: ${childCreator0}`);
    console.log(`   父NFT ID: ${parentTokenId0}`);

    // 创建带自定义URI的子NFT
    console.log("\n🎨 创建带自定义URI的子NFT...");
    const childUri2 = "ipfs://QmChildNFTHash987654321";
    await mainNFT.connect(user1).createChildNFTWithURI(0, user1.address, childUri2);
    console.log(`✅ 用户1为自己创建了带自定义URI的子NFT #1`);

    const childTokenUri1 = await childNFT.tokenURI(1);
    console.log(`   子NFT URI: ${childTokenUri1}`);

    // 获取NFT元数据信息
    console.log("\n📊 获取NFT元数据信息...");

    // 获取主NFT元数据
    console.log("主NFT信息:");
    for (let i = 0; i < 2; i++) {
        try {
            const uri = await mainNFT.tokenURI(i);
            const owner = await mainNFT.ownerOf(i);
            console.log(`   NFT #${i}: ${uri} (所有者: ${owner})`);
        } catch (error) {
            console.log(`   NFT #${i}: 不存在`);
        }
    }

    // 获取子NFT元数据
    console.log("子NFT信息:");
    for (let i = 0; i < 2; i++) {
        try {
            const uri = await childNFT.tokenURI(i);
            const owner = await childNFT.ownerOf(i);
            const creator = await childNFT.getChildCreator(i);
            const parentId = await childNFT.getParentTokenId(i);
            console.log(`   子NFT #${i}: ${uri}`);
            console.log(`       所有者: ${owner}`);
            console.log(`       创建者: ${creator}`);
            console.log(`       父NFT ID: ${parentId}`);
        } catch (error) {
            console.log(`   子NFT #${i}: 不存在`);
        }
    }

    // 演示URI更新
    console.log("\n🔄 演示URI更新...");
    const newUri = "ipfs://QmUpdatedHash111222333";
    await mainNFT.connect(user1).setSpecificTokenURI(0, newUri);
    console.log(`✅ 用户1更新了主NFT #0的URI`);

    const updatedUri = await mainNFT.tokenURI(0);
    console.log(`   新URI: ${updatedUri}`);

    // 显示合约状态
    console.log("\n📈 合约状态统计:");
    const mainNftTotalSupply = await mainNFT.totalSupply();
    const childNftTotalSupply = await childNFT.totalSupply();
    const maxAmount = await mainNFT.MAX_AMOUNT();

    console.log(`   主NFT总供应量: ${mainNftTotalSupply}/${maxAmount}`);
    console.log(`   子NFT总供应量: ${childNftTotalSupply}`);

    // 余额信息
    console.log("\n💰 用户余额:");
    console.log(`   用户1主NFT余额: ${await mainNFT.balanceOf(user1.address)}`);
    console.log(`   用户1子NFT余额: ${await childNFT.balanceOf(user1.address)}`);
    console.log(`   用户2主NFT余额: ${await mainNFT.balanceOf(user2.address)}`);
    console.log(`   用户2子NFT余额: ${await childNFT.balanceOf(user2.address)}`);

    console.log("\n🎉 交互演示完成!");
    console.log("\n📝 合约地址记录:");
    console.log(`MainNFT: ${await mainNFT.getAddress()}`);
    console.log(`ChildNFT: ${await childNFT.getAddress()}`);
}

main().catch((error) => {
    console.error("❌ 执行出错:", error);
    process.exitCode = 1;
}); 