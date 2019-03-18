package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type config struct {
    key string `yaml:"key"`
    server string `yaml:"server"`
    webservice string `yaml:"webservice"`
    password string `yaml:"password"`
}

func (c *config) getConf() *config {

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
    var c config
    c.getConf()

    fmt.Println(c)
}
