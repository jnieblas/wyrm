# wyrm
CLI tool that simplifies command execution & management.

# Setup
From root dir:
1. Retrieve the go install path: `go list -f ‘{{.Target}}’`.
2. Add the go install path to your system's shell path: `export PATH=$PATH:/yourPath/go/bin`.
3. Run `go build` to generate the wyrm binary.
4. Run `go install` to make wyrm globally accessible.
5. Run `wyrm -provision` to instantiate wyrm's database.

# Commands
`-i [-name=script_name]` - Print info about one or many scripts. No parameter will display info for all scripts.

`-c` - Create a script.

`-u` - Update an existing script.

`wyrm <script_name>` - Invoke a script.

`-l` - Enable logging
