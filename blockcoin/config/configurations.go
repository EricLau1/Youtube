package config

type Database struct {
  Name string
  User string
  Pass string
}

type Jwt struct {
  SecretKey []byte
}

type Config struct {
  Database Database
  Jwt Jwt
}

func LoadConfigs() Config {
  return Config{Database{"blockcoin", "postgres", "@root"},Jwt{[]byte("mindawakebodyasleep")}}
}
