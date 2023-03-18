#!/bin/bash

files=`ls *.md`
for i in $files; do
	rename="${i/md/org}"
	pandoc -s "$i" -o "$rename"
	echo "$i -> $rename"
done
