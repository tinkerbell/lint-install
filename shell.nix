let _pkgs = import <nixpkgs> { };
in { pkgs ? import (_pkgs.fetchFromGitHub {
  owner = "NixOS";
  repo = "nixpkgs";
  #branch@date: nixpkgs-unstable@2022-08-17
  rev = "d61d4e71ba9a8f56e9f2092b7cfa9cffa4253971";
  sha256 = "0x76l64pchaqaw8v0b331g4jm3li1xcpkbwxpn9n6a769c4ynj58";
}) { } }:

with pkgs;

mkShell {
  buildInputs = [
    git
    gnumake
    go_1_18
    nixfmt
    nodePackages.prettier
    python3Packages.pip
    python3Packages.setuptools
    python3Packages.wheel
  ];
}
