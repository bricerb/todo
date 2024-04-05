package elog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	PANIC uint8 = 0
	ERROR uint8 = 1
)

var levels = [2]string{"FATAL: ", "ERROR: "}

func New(level uint8, str string, err error) {
	if err != nil {

		// Format
		err_str := fmt.Sprintf("%s %s -> %v", levels[level], str, err.Error())

		// Write
		mtx := &sync.Mutex{}
		writeLog(mtx, err_str)

		// PANIC
		if level == PANIC {
			panic(err_str)
		}

		// Console output
		log.Println(err_str)

	}
}

func writeLog(mtx *sync.Mutex, err_str string) {
	mtx.Lock()
	defer mtx.Unlock()

	// .log dir
	log_dir := "./logs/"
	err := os.MkdirAll(log_dir, 0755)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}

	// .log file
	strTime := time.Now().Format("01-January-2024")
	f_name := log_dir + strTime + ".log"

	// Create new file each day
	l_file, err := os.OpenFile(f_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		os.Stderr.WriteString(err.Error() + f_name + "\n")
		os.Exit(-1)
	}
	defer l_file.Close()

	// .log Timestamp
	strTime = time.Now().Format("10:11:12.00000")

	// Write to .log file
	fmt.Fprintf(l_file, strTime + " %s\n", err_str)

}