package main 



import "fmt"



func callBackHelp(){
    
        fmt.Println("Welcome to the Pokedex!")
        fmt.Println("usage!:")
     avilablecommand := getCommand()
   for _,command := range avilablecommand {

        fmt.Printf(" %s: %s/n",command.name,command.description)
    }
} 
