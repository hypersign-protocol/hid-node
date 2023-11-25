#!/bin/bash

echo "================================================"
echo "   Setup & Starting Hypersign 3 Node Network    "
echo "================================================"

./setup.sh
./start.sh 
./fund.sh 
./createValidators.sh

echo "================================================"
echo "      Hypersign 3 Node Network Setup End        " 
echo "================================================"




