# wyrm
CLI tool that simplifies command execution & management.

# Setup
From root dir:
1. Retrieve the go install path: `go list -f ‘{{.Target}}’`.
2. Add the go install path to your system's shell path: `export PATH=$PATH:/yourPath/go/bin`.
3. Run `go build` to generate the wyrm binary.
4. Run `go install` to make wyrm globally accessible.
4. Run `wyrm -provision` to instantiate wyrm's database.
