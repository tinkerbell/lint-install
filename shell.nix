let _pkgs = import <nixpkgs> { };
in { pkgs ? import (_pkgs.fetchFromGitHub {
  owner = "NixOS";
  repo = "nixpkgs";
  #branch@date: nixpkgs-unstable@2021-11-12
  rev = "2fbba4b4416446721ebfb2e0bfcc9e45d8ddb4d2";
  sha256 = "1yw2p38pdvx63n21g6pmn79xjpxa93p1qpcadrlmg0j0zjnxkwr8";
}) { } }:

with pkgs;

mkShell {
  buildInputs = [
    git
    gnumake
    go
    nixfmt
    nodePackages.prettier
    python3Packages.pip
    python3Packages.setuptools
    python3Packages.wheel
  ];
}
