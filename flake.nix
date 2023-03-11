{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    flake-utils.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let 
        pkgs = import nixpkgs {
          system = system;
        };
        dynamotableviz = pkgs.callPackage ./dynamotableviz.nix {};
      in
      {
        defaultPackage = dynamotableviz;
        packages = { 
          dynamotableviz = dynamotableviz;
        };
        devShells = {
          default = pkgs.mkShell {
            packages = [ dynamotableviz ];
          };
          dynamotableviz = pkgs.mkShell {
            packages = [ dynamotableviz ];
          };
        };
      }
    );
}
