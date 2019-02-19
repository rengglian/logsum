package logsum
	
import (
	/*nothing to import*/
)
/*list of all files, last entry > only new files*/
func Files(fileList []string, lastEntry string) ([]string){
	
	var filesToRead []string

	/*look for last record in file list and create a list with only newest files*/
	for i, str:= range fileList {
		if str == lastEntry {
			if i+1 == len(fileList) {
				filesToRead = nil
			} else {
				filesToRead = fileList[i+1:len(fileList)]
			}
			break
		} else {
			filesToRead = fileList[1:len(fileList)]
		}
	}
		
	return filesToRead
	
}