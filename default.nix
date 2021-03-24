{ pkgs ? import <nixpkgs> {} }:
  pkgs.buildGoPackage rec {
    name = "node-id";
    goPackagePath = "node-id";
    version = "1.9.25";
    goDeps = ./deps.nix;
    src = ./.;
    doCheck = false;

    buildPhase = ''
      runHook preBuild
       (
         cd go/src/${goPackagePath}
         go build -o node-id main.go
       )
       runHook postBuild
    '';
    installPhase = ''
      runHook preInstall
      install -D go/src/${goPackagePath}/node-id $out/bin/node-id
      runHook postInstall
    '';
}
