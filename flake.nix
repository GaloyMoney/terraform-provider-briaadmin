{
  description = "terraform-provider-bria";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        buildDeps = with pkgs; [
          go
          gotools
          gnumake
        ];
        devDeps = with pkgs;
        buildDeps ++ [
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc
          terraform
        ];
      in
      rec {
        packages = {
          terraform-provider-bria = pkgs.buildGoModule rec {
            pname = "terraform-provider-bria";
            version = "0.1.0";

            src = ./.;

            vendorSha256 = null;

            subPackages = [ "admin-provider" ];

            buildPhase = ''
            go build -o $out/bin/${pname} ./admin-provider
            '';

            checkPhase = ''
            go test ./... -v
            '';
          };
        };

        defaultPackage = packages.terraform-provider-bria;

        devShell = pkgs.mkShell {
          buildInputs = devDeps;

          shellHook = ''
          export PATH="$PATH:${pkgs.gnumake}/bin"
          alias ll="ls -al"
          alias tf=terraform
          '';
        };
      });
    }

