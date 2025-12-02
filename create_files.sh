#!/bin/bash

for y in `seq -w 2025 2025`; do
  mkdir $y

  for i in `seq -w 1 12`; do
    mkdir $y/day-$i
    mkdir $y/day-$i/part-1
    mkdir $y/day-$i/part-2
    touch $y/day-$i/part-1/main.go
    touch $y/day-$i/part-2/main.go
    touch $y/day-$i/README.md
    touch $y/day-$i/input.txt
  done
done