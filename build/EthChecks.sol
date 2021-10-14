// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.3 and less than 0.8.0
pragma solidity ^0.8.3;

contract EthChecks {



    bytes32 private hashedSecret = 0x27edca7b1e5bb188201b8a1e971e153b6f8fb0906d75a7d6a8133b93ad1df65f;

    constructor()

    function EthChecks(string memory _word) public view returns(bool){
        return keccak256(abi.encodePacked(_word)) == answer;
    }
}
