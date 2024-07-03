// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Counters.sol";
import "./NFT.sol";
import "./Token.sol";

contract Controller {
    using Counters for Counters.Counter;

    //
    // STATE VARIABLES
    //
    Counters.Counter private _sessionIdCounter;
    GeneNFT public geneNFT;
    PostCovidStrokePrevention public pcspToken;

    struct UploadSession {
        uint256 id;
        address user;
        string proof;
        bool confirmed;
    }

    struct DataDoc {
        string id;
        string hashContent;
    }

    mapping(uint256 => UploadSession) sessions;
    mapping(string => DataDoc) docs;
    mapping(string => bool) docSubmits;
    mapping(uint256 => string) nftDocs;

    //
    // EVENTS
    //
    event UploadData(string docId, uint256 sessionId);
    event Confirm (string docId, uint256 sessionId,uint256 riskScore);

    constructor(address nftAddress, address pcspAddress) {
        geneNFT = GeneNFT(nftAddress);
        pcspToken = PostCovidStrokePrevention(pcspAddress);
    }

    function uploadData(string memory docId) public returns (uint256) {
        // TODO: Implement this method: to start an uploading gene data session. The doc id is used to identify a unique gene profile. Also should check if the doc id has been submited to the system before. This method return the session id
        require(bytes(docId).length > 0,"docId empty not allow ");
        require(!docSubmits[docId] ,"Doc already been submitted" );

        uint256 _currentSessionId = _sessionIdCounter.current();
        _sessionIdCounter.increment();

        sessions[_currentSessionId] = UploadSession(_currentSessionId, msg.sender, "", false);
        

        emit UploadData(docId,_currentSessionId);
        return _currentSessionId;
    }

    function confirm(
        string memory docId,
        string memory contentHash,
        string memory proof,
        uint256 sessionId,
        uint256 riskScore
    ) public {
        require(bytes(docId).length > 0,"docId empty not allow ");
        require(bytes(contentHash).length > 0,"docId empty not allow ");
        require(bytes(proof).length > 0,"docId empty not allow ");
        
        require(!docSubmits[docId] ,"Doc already been submitted" );
        
        require(sessions[sessionId].user == msg.sender,"Invalid session owner");

        require(!sessions[sessionId].confirmed,"Session is ended");

        // TODO: Implement this method: The proof here is used to verify that the result is returned from a valid computation on the gene data. For simplicity, we will skip the proof verification in this implementation. The gene data's owner will receive a NFT as a ownership certicate for his/her gene profile.

        // TODO: Verify proof, we can skip this step

        // TODO: Update doc content
        docs[docId] = DataDoc(docId, contentHash);
        // TODO: Mint NFT
        geneNFT.safeMint(msg.sender);
        // TODO: Reward PCSP token based on risk stroke
        pcspToken.reward(msg.sender, riskScore);
        // TODO: Close session
        docSubmits[docId] = true;
        sessions[sessionId].confirmed = true;
        sessions[sessionId].proof = proof;
        emit Confirm(docId, sessionId, riskScore);
    }

    function getSession(uint256 sessionId) public view returns(UploadSession memory) {
        return sessions[sessionId];
    }

    function getDoc(string memory docId) public view returns(DataDoc memory) {
        return docs[docId];
    }
}