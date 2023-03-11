{ config, lib, pkgs, fetchFromGitHub, ... }:

pkgs.buildGoModule rec {
  pname = "dynamotableviz";
  version = "v0.0.13";
  src = pkgs.fetchFromGitHub {
    owner = "a-h";
    repo = "dynamotableviz";
    rev = version;
    sha256 = "8dT4winr0I1h5EnREcmG5CMyvbKOYkBiyc5AJ2hX2Ew=";
  };
  vendorSha256 = "BPsosNyuW0PBY3PZGXyNuiQ8pczJ6/9hcYnKtpQKzjo=";
}
