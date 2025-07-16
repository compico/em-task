package config

import "time"

type HttpServer interface {
	GetAddr() string
	GetReadTimeout() time.Duration
	GetReadHeaderTimeout() time.Duration
	GetWriteTimeout() time.Duration
	GetIdleTimeout() time.Duration
	GetMaxHeaderBytes() int
	GetDisableGeneralOptionsHandler() bool
}

type httpServer struct {
	Addr                         string        `yaml:"addr"`
	ReadTimeout                  time.Duration `yaml:"read_timeout"`
	ReadHeaderTimeout            time.Duration `yaml:"read_header_timeout"`
	WriteTimeout                 time.Duration `yaml:"write_timeout"`
	IdleTimeout                  time.Duration `yaml:"idle_timeout"`
	MaxHeaderBytes               int           `yaml:"max_header_bytes"`
	DisableGeneralOptionsHandler bool          `yaml:"disable_general_options_handler"`
}

func (s *httpServer) GetAddr() string {
	return s.Addr
}

func (s *httpServer) GetReadTimeout() time.Duration {
	return s.ReadTimeout
}

func (s *httpServer) GetReadHeaderTimeout() time.Duration {
	return s.ReadHeaderTimeout
}

func (s *httpServer) GetWriteTimeout() time.Duration {
	return s.WriteTimeout
}

func (s *httpServer) GetIdleTimeout() time.Duration {
	return s.IdleTimeout
}

func (s *httpServer) GetMaxHeaderBytes() int {
	return s.MaxHeaderBytes
}

func (s *httpServer) GetDisableGeneralOptionsHandler() bool {
	return s.DisableGeneralOptionsHandler
}
