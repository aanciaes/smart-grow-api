package handlers

import (
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

type LogRecorder struct {
	http.ResponseWriter
	status int
}

func (r *LogRecorder) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (rec *LogRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func writeLogLine (out io.Writer, req *http.Request, status int) {
	host, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		host = req.RemoteAddr
	}

	uri := req.RequestURI
	t := time.Now()

	buf := make([]byte, 0, )
	buf = append(buf, host...)
	buf = append(buf, " - "...)
	buf = append(buf, `[ "`...)
	buf = append(buf, t.Format("02/Jan/2006:15:04:05 -0700")...)
	buf = append(buf, `] `...)
	buf = append(buf, req.Method...)
	buf = append(buf, " "...)
	buf = append(buf, uri...)
	buf = append(buf, " "...)
	buf = append(buf, req.Proto...)
	buf = append(buf, `" `...)
	buf = append(buf, strconv.Itoa(status)...)
	buf = append(buf, " "...)
	buf = append(buf, ` "`...)
	buf = append(buf, req.Referer()...)
	buf = append(buf, `"`...)
	buf = append(buf, req.UserAgent()...)
	buf = append(buf, '"', '\n')
	out.Write(buf)
}

func LoggingWrapper (out io.Writer, h http.Handler) http.Handler {


	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		rec := LogRecorder{w, 200}

		h.ServeHTTP(&rec, r)

		writeLogLine(out, r, rec.status)
	})
}