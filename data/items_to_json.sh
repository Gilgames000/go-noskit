#!/bin/sh
# usage: items_to_json.sh <Items.dat_location>
{
printf '{"items":['
sed 's:\r:\n:g' "$@" | grep -E 'VNUM|INDEX' | perl -pe 's/\s+VNUM\s+(\d+)\s+\d+\n/{"vnum":\1,/' | perl -pe 's/\s+INDEX\s+(\d+).+\n/"\"inventory_type\":".($1%4)."},",/e' | sed 's:.$::'
printf ']}'
} > items.json
