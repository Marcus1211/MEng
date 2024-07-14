#!/bin/bash

cd data_graph

for FILE in ./*.txt; do
  go run txt-to-json.go $FILE
done

rm -rf *.txt
