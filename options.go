package mbserver

type FunctionCallback func(code FuncCode, frame Framer)

var Noop = func(FuncCode, Framer) {}

type ServerOptions struct {
	OnFunction FunctionCallback
}

type ServerOption func(*ServerOptions)

// OnFunction will set a callback when a function is called
func OnFunction(value FunctionCallback) ServerOption {
	return func(args *ServerOptions) {
		args.OnFunction = value
	}
}
