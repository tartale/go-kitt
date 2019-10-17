package config

type Type struct {
	FormatTool      string
	LogAllMethods   bool
	TraceAllMethods bool
}

var Config = createConfig()

func createConfig() Type {
	return Type{
		FormatTool:    "goimports",
		LogAllMethods: true,
	}
}
