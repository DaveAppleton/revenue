pragma solidity ^0.7.0;
// SPDX-License-Identifier: MIT

import "./SafeMath.sol";

// assumes unit (decimals = 0) tokens
contract stakedRevenueToken {
    using SafeMath for uint256;
    uint256 potPerToken;
    uint256 numTokens;
    uint256 pot;
    uint256 pendingPot;
    uint256 tokenBase;

    string  public name = "StakedRevenueContract";

    struct Balance {
        uint256    tokens;
        uint256    lastPotPerToken;
        uint256    accumulatedReward;
    }

    mapping( address => Balance ) balances;

    constructor(uint256 unitToken) {
        tokenBase = unitToken;
    }

    receive() external payable {
        pot = pot.add(msg.value);
        if (numTokens > 0) {
            uint256 amount = msg.value.add(pendingPot);
            pendingPot = 0;
            potPerToken = potPerToken.add(amount.mul(tokenBase).div(numTokens));
        } else {
            pendingPot = pendingPot.add(msg.value);
        }
    }

    function depositTokens(uint256 tokens) external {
        updateBalance();
        numTokens = numTokens.add(tokens);
        Balance storage myBalance = balances[msg.sender];
        // update rewards
        myBalance.tokens = myBalance.tokens.add(tokens);
        myBalance.lastPotPerToken = potPerToken;
    }

    function updateBalance() internal {
        Balance storage myBalance = balances[msg.sender];
        uint256 reward = myBalance.tokens.mul(potPerToken.sub(myBalance.lastPotPerToken)).div(tokenBase);
        myBalance.accumulatedReward = myBalance.accumulatedReward.add(reward);
    }

    // debug function
    function pendingReward(address account) public view returns (uint) {
        Balance storage myBalance = balances[account];
        uint256 reward = myBalance.tokens.mul(potPerToken.sub(myBalance.lastPotPerToken)).div(tokenBase);
        return reward;
    }

    function transferTokens(address to, uint256 amount) internal {

    }

    function withDrawRewards() external {
        _withdrawRewards(false);
    }

    function _withdrawRewards(bool includingTokens) internal {
        updateBalance();
        Balance storage myBalance = balances[msg.sender];
        uint256 toPay = myBalance.accumulatedReward;
        myBalance.accumulatedReward = 0;
        if (includingTokens) {
            transferTokens(msg.sender,myBalance.tokens);
            delete balances[msg.sender];
        }
        myBalance.lastPotPerToken = potPerToken;
        msg.sender.transfer(toPay);
    }

    function withdrawTokens() external {
        _withdrawRewards(true);
    }

    function balance(address account) public view returns (
        uint256    tokens,
        uint256    lastPotPerToken,
        uint256    accumulatedReward
    ) {
        Balance memory myBalance = balances[account];
        return (myBalance.tokens, myBalance.lastPotPerToken, myBalance.accumulatedReward);
    }

    function stats() public view returns (
        uint256 _potPerToken,
        uint256 _numTokens,
        uint256 _pot
    ) {
        return (potPerToken,numTokens,pot);
    }



}