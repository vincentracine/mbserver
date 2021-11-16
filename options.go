package mbserver

var Noop = func(Framer) {}

type ServerOptions struct {
	OnReadCoils       func(frame Framer)
	OnWriteSingleCoil func(frame Framer)
}

type ServerOption func(*ServerOptions)

// OnReadCoils will set a callback when ReadCoils is called
func OnReadCoils(value func(frame Framer)) ServerOption {
	return func(args *ServerOptions) {
		args.OnReadCoils = value
	}
}

// OnWriteSingleCoil will set a callback when WriteSingleCoil is called
func OnWriteSingleCoil(value func(frame Framer)) ServerOption {
	return func(args *ServerOptions) {
		args.OnWriteSingleCoil = value
	}
}
