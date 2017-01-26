package gendriver

var (
	// Registry is where each Driver is stored.
	//
	// Each Driver is meant to create a single Go source code file.
	Registry Registrar
)

func init() {
	Registry = new(internalRegistrar)
}
