package fsutil

import "os"

func SyncableFile(path string)(bool,error){
	
	fileInformation, err := os.Stat(path)		
	if err != nil {
		// TODO: log error 
		return false,err
	}
	
	isRegular := fileInformation.Mode().IsRegular()
	if isRegular {
		return isRegular,nil	
	}
	isDirectory := fileInformation.Mode().IsDir()
	if isDirectory{
		return isDirectory,nil
	}
			
	return false,nil
}
