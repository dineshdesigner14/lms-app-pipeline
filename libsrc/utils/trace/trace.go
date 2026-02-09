package trace

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/libsrc/utils/dtutil"
	"os"
	"sync"
)

var processname string
var tracedate string
var tracefile string
var tracefilepath string
var tracepath string
var traceptr *os.File
var writerbuf *bufio.Writer

var lgmutex = &sync.Mutex{}
var traceLevel = debugdef.DEBUG_LEVEL_SECURED

func SetTrace(TracePath string) {
	tracepath = TracePath
}

func SetTraceLevel(tLevel int) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	traceLevel = tLevel
}

func GetTraceLevel() int {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	return traceLevel
}

func OpenTrace(TraceFile string, ProcessName string) int {
	var tracedir string
	processname = ProcessName
	tracefile = TraceFile

	if tracepath == "" {
		if !globaldef.IsAppBaseDirExists() {
			fmt.Printf("\nIsAppBaseDirExists() failed....OpenTrace() failed....\n")
			os.Exit(0)
		}
		tracedate = dtutil.GetDate("dd-mm-yyyy")
		tracedir = fmt.Sprintf("%s/%s/%s", globaldef.GetAppBaseDir(), "log", tracedate)
	} else {
		tracedate = dtutil.GetDate("dd-mm-yyyy")
		tracedir = fmt.Sprintf("%s/%s", tracepath, tracedate)
	}
	os.Mkdir(tracedir, 0777)
	tracefilepath = fmt.Sprintf("%s/%s.debug", tracedir, tracefile)
	var err error
	traceptr, err = os.OpenFile(tracefilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return -1
	}
	writerbuf = bufio.NewWriter(traceptr)
	return 1
}

func ModifyTrace() {
	var tracedir string
	var Ltracedate string
	Ltracedate = dtutil.GetDate("dd-mm-yyyy")
	if Ltracedate != tracedate {
		tracedate = dtutil.GetDate("dd-mm-yyyy")
		if tracepath == "" {
			tracedir = fmt.Sprintf("%s/%s/%s", globaldef.GetAppBaseDir(), "log", tracedate)
		} else {
			tracedir = fmt.Sprintf("%s/%s", tracepath, tracedate)
		}
		os.Mkdir(tracedir, 0777)
		tracefilepath = fmt.Sprintf("%s/%s.debug", tracedir, tracefile)
		var err error
		traceptr.Close()
		traceptr, err = os.OpenFile(tracefilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		writerbuf = bufio.NewWriter(traceptr)
	}

}

func Log(tLevel int, format string, a ...interface{}) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	if tLevel <= traceLevel {
		ModifyTrace()
		Prefix := fmt.Sprintf("(%s)(%s)(%s)(%d)", dtutil.GetDate("dd/mm/yyyy"), dtutil.GetTime("hh:mm:ss:ms"), processname, os.Getpid())
		fmt.Fprintf(writerbuf, "%s", Prefix)
		fmt.Fprintf(writerbuf, format, a...)
		fmt.Fprintf(writerbuf, "\n")
		writerbuf.Flush()
	}
}

func LgReq(ReqNum string, format string, a ...interface{}) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	ModifyTrace()
	Prefix := fmt.Sprintf("(%s)(%s)(%s)(%s)(%d)", dtutil.GetDate("dd/mm/yyyy"), dtutil.GetTime("hh:mm:ss:ms"), ReqNum, processname, os.Getpid())
	fmt.Fprintf(writerbuf, "%s", Prefix)
	fmt.Fprintf(writerbuf, format, a...)
	fmt.Fprintf(writerbuf, "\n")
	writerbuf.Flush()
}

func LogHex(tLevel int, dumpdata []byte) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	if tLevel <= traceLevel {
		fmt.Fprintf(writerbuf, "%s", hex.Dump(dumpdata))
		fmt.Fprintf(writerbuf, "\n")
		writerbuf.Flush()
	}
}

func LogDump(tLevel int, format string, a ...interface{}) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	if tLevel <= traceLevel {
		ModifyTrace()
		fmt.Fprintf(writerbuf, format, a...)
		fmt.Fprintf(writerbuf, "\n")
		writerbuf.Flush()
	}
}

func Lg(format string, a ...interface{}) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	ModifyTrace()
	Prefix := fmt.Sprintf("(%s)(%s)(%s)(%d)", dtutil.GetDate("dd/mm/yyyy"), dtutil.GetTime("hh:mm:ss:ms"), processname, os.Getpid())
	fmt.Fprintf(writerbuf, "%s", Prefix)
	fmt.Fprintf(writerbuf, format, a...)
	fmt.Fprintf(writerbuf, "\n")
	writerbuf.Flush()
}

func LgHex(dumpdata []byte) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	fmt.Fprintf(writerbuf, "%s", hex.Dump(dumpdata))
	fmt.Fprintf(writerbuf, "\n")
	writerbuf.Flush()
}

func LgDump(format string, a ...interface{}) {
	lgmutex.Lock()
	defer lgmutex.Unlock()
	ModifyTrace()
	fmt.Fprintf(writerbuf, format, a...)
	fmt.Fprintf(writerbuf, "\n")
	writerbuf.Flush()
}
