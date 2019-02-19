package logsum

import (
    "os"
	"encoding/csv"
	"strings"
)

func Check(fileName string) (string, error){
	
	/*open summary file*/

	var lastRecord string
	var file *os.File
	var records [][]string
	var err error	
	
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		if strings.HasPrefix(err.Error(), "FindFirstFile") { 
			err = nil 
		}
	} else {
		file, err = os.OpenFile(fileName, os.O_RDONLY, 0600)
		defer file.Close()
        /*read summary file*/
        r := csv.NewReader(file)
        r.Comma = ';'
        records, err = r.ReadAll()
		if err == nil {
			lastRecord = records[len(records)-1][len(records[0])-1];
		}else{
			if strings.HasSuffix(err.Error(), "wrong number of fields") { 
				err = nil 
			}
		}
	}
	return lastRecord, err
	
}