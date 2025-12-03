#!/bin/bash

for y in `seq -w 2025 2025`; do
  mkdir $y

  for i in `seq -w 1 12`; do
    mkdir $y/day-$i
    mkdir $y/day-$i/solutions
    touch $y/day-$i/solutions/part1.go
    touch $y/day-$i/solutions/part2.go
    touch $y/day-$i/main.go
    touch $y/day-$i/README.md
    touch $y/day-$i/input.txt
  done
done