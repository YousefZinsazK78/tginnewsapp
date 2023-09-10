package utils

import (
	"log"
)

var ErrorLogger *log.Logger
var GeneralLogger *log.Logger

// func init() {
// 	logFile, err := os.OpenFile("./log/general-log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("error opening log file")
// 	}
// 	// defer logFile.Close()
// 	GeneralLogger = log.New(logFile, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	ErrorLogger = log.New(logFile, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
// }
