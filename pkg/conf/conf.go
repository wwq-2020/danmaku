package conf

// Conf Conf
type Conf struct	{
	Server *Server `toml:"server"`
}

// Server Server
type Server struct{
	Addr string `toml:"addr"`
}