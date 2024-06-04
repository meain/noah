{
  description = "Note taking helper";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system}; in
      {
        packages = rec {
          noah = pkgs.buildGoModule {
            pname = "noah";
            version = "dev";
            src = ./.;
	    vendorHash = null;
            doCheck = false;
          };

          default = noah;
        };

        devShells.default = pkgs.mkShell {
          # needed for dlv to work (https://github.com/NixOS/nixpkgs/issues/18995)
          hardeningDisable = [ "fortify" ];
          packages = with pkgs; [ go delve ];
        };
      }
    );
}
