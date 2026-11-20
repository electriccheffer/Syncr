package fsutil_test

import "errors"
import "testing"
import "os"
import "path/filepath"
import "syscall"
import "github.com/electriccheffer/Syncr/internal/fsutil"

func TestAbsolutePathRegularFileExists(t *testing.T){

	directory := t.TempDir()
	path := filepath.Join(directory,"regularfile")
	_,err := os.Create(path)	
	if err != nil{
		t.Errorf("unexpected error creating test file: %s",err.Error())
	}

	path,err = filepath.Abs(path)
	if err != nil {
		t.Errorf("unexpected error in absolute path: %s", err.Error())
	}

	result, err := fsutil.SyncableFile(path)
	if err != nil {
		t.Errorf("unexpected error returned: %s",err.Error())
	}
	if result != true{
		t.Error("expected true returned false")
	}
}

func TestAbsolutePathFileNoExist(t *testing.T){

	path := "noexist"
	result,err := fsutil.SyncableFile(path)
	if err == nil {
		t.Error("expected error no error thrown")
	}else{
		if !errors.Is(err,os.ErrNotExist){
			t.Errorf("error of incorrect type: %s",err.Error())
		}
	}
	if result {
		t.Error("expected false, got true")
	}
}

func TestAbsolutePathDirectoryCaseExists(t *testing.T){

	directory := t.TempDir()
	path,err := filepath.Abs(directory)
	if err != nil {
		t.Errorf("unexpected error creating absolute path:%s",err.Error())
	}
	
	result, err := fsutil.SyncableFile(path)
	if err != nil {
		t.Errorf("unexpected error returned: %s",err.Error())
	}
	if !result{
		t.Error("expected true, got false")	
	}
}

func TestAbsolutePathNonRegularFile(t *testing.T){

	directory := t.TempDir()
	path := filepath.Join(directory,"regularfile")
	
	err := syscall.Mkfifo(path,0o600)
	if err != nil {
		t.Errorf("unexpected error makeing fifo:%s", err.Error())
	}
	
	path,err = filepath.Abs(path)
	if err != nil {
		t.Errorf("unexpected error getting absolute path: %s",err.Error())
	}
	
	result, err := fsutil.SyncableFile(path)
	if err != nil {
		t.Errorf("unexpected error returned:%s",err.Error())
	}
	if result{
		t.Error("expected false got true")
	}
}
