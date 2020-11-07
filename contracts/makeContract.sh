#!/bin/bash
solc use 0.7.4
solc --combined-json abi,bin -o ./oy  revenue.sol  --overwrite --optimize
abigen --combined-json ./oy/combined.json --pkg contracts --out core.go
# abigen --combined-json ../bin/contracts/stakedRevenueToken.json --pkg contracts --out core.go