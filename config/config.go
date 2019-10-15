package config

type Type struct {
	LogAllMethods   bool
	TraceAllMethods bool
}

var Config = createConfig()

func createConfig() Type {
	return Type{
		LogAllMethods: true,
	}
}
