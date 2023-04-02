package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *gorm.DB
}

type ServerConfig struct {
	Host string
	Port string
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseConnection(app.Db)
}

// func Server(serverUrl string) ServerConfig {
// 	var serverConfig ServerConfig
// 	prt, err := http.Get(serverUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := ioutil.ReadAll(prt.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sb := string(body)
// 	p := strings.TrimSuffix(sb, "\n")
// 	parser, err := gojq.NewStringQuery(p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// sName, _ := parser.QueryToString("data.[1].service_name")
// 	host, _ := parser.QueryToString("data.[2].service_host")
// 	sPort, _ := parser.QueryToString("data.[2].service_port")

// 	port := fmt.Sprintf(":%s", sPort)
// 	serverConfig.Host = host
// 	serverConfig.Port = port
// 	return serverConfig
// }

// func JwtService(jwtUrl string) ServerConfig {
// 	var serverConfig ServerConfig

// 	prt, err := http.Get(jwtUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := ioutil.ReadAll(prt.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sb := string(body)
// 	p := strings.TrimSuffix(sb, "\n")
// 	parser, err := gojq.NewStringQuery(p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// sName, _ := parser.QueryToString("data.[1].service_name")
// 	host, _ := parser.QueryToString("data.[1].service_host")
// 	sPort, _ := parser.QueryToString("data.[1].service_port")

// 	port := fmt.Sprintf(":%s", sPort)
// 	serverConfig.Host = host
// 	serverConfig.Port = port
// 	return serverConfig
// }
