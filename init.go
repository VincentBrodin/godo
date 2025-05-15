package main
					
const Init = 
`commands:

  loop:
    run: echo "Hello world"
    times: 10
    description: "Runs command 10 times"

  fail:
    run: 
      - a
      - $b
      - echo Works
    description: This code will fail twice

  os:    
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
# times - how many times the command will be ran
# description
# variants - allows for better control over how commands are ran in different enviroments 
#     variants can have these values
#       run - the same as commands run (if the command has run commands the variants will be ignored)
#       platform - can be any value that is in GOOS or defualt`
