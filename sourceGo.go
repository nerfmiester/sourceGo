package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nerfmiester/sourceGo/Config"

	"github.com/ivpusic/toml"
)

var name = flag.String("name", "World", "A name to say hello to")

var spanish bool

var useJSON bool

var useToml bool

var shootme bool

func checkHowIs(h Config.Hoogie) {
	fmt.Println(h)
	fmt.Println(h.Shootme())
}

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.") //#3
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
	flag.BoolVar(&useJSON, "j", false, "Use json language.")    //#3
	flag.BoolVar(&useToml, "t", false, "Use TOML language.")    //#3
	flag.BoolVar(&useJSON, "json", false, "Use json language.") //#3
	flag.BoolVar(&shootme, "shoot", false, "shoot me.")         //#3
}

func getName() string {
	return "World!"
}

func main() {
	fmt.Println("Hello", getName())

	flag.Parse()

	if shootme == true {

		things := []Config.Hoogie{Config.NewMawhat(), Config.NewDoohickie()}

		for _, thing := range things {

			checkHowIs(thing)

		}
		/*fmt.Println("I am ", m.state)
		m.shootme()
		fmt.Println("I am ", m.state)*/
	} else {

		if useToml == true {

			var tomlConfig Config.TomlConfig

			if _, err := toml.DecodeFile("definition.toml", &tomlConfig); err != nil {
				log.Fatal(err)
				// return
			}
			fmt.Printf("Title: %s\n", tomlConfig.Title)
			fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
				tomlConfig.Owner.Name, tomlConfig.Owner.Org, tomlConfig.Owner.Bio,
				tomlConfig.Owner.DOB)
			fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
				tomlConfig.DB.Server, tomlConfig.DB.Ports, tomlConfig.DB.ConnMax,
				tomlConfig.DB.ConnMax)
			for serverName, server := range tomlConfig.Servers {
				fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
			}
			fmt.Printf("Client data: %v\n", tomlConfig.Clients.Data)
			fmt.Printf("Client hosts: %v\n", tomlConfig.Clients.Hosts)

		}

		if useJSON == true {
			file, err1 := os.Open("config.json")
			if err1 != nil {
				fmt.Println("Error:", err1)
			} //#2
			defer file.Close()
			fmt.Println("Hello json")

			decoder := json.NewDecoder(file) //#3
			conf := Config.Cfg{}             //#3
			err := decoder.Decode(&conf)     //#3
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println(conf.Path)
		}
		if spanish == true {
			fmt.Printf("Hola %s!\n", *name)
		} else {
			fmt.Printf("Hello %s!\n", *name)
		}
	}
}
