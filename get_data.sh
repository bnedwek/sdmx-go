#!/bin/bash

URL="https://www.federalreserve.gov/datadownload/Output.aspx?rel=sloos&filetype=zip"

cd data
wget $URL -O 'sloos.zip'

unzip -o sloos.zip
rm sloos.zip