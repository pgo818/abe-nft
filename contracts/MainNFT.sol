// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
//ipfs://
contract MainNFT is ERC721, ERC721Enumerable, ERC721URIStorage, Ownable {
    uint256 public MAX_AMOUNT = 10000;
 
    uint256 private _nextTokenId;
    string private _baseTokenURI;
    
    // 子NFT合约地址
    address public childNFTContract;
    
    event ChildNFTCreated(uint256 indexed tokenId, address indexed receiver);

    constructor() ERC721("MainABE", "MABE") Ownable(msg.sender) {}
    
    // 设置子NFT合约地址
    function setChildNFTContract(address _childNFTContract) external onlyOwner {
        childNFTContract = _childNFTContract;
    }

    // 铸造主NFT - 接受URI参数
    function mintTo(address to,string memory uri) public payable {
        require(totalSupply() < MAX_AMOUNT, "NFT is sold out!");
        
        uint256 tokenId = _nextTokenId;
        _nextTokenId++;
        
        _safeMint(to, tokenId);
        
        // 直接设置传入的URI
        _setTokenURI(tokenId, uri);
    }
    

    
    // 从IPFS URI中提取哈希值
    function extractIPFSHash(string memory ipfsURI) public pure returns (string memory) {
        bytes memory uriBytes = bytes(ipfsURI);
        
        // 检查URI是否以"ipfs://"开头
        if (uriBytes.length < 7) return ipfsURI;
        
        bytes memory ipfsPrefix = bytes("ipfs://");
        bool isIPFS = true;
        
        for (uint i = 0; i < 7; i++) {
            if (ipfsPrefix[i] != uriBytes[i]) {
                isIPFS = false;
                break;
            }
        }
        
        if (!isIPFS) return ipfsURI;
        
        // 提取哈希部分（去掉"ipfs://"前缀）
        string memory hash = "";
        for (uint i = 7; i < uriBytes.length; i++) {
            hash = string(abi.encodePacked(hash, bytes1(uriBytes[i])));
        }
        
        return hash;
    }
    
    // 使用HTTP网关URL铸造NFT
    function mintWithHttpGateway(address to,string memory ipfsURI) public payable {
        require(totalSupply() < MAX_AMOUNT, "NFT is sold out!");
        
        uint256 tokenId = _nextTokenId;
        _nextTokenId++;
        
        _safeMint(to, tokenId);
        
        // 提取IPFS哈希并构建HTTP网关URL
        string memory hash = extractIPFSHash(ipfsURI);
        string memory httpGatewayURI = string(abi.encodePacked("https://ipfs.io/ipfs/", hash));
        
        // 使用HTTP网关URL
        _setTokenURI(tokenId, httpGatewayURI);
    }
    // 通过主NFT为指定地址铸造子NFT并设置自定义URI
    function createChildNFTWithURI(address to,uint256 tokenId, address receiver, string memory uri) public {
        require(_exists(tokenId), "Token does not exist");
        require(ownerOf(tokenId) == to, "You must own the token");
        require(childNFTContract != address(0), "Child NFT contract not set");
        
        // 调用子合约的带URI的铸造函数
        (bool success, ) = childNFTContract.call(
            abi.encodeWithSignature(
                "mintChildNFTWithURI(uint256,address,address,string)", 
                tokenId, 
                receiver, 
                msg.sender,
                uri
            )
        );
        
        require(success, "Child NFT creation with URI failed");
        
        emit ChildNFTCreated(tokenId, receiver);
    }
       
    // 设置特定NFT的URI
    function setSpecificTokenURI(address to,uint256 tokenId, string memory newTokenURI) public {
        require(_exists(tokenId), "URI set of nonexistent token");
        require(ownerOf(tokenId) == to || to == owner(), "Not authorized");
        
         _setTokenURI(tokenId, newTokenURI);

    }
    
    // 检查token是否存在的内部函数
    function _exists(uint256 tokenId) internal view returns (bool) {
        return tokenId < _nextTokenId && _ownerOf(tokenId) != address(0);
    }

    // The following functions are overrides required by Solidity
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function _increaseBalance(address account, uint128 value) 
        internal 
        override(ERC721, ERC721Enumerable) 
    {
        super._increaseBalance(account, value);
    }
    
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }
        
    // 提取合约资金的函数
    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        payable(msg.sender).transfer(balance);
    }
} 