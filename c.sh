#!/bin/bash
cd contracts
~/go/bin/abigen --sol revenue.sol --pkg contracts --out core.go
cd ..
