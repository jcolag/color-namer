#!/bin/sh
tmpnames=$(mktemp)
tmpcodes=$(mktemp)
outfile=allcolors.csv
url=https://en.wikipedia.org/wiki/List_of_colors_%28compact%29

curl $url | grep "width:9em;padding:5px;margin:auto;" | sed 's/<[^>]*>//g' > "$tmpnames"
curl $url | grep "hsv(" | cut -f2 -d'"' | sed 's/,—,/,0,/g' > "$tmpcodes"
#echo "Name,Hue,Saturation,Value,Red,Green,Blue,RGB" > $outfile
paste -d';' "$tmpnames" "$tmpcodes" | sed 's/); /;/g;s/hsv(//g;s/rgb(//g;s/;/,/g;s/[°%]//g' >> $outfile

