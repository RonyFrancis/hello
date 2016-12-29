package yaml2

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Src Options
}
type Options struct {
	Iv  string
	Key string
}

func Yaml(loc string) Config {
	fmt.Println(loc)
	file, _ := filepath.Abs(loc)
	yamlFile, _ := ioutil.ReadFile(file)
	fmt.Println(string(yamlFile))

	var config Config
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	//fmt.Printf(config.Src)
	//a := config.Src

	fmt.Println("+++++++++++++++++++++++++++++++")
	return config
}
