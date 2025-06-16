// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
//ipfs://QmRA97iTYaU6JqxVB8fUG2jtWA1Y1NszbEE9i6uXNfRGsA
//0x7f9a4F4D8dEC6A733883D4Aedb5FCE323B22673C
interface IMainNFT {
    function ownerOf(uint256 tokenId) external view returns (address);
}

contract ChildNFT is ERC721, ERC721Enumerable, ERC721URIStorage, Ownable {
    uint256 private _nextTokenId;
    string private _baseTokenURI; 
    // 主NFT合约地址
    address public mainNFTContract;
    
    // 子NFT与创建者的映射
    mapping(uint256 => address) private _childCreators; // 子NFT ID => 创建者(主NFT所有者)
    
    // 子NFT与主NFT的映射
    mapping(uint256 => uint256) private _parentTokens; // 子NFT ID => 主NFT ID
    
    event ChildTokenMinted(uint256 indexed parentTokenId, uint256 indexed childTokenId, address indexed receiver);
    
    constructor() ERC721("ChildABE", "CABE") Ownable(msg.sender) {}
    
    
    // 设置主NFT合约地址
    function setMainNFTContract(address _mainNFTContract) external onlyOwner {
        mainNFTContract = _mainNFTContract;
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
    // 为指定地址铸造子NFT并设置自定义URI
    function mintChildNFTWithURI(
        uint256 parentTokenId,
        address receiver,
        address creator,
        string memory uri
    ) external {
        // 确保调用者是主NFT合约
        require(msg.sender == mainNFTContract, "Only main NFT contract can call this");
        
        uint256 childTokenId = _nextTokenId;
        _nextTokenId++;
        
        // 铸造子NFT给指定接收者
        _safeMint(receiver, childTokenId);

        // 提取IPFS哈希并构建HTTP网关URL
        string memory hash = extractIPFSHash(uri);
        string memory httpGatewayURI = string(abi.encodePacked("https://ipfs.io/ipfs/", hash));
        // 设置自定义元数据URI
        _setTokenURI(childTokenId, httpGatewayURI);
       
        // 记录关联关系
        _childCreators[childTokenId] = creator; // 记录创建者(主NFT所有者)
        _parentTokens[childTokenId] = parentTokenId; // 记录关联的主NFT
        
        emit ChildTokenMinted(parentTokenId, childTokenId, receiver);
    }
      
    
    // 获取子NFT的创建者(主NFT所有者)
    function getChildCreator(uint256 childTokenId) public view returns (address) {
        require(_exists(childTokenId), "Child token does not exist");
        return _childCreators[childTokenId];
    }
    
    // 获取子NFT对应的主NFT ID
    function getParentTokenId(uint256 childTokenId) public view returns (uint256) {
        require(_exists(childTokenId), "Child token does not exist");
        return _parentTokens[childTokenId];
    }
    
    // 设置子NFT的URI
    function setChildTokenURI(string memory newURI) public onlyOwner {
        _baseTokenURI = newURI;
    }
    
    // 设置某个特定NFT的URI
    function setSpecificTokenURI(uint256 tokenId, string memory newTokenURI) public {
        require(_exists(tokenId), "URI set of nonexistent token");
        require(msg.sender == _childCreators[tokenId] || msg.sender == owner(), "Not authorized");
        
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
    
    // 重写update函数，简化权限检查
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address)
    {
        // 简单处理，使用OpenZeppelin标准权限检查
        return super._update(to, tokenId, auth);
    }
} 