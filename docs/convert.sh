#!/bin/bash

files=`ls *.org`
for i in $files; do
	rename="${i/org/md}"
	pandoc -s "$i" -o "$rename"
	echo "$i -> $rename"
done
