#! /bin/bash

deburl="https://dl.ui.com/unifi/$1/unifi_sysvinit_all.deb"
wkdir="$(mktemp -d)"
deb="$wkdir\unifi.deb"

curl -o "$deb" "$deburl"

mkdir -p "$wkdir/unifi"
dpkg-deb -R "$deb" "$wkdir/unifi"

cp "$wkdir/unifi/usr/lib/unifi/lib/ace.jar" ./
unzip -o ace.jar -d ./ace/

mkdir -p "$1"

cp ./ace/api/fields/*.json "./$1/"

./settings.sh "$1"

rm -rf ace ace.jar

go run main.go "$1" "../unifi/"

gofmt -w -s ./../unifi/