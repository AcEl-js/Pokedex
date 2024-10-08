package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)


type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config,mapArea string,name string)error 
}

// mangen the command response 

func getCommand() map[string]cliCommand { 
return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    callBackHelp,
    },
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    callBackExit,
    },
    "map": {
        name:        "map",
        description: "Get Pokemon map",
        callback:    callBackMap,
    },
    "mapp": {
        name:        "mapp",
        description: "go to privous map location",
        callback:    callBackMapPrivous,
    },
    "explore": {
        name:        "explore",
        description: "exolore existing pokemons in a location Area",
        callback:    callBackExplore,
    },
    "catch": {
        name:        "catch",
        description: "catche a pokemon",
        callback:    callBackCatch,
    },
}
}
    
// handel the command format

func cleanCli(command string)[]string{
    
    lowerCase := strings.ToLower(command)
    words := strings.Fields(lowerCase)
    return words
}

// creat a REPL program 

func startRepl (cfg *config) {

    scanner := bufio.NewScanner(os.Stdin)
    for {
    fmt.Print("pokedex >")
        scanner.Scan()
        cleaned := cleanCli(scanner.Text())
       if (len(cleaned) == 0){
            continue
        }

        commandName := cleaned[0]
        availableCommand := getCommand()
        command, ok := availableCommand[commandName]
        if !ok {
            fmt.Println("invalid command (:")
        continue
        }


        if commandName == "explore" && len(cleaned)==1{
                fmt.Println("please enter the location area")

        }
        if len(cleaned) ==2 {
            switch commandName{
            case "explore":
            err := command.callback(cfg,cleaned[1],"")
            if err != nil{
            fmt.Println(err)
            }
            continue
             

            case "catch":
                err := command.callback(cfg,"", cleaned[1])
            if err != nil{
                    fmt.Println(err)
                }
            continue
            }

        } 
        err := command.callback(cfg,"","")
        if err != nil{
            fmt.Errorf("failed to call map",err)
        }




        }
    }
