package mbserver

type FunctionCallback func(frame Framer)

var Noop = func(frame Framer) {}

type ServerOptions struct {
	OnReadCoils             FunctionCallback
	OnReadDiscreteInputs    FunctionCallback
	OnReadHoldingRegisters  FunctionCallback
	OnReadInputRegisters    FunctionCallback
	OnWriteSingleCoil       FunctionCallback
	OnWriteHoldingRegister  FunctionCallback
	OnWriteMultipleCoils    FunctionCallback
	OnWriteHoldingRegisters FunctionCallback
}

type ServerOption func(*ServerOptions)

// OnReadCoils will set a callback when ReadCoils is called
func OnReadCoils(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnReadCoils = value
	}
}

// OnReadDiscreteInputs will set a callback when ReadDiscreteInputs is called
func OnReadDiscreteInputs(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnReadDiscreteInputs = value
	}
}

// OnReadHoldingRegisters will set a callback when ReadHoldingRegisters is called
func OnReadHoldingRegisters(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnReadHoldingRegisters = value
	}
}

// OnReadInputRegisters will set a callback when ReadInputRegisters is called
func OnReadInputRegisters(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnReadInputRegisters = value
	}
}

// OnWriteSingleCoil will set a callback when WriteSingleCoil is called
func OnWriteSingleCoil(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnWriteSingleCoil = value
	}
}

// OnWriteHoldingRegister will set a callback when WriteHoldingRegister is called
func OnWriteHoldingRegister(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnWriteHoldingRegister = value
	}
}

// OnWriteMultipleCoils will set a callback when WriteMultipleCoils is called
func OnWriteMultipleCoils(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnWriteMultipleCoils = value
	}
}

// OnWriteHoldingRegisters will set a callback when WriteHoldingRegisters is called
func OnWriteHoldingRegisters(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnWriteHoldingRegisters = value
	}
}
