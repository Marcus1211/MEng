#!/bin/bash

cd data_graph

### Download all test data files ###
curl -O https://snap.stanford.edu/data/soc-pokec-relationships.txt.gz
curl -O https://snap.stanford.edu/data/twitch.zip
curl -O https://snap.stanford.edu/data/facebook_combined.txt.gz
curl -O https://snap.stanford.edu/data/git_web_ml.zip
curl -O https://snap.stanford.edu/data/soc-LiveJournal1.txt.gz
curl -O https://snap.stanford.edu/data/email-Enron.txt.gz
curl -O https://snap.stanford.edu/data/email-EuAll.txt.gz
curl -O https://snap.stanford.edu/data/p2p-Gnutella31.txt.gz
curl -O https://snap.stanford.edu/data/bigdata/communities/com-lj.ungraph.txt.gz
curl -O https://snap.stanford.edu/data/bigdata/communities/com-amazon.ungraph.txt.gz
curl -O https://snap.stanford.edu/data/web-Stanford.txt.gz
curl -O https://snap.stanford.edu/data/web-Google.txt.gz
curl -O https://snap.stanford.edu/data/amazon0505.txt.gz
curl -O https://snap.stanford.edu/data/soc-Slashdot0811.txt.gz

### unzip all compressed files ###
find . -name "*.gz" -exec gunzip {} \;
unzip twitch.zip
mv ./twitch/PTBR/musae_PTBR_features.json ./
unzip git_web_ml.zip
mv ./git_web_ml/musae_git_features.json ./
rm -rf twitch git_web_ml

### remove all lines that start with '#' ###
for FILE in ./*.txt; do
  sed -i.bak '/^#/d' ./"$FILE"
done

rm -rf *.bak
