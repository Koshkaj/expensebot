package config

type Config struct {
	DbType    string
	StoreType string
	Server    ServerConfig
}

type MongoConfig struct {
	URI        string
	DB_NAME    string
	COLLECTION string
}

type MemoryConfig struct {
}

type ServerConfig struct {
	Port string
}

type GoogleProcessorConfig struct {
	Location    string
	ProjectID   string
	ProcessorID string
	Endpoint    string
}

type StoreConfig struct {
	DirectoryName string
}
