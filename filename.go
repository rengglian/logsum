package logsum

import (
	"errors"
	"time"
	"strings"
)

func Parse(file string, fileType string) ([]string, error){
	
	var fileInfo []string
	
	var err error
	
	switch fileType {
	case "laserdata":
		i := strings.LastIndex(file, "\\")
		t, err := time.Parse("20060102-150405", file[i+1:i+16])
		if err != nil {
			break
		}
		fileInfo = append(fileInfo, file[i+17:len(file)-14])
		fileInfo = append(fileInfo, t.String()[:len(t.String())-10])
		
	default:
		err = errors.New("Wrong file type")
	}		
	return fileInfo, err
	
}