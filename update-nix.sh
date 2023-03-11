#!/usr/bin/env bash

sed -e "s/__VERSION__/$(git describe --tags --abbrev=0)/g" dynamotableviz.nix.tmpl > dynamotableviz.nix

# We first try to build and it fails with hash mismatch, and we use it to populate sha256.
nix-build -E 'with import <nixpkgs> { }; callPackage ./dynamotableviz.nix { }'
SRC_SHA256="$(nix-build -E 'with import <nixpkgs> { }; callPackage ./dynamotableviz.nix { }' 2>&1 | grep -oE 'got:\s+sha256-\S+' | cut -d "-" -f 2)"
sed -i -e "s|sha256 = lib.fakeSha256;|sha256 = \"$SRC_SHA256\";|g" dynamotableviz.nix

# We try again to build and it fails with hash mismatch, and we use it to populate vendorSha256.
VENDOR_SHA256="$(nix-build -E 'with import <nixpkgs> { }; callPackage ./dynamotableviz.nix { }' 2>&1 | grep -oE 'got:\s+sha256-\S+' | cut -d "-" -f 2)"
sed -i -e "s|vendorSha256 = lib.fakeSha256;|vendorSha256 = \"$VENDOR_SHA256\";|g" dynamotableviz.nix

