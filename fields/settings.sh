#!/usr/bin/env bash

ver="$1"
keys=$(jq -r keys[] "$ver/Setting.json")

while IFS= read -r key; do
    readarray -td ' ' arr <<< "${key//_/ }"
    fn=$(printf %s "${arr[@]^}")
    echo "... $key $fn ..."
    jq ".$key" "$ver/Setting.json" > "$ver/Setting$fn.json"
done <<< "$keys"
