#!/bin/bash

#perhaps it's better to do this manually
trash out.txt

numRuns=10

for i in $(seq 1 $numRuns)
do
    echo "RUN #$i of $numRuns"
    make run >> out.txt;
done

echo "-----------RESULTS--------------"
echo "player1: "$(grep "winner: player1" out.txt | wc -l)
echo "player2: "$(grep "winner: player2" out.txt | wc -l)
echo "TIE: "$(grep "TIE" out.txt | wc -l)
