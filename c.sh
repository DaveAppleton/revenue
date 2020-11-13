#!/bin/bash
cd contracts
~/go/bin/abigen --sol xStaking.sol --pkg contracts --out core.go
cd ..
