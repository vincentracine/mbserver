// Package mbserver implments a Modbus server (slave).
package mbserver

import (
	"io"
	"net"

	"github.com/goburrow/serial"
)

// Server is a Modbus slave with allocated memory for discrete inputs, coils, etc.
type Server struct {
	// Debug enables more verbose messaging.
	Debug            bool
	listeners        []net.Listener
	ports            []serial.Port
	requestChan      chan *Request
	function         [256]func(*Server, Framer) ([]byte, *Exception)
	DiscreteInputs   []byte
	Coils            []byte
	HoldingRegisters []uint16
	InputRegisters   []uint16
	options          ServerOptions
}

// Request contains the connection and Modbus frame.
type Request struct {
	conn  io.ReadWriteCloser
	frame Framer
}

// NewServer creates a new Modbus server (slave).
func NewServer(setOptions ...ServerOption) *Server {
	s := &Server{
		options: ServerOptions{
			OnFunction: Noop,
		},
	}
	for _, setOption := range setOptions {
		setOption(&s.options)
	}

	// Allocate Modbus memory maps.
	s.DiscreteInputs = make([]byte, 65536)
	s.Coils = make([]byte, 65536)
	s.HoldingRegisters = make([]uint16, 65536)
	s.InputRegisters = make([]uint16, 65536)

	// Add default functions.
	s.function[FuncCodeReadCoils] = ReadCoils
	s.function[FuncCodeReadDiscreteInputs] = ReadDiscreteInputs
	s.function[FuncCodeReadHoldingRegisters] = ReadHoldingRegisters
	s.function[FuncCodeReadInputRegisters] = ReadInputRegisters
	s.function[FuncCodeWriteSingleCoil] = WriteSingleCoil
	s.function[FuncCodeWriteSingleRegister] = WriteHoldingRegister
	s.function[FuncCodeWriteMultipleCoils] = WriteMultipleCoils
	s.function[FuncCodeWriteMultipleRegisters] = WriteHoldingRegisters

	s.requestChan = make(chan *Request)
	go s.handler()

	return s
}

// RegisterFunctionHandler override the default behavior for a given Modbus function.
func (s *Server) RegisterFunctionHandler(funcCode uint8, function func(*Server, Framer) ([]byte, *Exception)) {
	s.function[funcCode] = function
}

func (s *Server) handle(request *Request) Framer {
	var exception *Exception
	var data []byte

	response := request.frame.Copy()

	function := request.frame.GetFunction()
	if s.function[function] != nil {
		data, exception = s.function[function](s, request.frame)
		response.SetData(data)
	} else {
		exception = &IllegalFunction
	}

	if exception != &Success {
		response.SetException(exception)
	}

	return response
}

// All requests are handled synchronously to prevent modbus memory corruption.
func (s *Server) handler() {
	for {
		request := <-s.requestChan
		response := s.handle(request)
		request.conn.Write(response.Bytes())
	}
}

// Close stops listening to TCP/IP ports and closes serial ports.
func (s *Server) Close() {
	for _, listen := range s.listeners {
		listen.Close()
	}
	for _, port := range s.ports {
		port.Close()
	}
}
