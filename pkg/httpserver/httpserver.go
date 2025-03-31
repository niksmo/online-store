package httpserver

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const shutdownTimeout = 5 * time.Second

type HTTPServer struct {
	addr     string
	FiberApp *fiber.App
}

func New(addr string) HTTPServer {
	return HTTPServer{addr: addr, FiberApp: fiber.New()}
}

func (s HTTPServer) Listen(errCb func(err error)) {
	err := s.FiberApp.Listen(s.addr)
	errCb(err)
}

func (s HTTPServer) Close() error {
	return s.FiberApp.ShutdownWithTimeout(shutdownTimeout)
}
