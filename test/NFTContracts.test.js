const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("NFT Contracts", function () {
    let mainNFT;
    let childNFT;
    let owner;
    let addr1;
    let addr2;

    beforeEach(async function () {
        [owner, addr1, addr2] = await ethers.getSigners();

        // 部署MainNFT合约
        const MainNFT = await ethers.getContractFactory("MainNFT");
        mainNFT = await MainNFT.deploy();

        // 部署ChildNFT合约
        const ChildNFT = await ethers.getContractFactory("ChildNFT");
        childNFT = await ChildNFT.deploy();

        // 设置合约之间的关联
        await mainNFT.setChildNFTContract(await childNFT.getAddress());
        await childNFT.setMainNFTContract(await mainNFT.getAddress());
    });

    describe("MainNFT Deployment", function () {
        it("Should set the right name and symbol", async function () {
            expect(await mainNFT.name()).to.equal("MainABE");
            expect(await mainNFT.symbol()).to.equal("MABE");
        });

        it("Should set the right owner", async function () {
            expect(await mainNFT.owner()).to.equal(owner.address);
        });

        it("Should set the correct max amount", async function () {
            expect(await mainNFT.MAX_AMOUNT()).to.equal(3);
        });

        it("Should correctly set child contract address", async function () {
            expect(await mainNFT.childNFTContract()).to.equal(await childNFT.getAddress());
        });
    });

    describe("ChildNFT Deployment", function () {
        it("Should set the right name and symbol", async function () {
            expect(await childNFT.name()).to.equal("ChildABE");
            expect(await childNFT.symbol()).to.equal("CABE");
        });

        it("Should set the right owner", async function () {
            expect(await childNFT.owner()).to.equal(owner.address);
        });

        it("Should correctly set main contract address", async function () {
            expect(await childNFT.mainNFTContract()).to.equal(await mainNFT.getAddress());
        });
    });

    describe("MainNFT Minting", function () {
        it("Should mint a token with URI", async function () {
            const testURI = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";

            await mainNFT.connect(addr1).mint(testURI);

            expect(await mainNFT.balanceOf(addr1.address)).to.equal(1);
            expect(await mainNFT.ownerOf(0)).to.equal(addr1.address);
            expect(await mainNFT.tokenURI(0)).to.equal(testURI);
            expect(await mainNFT.totalSupply()).to.equal(1);
        });

        it("Should mint with HTTP gateway conversion", async function () {
            const ipfsURI = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";

            await mainNFT.connect(addr1).mintWithHttpGateway(ipfsURI);

            const expectedHttpURI = "https://ipfs.io/ipfs/QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";
            expect(await mainNFT.tokenURI(0)).to.equal(expectedHttpURI);
        });

        it("Should not exceed max supply", async function () {
            const testURI = "ipfs://QmTest";

            // 铸造3个NFT（达到上限）
            await mainNFT.connect(addr1).mint(testURI);
            await mainNFT.connect(addr1).mint(testURI);
            await mainNFT.connect(addr1).mint(testURI);

            // 尝试铸造第4个应该失败
            await expect(
                mainNFT.connect(addr1).mint(testURI)
            ).to.be.revertedWith("NFT is sold out!");
        });

        it("Should extract IPFS hash correctly", async function () {
            const ipfsURI = "ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";
            const expectedHash = "QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA";

            expect(await mainNFT.extractIPFSHash(ipfsURI)).to.equal(expectedHash);
        });
    });

    describe("Child NFT Creation", function () {
        beforeEach(async function () {
            // 先铸造一个主NFT
            await mainNFT.connect(addr1).mint("ipfs://QmMainNFT");
        });

        it("Should create child NFT with custom URI", async function () {
            const childURI = "ipfs://QmChildNFT";

            await mainNFT.connect(addr1).createChildNFTWithURI(0, addr2.address, childURI);

            expect(await childNFT.balanceOf(addr2.address)).to.equal(1);
            expect(await childNFT.ownerOf(0)).to.equal(addr2.address);
            expect(await childNFT.getParentTokenId(0)).to.equal(0);
            expect(await childNFT.getChildCreator(0)).to.equal(addr1.address);

            const expectedHttpURI = "https://ipfs.io/ipfs/QmChildNFT";
            expect(await childNFT.tokenURI(0)).to.equal(expectedHttpURI);
        });

        it("Should not allow non-owner to create child NFT", async function () {
            const childURI = "ipfs://QmChildNFT";

            await expect(
                mainNFT.connect(addr2).createChildNFTWithURI(0, addr2.address, childURI)
            ).to.be.revertedWith("You must own the token");
        });

        it("Should not create child NFT for non-existent token", async function () {
            const childURI = "ipfs://QmChildNFT";

            await expect(
                mainNFT.connect(addr1).createChildNFTWithURI(999, addr2.address, childURI)
            ).to.be.revertedWith("Token does not exist");
        });
    });

    describe("Token URI Management", function () {
        beforeEach(async function () {
            await mainNFT.connect(addr1).mint("ipfs://QmMainNFT");
        });

        it("Should allow owner to update main NFT URI", async function () {
            const newURI = "ipfs://QmNewMainNFT";

            await mainNFT.connect(addr1).setSpecificTokenURI(0, newURI);

            expect(await mainNFT.tokenURI(0)).to.equal(newURI);
        });

        it("Should allow contract owner to update any NFT URI", async function () {
            const newURI = "ipfs://QmNewMainNFT";

            await mainNFT.connect(owner).setSpecificTokenURI(0, newURI);

            expect(await mainNFT.tokenURI(0)).to.equal(newURI);
        });

        it("Should not allow unauthorized users to update URI", async function () {
            const newURI = "ipfs://QmNewMainNFT";

            await expect(
                mainNFT.connect(addr2).setSpecificTokenURI(0, newURI)
            ).to.be.revertedWith("Not authorized");
        });
    });

    describe("Withdrawal", function () {
        it("Should allow owner to withdraw funds", async function () {
            // 铸造一些NFT来生成资金（如果有支付要求）
            await mainNFT.connect(addr1).mint("ipfs://QmTest");

            const contractBalance = await ethers.provider.getBalance(await mainNFT.getAddress());

            if (contractBalance > 0n) {
                const ownerBalanceBefore = await ethers.provider.getBalance(owner.address);

                const tx = await mainNFT.withdraw();
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed * receipt.gasPrice;

                const ownerBalanceAfter = await ethers.provider.getBalance(owner.address);

                expect(ownerBalanceAfter).to.equal(
                    ownerBalanceBefore + contractBalance - gasUsed
                );
            }
        });

        it("Should prevent non-owner from withdrawing", async function () {
            await expect(
                mainNFT.connect(addr1).withdraw()
            ).to.be.revertedWithCustomError(mainNFT, "OwnableUnauthorizedAccount");
        });
    });
}); 