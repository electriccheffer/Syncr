package main

import "testing"

func TestParseArgsSuccessAddFile(t *testing.T){

	args := []string{"add","-f","somefile.txt","-r","123.456.789.123"}
	command, err := ParseArguments(args)
	if err != nil{
		t.Errorf("unexpected error: %s", err.Error())
	}
	switch runCommand := command.(type) {

		case *AddCommand:
			break 
		default: 
			t.Errorf("incorrect type was returned expected: AddCommand got:%T",
				runCommand)
	}
	
}
