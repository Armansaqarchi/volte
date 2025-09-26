// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract VolteContract {

    address public admin;

    constructor() {
        admin = msg.sender;

    }

    mapping (string /* eventID */ => bytes /* NullifierRootHash */) public nullifierMerkleRoots;
    mapping (string /* eventID */ => bytes /* VoteRootHash */) public voteMerkleRoots;
    mapping (string /* eventID */ => bytes /* EventDetailsHash */) public eventHashes;
    mapping (string /* eventID */ => bytes /* UsedNullifiers */) public usedNullifiers;

//    function verifyNullifierProof() external {
//        // verifies nullifier proof given a leaf hash and set of
//        // siblings hash corresponding to the leafs path up to the root.
//
//        // Match the calculated root against rootHash in `nullifierMerkleRoots`
//
//    }
//
//
//    function verifyEncodedCipherText() external{
//
//    }

    modifier onlyOwner() {
        require(msg.sender == admin, "Only owner is allowed to execute this transaction.");
        _;
    }


    function SetNullifierMerkleRoot(string calldata eventID, bytes calldata value) onlyOwner external {
        //check before to make sure its a valid hash
       nullifierMerkleRoots[eventID] = value;
    }

    function SetVoteMerkleRoot(string calldata eventID, bytes calldata value) onlyOwner external {
        // check before to make sure its a valid hash
        voteMerkleRoots[eventID] = value;
    }

    function SetEventHash(string calldata eventID, bytes calldata value) onlyOwner external {
        // A change in event details changes the event hash completely.
        eventHashes[eventID] = value;
    }

    function GetNullifierMerkleRoot(string calldata eventID) external view returns (bytes memory) {
        return nullifierMerkleRoots[eventID];
    }

    function GetVoteMerkleRoot(string calldata eventID) external view returns (bytes memory) {
        return voteMerkleRoots[eventID];
    }

    function GetEventHash(string calldata eventID) external view returns (bytes memory) {
        return eventHashes[eventID];
    }

}