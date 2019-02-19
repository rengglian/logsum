package logsum

import (
    "io/ioutil"
	"strings"
    "errors"
    "strconv"
    "log"
)

type configStruc struct {
    Version versionStruc
    LogFilePath string
    PS1Thresh float64
    PS5Thresh float64
    NrLines int64
}

type versionStruc struct{
    Major string
    Minor string
}

func Split(r rune) bool {
    return r == '\n' || r == '='
}

func Read(fileName string) (configStruc, error){
	
    var conf configStruc
    var err error
	/*open summary file*/
    content, err := ioutil.ReadFile(fileName)
    if err == nil {
        result := string(content)
        temp := strings.FieldsFunc(result,Split)
        if len(temp) != 11 {
            err = errors.New("Corrupt config file")
        } else {
            conf.Version.Major, conf.Version.Minor = versionParse(temp[2][0:len(temp[2])-1])
            conf.LogFilePath = temp[4][0:len(temp[4])-1]
            conf.PS1Thresh, err = strconv.ParseFloat(temp[6][0:len(temp[6])-1], 64)
            checkError("conv", err)
            conf.PS5Thresh, err = strconv.ParseFloat(temp[8][0:len(temp[8])-1], 64)
            checkError("conv", err)
            conf.NrLines, err = strconv.ParseInt(temp[10], 10, 64)
            checkError("conv", err)
        }
    }
	return conf, err
}

func versionParse(str string) (string, string){
    tmp := strings.SplitN(str, ".", 2)
    var major string
    var minor string
    if len(tmp)>1 {
        major = tmp[0]
        minor = tmp[1]
    } else {
        major = "0"
        minor = "0"
    }
    return major, minor
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}