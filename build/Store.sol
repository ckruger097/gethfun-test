// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.3 and less than 0.8.0
pragma solidity ^0.8.3;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}