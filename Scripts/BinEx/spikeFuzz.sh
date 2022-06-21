#!/bin/bash -e

directory=$(pwd)
files="./*.spk"

if [ $# -lt 4 ]; then
	#If fewer than 4 parameters are given, then display an error and exit
	echo "Usage: spikeFuzz IP PORT SKIPVAR SKIPSTRnn."
	exit
else
	echo "Starting fuzz with spk files in " $directory
fi

for f in $files
do
	spike-fuzzer-generic-send_tcp $1 $2 $f 0 0
done
