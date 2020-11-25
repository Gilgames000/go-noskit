#!/bin/sh
# usage: items_to_csv.sh <Items.dat_location>
{
printf 'vnum,inventory_pocket\n'
sed 's:\r:\n:g' "$@" | grep -E 'VNUM|INDEX' | perl -pe 's/\s+VNUM\s+(\d+)\s+\d+\n/$1,/' | perl -pe 's/\s+INDEX\s+(\d+).+/$1%4/e'
} > items.csv
