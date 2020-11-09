# experiments in revenue distribution

The revenue contract is designed to share revenue (currently  in ETH) between tokenholders.

in its current format, the tokens have been abstracted so we have functions depositing and withdrawing but vin reality they do not transfer tokens.

## The test script

The test script is a CSV file (each line must have 3 elements)

### Fields

- user 
    - who is interacting with the contract
- command 
    - which function to call or script operation to run
- amount
    - a numerical amount to support the operation

### operations

- profit
    - called from banker, causes ETH to be sent to the contract for sharing
- deposit
    - user deposits <N> tokens into the contract
- wdprofit
    - user withdraws profit but not tokens
- wdtokens
    - user withdraws tokens and profit
- balance
    - get balance information about user account
- pending
    - find out how much profit a user could claim
- stats
    - get some general contrat statistics
- contractbal
    -get the balance of ETH in the staking contract
    
