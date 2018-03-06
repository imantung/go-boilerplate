package testkit

import (
	"fmt"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

func FatalIfError(t *testing.T, err error) {
	if err != nil {
		fatal(t, err.Error(), 1)
	}
}

func FatalIfWrongError(t *testing.T, err error, message string) {
	if err.Error() != message {
		fatal(
			t,
			fmt.Sprintf("Wrong error message: %s", err.Error()),
			1,
		)
	}
}

func FatalIfWrongHttpCode(t *testing.T, rec *httptest.ResponseRecorder, code int) {
	FatalIf(
		t,
		rec.Code != code,
		fmt.Sprintf("wrong http code: %s", rec.Code),
	)
}

func FatalIf(t *testing.T, condition bool, message string) {
	if condition {
		fatal(t, message, 1)
	}
}

func fatal(t *testing.T, message string, funcLevel int) {
	_, file, no, ok := runtime.Caller(funcLevel + 1)
	if ok {
		simpleFileName := file[strings.LastIndex(file, "/")+1:]
		message = fmt.Sprintf("%s:%d: %s", simpleFileName, no, message)
	}

	t.Logf("%s\n", message)
	t.FailNow()
}
