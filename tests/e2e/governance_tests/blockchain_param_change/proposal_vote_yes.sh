#!/bin/bash

# Submit the proposal
echo "Submitting the proposal"
echo ""
vid-noded tx gov submit-proposal param-change ./param-change-upgrade-yes-vote.json --from node1 --yes
sleep 7
echo ""
echo "Proposal is submitted"
echo ""

# Vote for the proposal
echo "Voting yes for the proposal"
echo ""
vid-noded tx gov vote 1 yes --from node1 --yes
sleep 7
echo ""
echo "Vote given"
echo ""

echo "Waiting for the voting period to end...."
echo ""
sleep 65

# Check if the slashing_downtime_fraction changes from 0.01 to 0.365
EXPECTED_SLASH_FRACTION_DOWNTIME='"'0.365000000000000000'"'
ACTUAL_SLASH_FRACTION_DOWNTIME=$(vid-noded q slashing params --output json | jq '.slash_fraction_downtime')

if [ $ACTUAL_SLASH_FRACTION_DOWNTIME != $EXPECTED_SLASH_FRACTION_DOWNTIME ]; then
    echo "Slash Fraction Downtime did not update to ${EXPECTED_SLASH_FRACTION_DOWNTIME}, got ${ACTUAL_SLASH_FRACTION_DOWNTIME}"
    exit 1
else
    echo "Updated Slash Fraction Downtime: ${ACTUAL_SLASH_FRACTION_DOWNTIME}"
fi
