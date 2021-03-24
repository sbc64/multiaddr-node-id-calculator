with import <nixpkgs> { };

stdenv.mkDerivation {
  name = "go";
  buildInputs = [
    go_1_16
    gcc
    git
  ];
  shellHook = "
    export GOPATH=/home/sebas/.go
  ";
}
