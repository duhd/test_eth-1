package utils

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Config struct {
    Key string `json:"key"`
    Server string `yaml:"server"`
    Webservice string `yaml:"webservice"`
    Password string `yaml:"password"`
}

func (c *Config) GetConf() *Config {

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}
func main() {
    var c Config
    c.GetConf()

    fmt.Println(c)
}
