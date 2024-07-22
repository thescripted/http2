package http2

type HTTP2Server struct {
}

func NewServer() (HTTP2Server, error) {
	return HTTP2Server{}, nil
}

func (s *HTTP2Server) ListenAndServe(url string) error {
	return nil
}
