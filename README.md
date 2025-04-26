# GODO
godo is a modern take on Make with some needed QoL changes.

## Features
- **where** allows you to run a command from any child directory.
- **variants** os indipendent commands.
- **description** tell exactly what the build command does.
- **fallback** if you don't rember what build commands you have, simply run godo without a command and godo will tell you what commands you have.


## Syntax

```yaml
commands:
  test:
    run: 
      - go clean -testcache
      - go test ./... -v
    description: Runs all the test
  os: # Variants allow os indipendent commands to be ran the platform is based on GOOS. 
      # Here is a list of possible GOOS: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
    variants:
      - run: echo Windows
        platform: windows
      - run: echo Linux
        platform: linux
      - run: echo Unkown
        platform: defualt
    description: Tells you what os you are running


# A command can have these values
# run - one or more commands that will run in the given order
# where - enter the path from the current folder where the commands will be ran at
# type - how godo will run the command: raw, path or shell. But for most usecases this can be ignored
# description
# variants - allows for better control over how commands are ran in different enviroments 
#     variants can have these values
#       run - the same as commands run (if the command has run commands the variants will be ignored)
#       platform - can be any value that is in GOOS or defualt
#       type - the same as commands type
```

## Installation

Use `go install` to grab the latest build:

```bash
# Latest 
go install github.com/VincentBrodin/godo@latest
```

## Tips

You can run godo commands from inside of godo, so if you have big complex commands for your build system you can split them up into there own smaller godo commands.


### Contribute

If you like the idea of godo and what to help improve it please do. Any improvements are welcomed.


