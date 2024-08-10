package domain

type EnvVars struct {
	MongoDBCollNameContentItems string `mapstructure:"MONGO_DB_COLL_NAME_CONTENT_ITEMS"`
	MongoDBName                 string `mapstructure:"MONGO_DB_NAME"`
	MongoDBURI                  string `mapstructure:"MONGO_DB_URI"`
	ServerPort                  string `mapstructure:"SERVER_PORT"`
}
