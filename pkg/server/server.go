package server

import "net/http"

type Server struct {
}

func New() *Server {
	return &Server{}
}

var page = `<!DOCTYPE html>
<html>
	<body>
		<h1>This is the header</h1>
		<p>This is the body</p>
		<p>This is a second paragraph</p>
	</body>
</html>`

var user = `{
	"name": "jay rexiti",
	"age": 28
}`

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("my-header", "this is my string B")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}

func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("my-header", "this is my string B")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}
