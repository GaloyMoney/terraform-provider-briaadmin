{
  description = "A simple Golang project";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
  let pkgs = import nixpkgs { inherit system; };
    in
    {
      devShell = pkgs.mkShell {
        buildInputs = [
          pkgs.go
        ];
         shellHook = ''
            export TF_CLI_CONFIG_FILE=./terraform.rc
          '';
      };
    });
}
