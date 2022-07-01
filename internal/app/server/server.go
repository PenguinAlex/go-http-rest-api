package server

import "github.com/sirupsen/logrus"

//сервер ...
type Server struct {
	config *Config
	logger *logrus.Logger //либа для логгирования
}

//init server
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api-server")

	return nil
}

//настройка уровня логгирования
func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
