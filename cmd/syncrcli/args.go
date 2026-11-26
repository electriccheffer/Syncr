package main

import "errors"
import "flag"
import "os"

type Command interface {
	Run()error
}

type AddCommand struct{
	FileName string
	Remote string 
}

func (ac *AddCommand) Run()error{
	return nil
}

func ParseArguments(arguments []string)(Command,error){
	mode := arguments[0]
 	switch mode {
		case "add":
			addFlags := flag.NewFlagSet("add",flag.ContinueOnError)
			addFlags.SetOutput(os.Stderr)
			fileName := addFlags.String("f","","filepath to sync")
			remote := addFlags.String("r","","remote machine ip")
			if err := addFlags.Parse(arguments[1:]); err != nil {
				return nil, err
			}
			if *fileName == ""{
				return nil,errors.New("missing filepath -f name")
			}
			if *remote == ""{
				return nil,errors.New("missing remote -r argument")
			}
			return &AddCommand{FileName:*fileName,Remote:*remote}, nil
		
		default: 
			return nil,errors.New("invalid arguments")		
	}
}
