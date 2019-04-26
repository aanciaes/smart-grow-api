package handlers

import (
	"io"
	"net"
	"net/http"
	"time"
)

func writeLogLine (out io.Writer, req *http.Request) {
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
	buf = append(buf, `]`...)
	buf = append(buf, req.Method...)
	buf = append(buf, " "...)
	buf = append(buf, uri...)
	buf = append(buf, " "...)
	buf = append(buf, req.Proto...)
	buf = append(buf, ` "`...)
	buf = append(buf, req.Referer()...)
	buf = append(buf, `"`...)
	buf = append(buf, req.UserAgent()...)
	buf = append(buf, '"', '\n')
	out.Write(buf)
}


func LoggingWrapper (out io.Writer, h http.Handler) http.Handler {


	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		writeLogLine(out, r)
		h.ServeHTTP(w, r)
	})
}