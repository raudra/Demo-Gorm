package config

import(
	"gopkg.in/yaml.v2"
    "io/ioutil"
)


type Conf struct{
	Development map[string]interface{}
	Test map[string]interface{}
	Production map[string]interface{}
}

func (c *Conf)GetConf(filePath string) *Conf{
	yamlFile, err := ioutil.ReadFile(filePath)
    RaiseException(err)
    err = yaml.Unmarshal(yamlFile, c)
    RaiseException(err)
    return c
}

func (c *Conf)FetchConfig(filePath string) map[string]interface{}{
	c.GetConf(filePath)
	var config map[string]interface{}
	if IsGoenvDevelopment(){
		config = c.Development
	}else if IsGoenvProduction(){
		config = c.Production
	}else if IsGoenvTest(){
		config = c.Test
	}
	return config
}