package utils

import (
  "os"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
  "path/filepath"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "encoding/json"
  "crypto/md5"
  "encoding/hex"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/keystore"
)

type SpecialDate struct {
    time.Time
}

func (sd *SpecialDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
    fmt.Println("SpecialDate.UnmarshalYAML: parse date")
    var input string
    if err := unmarshal(&input); err != nil {
        fmt.Println("SpecialDate.UnmarshalYAML: Cannot parse string")
        return err
    }
    strInput := strings.Trim(input, `"`)
    newTime, err := time.Parse("2006-01-01T00:00:00Z", strInput)
    if err != nil {
      fmt.Println("SpecialDate.UnmarshalYAML: parse time error: ",err)
        return err
    }
    sd.Time = newTime
    return nil
}

type Config struct {
    Version string `yaml:"version" json:"version"`
    Released SpecialDate `yaml:"released" json:"released"`
    Update bool `yaml:"update" json:"update"`
    Jwt struct {
        Enable bool `yaml:"enable" json:"enable"`
        Signkey string `yaml:"signkey" json:"signkey"`
        ExpiredAt int `yaml:"expiredAt" json:"expiredAt"`
        LoadAccount bool `yaml:"loadAccount" json:"loadAccount"`
        AccountFile string `yaml:"accountFile" json:"accountFile"`
    } `yaml:"jwt" json:"jwt"`
    Mysql struct {
        Host string `yaml:"host" json:"host"`
        Port string `yaml:"port" json:"port"`
        Username string `yaml:"username" json:"username"`
        Password string `yaml:"password" json:"password"`
        Database string `yaml:"database" json:"database"`
        Debug bool `yaml:"debug" json:"debug"`
    } `yaml:"mysql" json:"mysql"`
    Channel struct {
        TransferQueue int `yaml:"transferqueue" json:"transferqueue"`
        LogQueue int `yaml:"logqueue" json:"logqueue"`
    } `yaml:"channel" json:"channel"`
    Webserver struct {
			  Port string `yaml:"port" json:"port"`
        Tls bool `yaml:"tls" json:"tls"`
        CertificateFile string `yaml:"certificateFile" json:"certificateFile"`
        KeyFile string `yaml:"keyFile" json:"keyFile"`
        MaxRpcConnection int `yaml:"maxrpc" json:"maxrpc"`
        MaxListenRpcConnection int `yaml:"maxlistenrpc" json:"maxlistenrpc"`
        RoutingMode int `yaml:"routingMode" json:"routingMode"`
        NonceMode int `yaml:"nonceMode" json:"nonceMode"`
		} `yaml:"webserver" json:"webserver"`
		Keys  struct {
        LoadKey  bool `yaml:"loadKey" json:"loadKey"`
			  KeyStore string `yaml:"keystore" json:"keystore"`
				Password string `yaml:"password" json:"password"`
		} `yaml:"keys" json:"keys"`
		Networks []struct {
				Name string `yaml:"name" json:"name"`
				Http string `yaml:"http" json:"http"`
				WebSocket string `yaml:"websocket" json:"websocket"`
				LocalAddr string `yaml:"local" json:"local"`
		} `yaml:"networks" json:"networks"`
		Redis struct {
        MaxConn int  `yaml:"maxconn" json:"maxconn"`
			  Host string `yaml:"host" json:"host"`
			  Password string `yaml:"password" json:"password"`
			  Db int `yaml:"db" json:"db"`
		} `yaml:"redis" json:"redis"`
		Contract struct {
        GasPrice string `yaml:"gasprice" json:"gasprice"`
        GasLimit uint64 `yaml:"gaslimit" json:"gaslimit"`
				Owner string `yaml:"owner" json:"owner"`
				InitialToken int64 `yaml:"initialToken" json:"initialToken"`
				MasterKey1 string `yaml:"masterkey1" json:"masterkey1"`
				MasterKey2 string `yaml:"masterkey2" json:"masterkey2"`
				Address string `yaml:"address" json:"address"`
		} `yaml:"contract" json:"contract"`
    F5Contract struct {
        GasPrice string `yaml:"gasprice" json:"gasprice"`
        GasLimitDefault uint64 `yaml:"gaslimitdefault" json:"gaslimitdefault"`
        GasLimit map[string]uint64 `yaml:"gaslimit" json:"gaslimit"`
        EthBudget string `yaml:"ethBudget" json:"ethBudget"`
        Owner string `yaml:"owner" json:"owner"`
      	Address string `yaml:"address" json:"address"`
		} `yaml:"f5contract" json:"f5contract"`
}
func (cf *Config) toJson() string {
    fmt.Println("Convert Config struct to json: ")
    b, err := json.Marshal(cf)
    if err != nil {
        fmt.Println(err)
        return "{}"
    }
    return string(b)
}

type JwtAccount struct {
  gorm.Model
  Username string `json:"username"`
	Password string `json:"password"`
  Active bool `gorm:"default:false"`
}

type JwtAccountList struct {
  Accounts []struct {
    Username string `yaml:"username" json:"username"`
    Password string `yaml:"password" json:"password"`
  } `yaml:"accounts" json:"accounts"`
}

type Setting struct {
  gorm.Model
  Version string
  Config string  `sql:"type:text"`
  Description string  `sql:"type:text"`
  Active bool `gorm:"default:false"`
}


func (s *Setting) GetConfig() *Config {
  cf := &Config{}
  err := json.Unmarshal([]byte(s.Config), cf)
  if err != nil {
      fmt.Println("Cannot parse string to Config", err.Error())
      return nil
  }
  return cf
}

var cfg *Config

func (cf *Config) MysqlConnectionUrl() string {
    return cf.Mysql.Username + ":" + cf.Mysql.Password + "@tcp(" +  cf.Mysql.Host + ":" + cf.Mysql.Port + ")/" +   cf.Mysql.Database + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func (cf *Config) JwtVerifyUsername(username string,password string) bool {
    db, err := gorm.Open("mysql", cfg.MysqlConnectionUrl())
    if cfg.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()
    acc := &JwtAccount{}
    fmt.Println("Start search username: ")
    if err := db.Where("username = ?", username).Where("password = ?", password).Where("active = ?", true).First(acc).Error; err != nil {
      fmt.Println("username: ",username," not existed or password failed or not activate")
      return false
  }
   return true
}
func (cf *Config) LoadJwtAccounts()  {
    filePath := cf.Jwt.AccountFile
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        panic("File: "+ filePath + " not existed")
    }
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
           panic("Read file: "+ filePath + " error: " + err.Error())
    }
    accountList := &JwtAccountList{}
    err = yaml.Unmarshal(yamlFile, accountList)
    if err != nil {
        panic("Unmarshal: " +  err.Error())
    }

    //Create connection to mysql
    db, err := gorm.Open("mysql", cf.MysqlConnectionUrl())
    if cf.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&JwtAccount{})

    for _, account :=  range accountList.Accounts {
        acc := JwtAccount{}
       //Check db if not exist with username
       if err := db.Where("username = ?", account.Username).First(&acc).Error; err != nil {
           fmt.Println("Jwt Account: ",account.Username," not existed. Update new Jwt account: " + err.Error())
           user := JwtAccount{
             Username: account.Username,
             Password: GetMD5Hash(account.Password),
             Active: false,
           }
           //fmt.Println("Create new record")
           db.Create(&user)
       }
    }
}



func (cf *Config) LoadKeyStoresToDB(){

    root := cf.Keys.KeyStore
    fmt.Println("Start load accounts from keystores: ",root)

    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
        fmt.Println("Cannot find file in folder: ", root)
         panic(err)
    }

    fmt.Println("Connect to mysql: ",cf.MysqlConnectionUrl())
    //Create connect to db
    db, err := gorm.Open("mysql", cf.MysqlConnectionUrl())
    if cf.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
    db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&TokenAccount{})


    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) != 3 {
            fmt.Println("File name not correct format:", file)
            continue
         }
         account := list[2]

         token := &TokenAccount{}
        // find token with address
         if err := db.Where("address = ?", account).First(token).Error; err != nil {
             fmt.Println("Not find account with address: ", account," Start read file")

             //Load account in db
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }
             fmt.Println("Decrypt keystore of account: ", account)
             //Store account private key
             accountKey, err := keystore.DecryptKey( []byte(keyjson), cf.Keys.Password)
             if err != nil {
                 fmt.Println("Cannot decrypt key file: ", err)
                 continue
             }
             privateKey :=  hex.EncodeToString(crypto.FromECDSA(accountKey.PrivateKey))
             new_account := &TokenAccount{
               Address: account,
               PrivateKey: privateKey,
               Active: true,
             }
             //fmt.Println("Create new record")
             db.Create(new_account)
         }
    }
    fmt.Println("End load accounts from keystores: ",root)
}

func LoadConfig(file string) *Config {
    fileCfg := LoadConfigFromFile(file)

    //fmt.Println("Config: ",fileCfg.toJson())
    //Create mysql connection
    db, err := gorm.Open("mysql", fileCfg.MysqlConnectionUrl())
    if fileCfg.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Setting{})

    setting := &Setting{}
    // find setting with version
    if err := db.Where("version = ?", fileCfg.Version).First(setting).Error; err != nil {
        fmt.Println("version: ",fileCfg.Version," not existed. Update new version")
        setting := Setting{
          Version: fileCfg.Version,
          Config: string(fileCfg.toJson()),
          Active: false,
        }
        //fmt.Println("Create new record")
        db.Create(&setting)
    } else {
       if fileCfg.Update {
           //Update current record
          setting.Config = string(fileCfg.toJson())
          db.Save(&setting)
          fmt.Println("version: ",fileCfg.Version," update successfully")
       }
    }
    //Load First active setting
    fmt.Println("Load first active record")
    activeSetting := &Setting{}
    if err := db.Where("active = ?", true).First(activeSetting).Error; err != nil {
        fmt.Println("Cannot find active config. Load first config:")
        db.First(activeSetting)
    }

    if activeSetting != nil {
      fmt.Println("Parse json to Config struct")
      //Load config
      cfg = activeSetting.GetConfig()
      return cfg
    }

    panic("Cannot load config ")
    return nil
}

func LoadConfigFromFile(file string) *Config {
     var fcfg = &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err: ", err)
     }

     err = yaml.Unmarshal(yamlFile, fcfg)
     if err != nil {
         fmt.Println("Unmarshal error: ", err)
     }
     fmt.Println("Load config: ",fcfg.Version, ", Released: ", fcfg.Released)

     if fcfg.Jwt.LoadAccount {
       fcfg.LoadJwtAccounts()
     }

     if fcfg.Keys.LoadKey {
         fcfg.LoadKeyStoresToDB()
     }

     return fcfg
}

func LoadKey(root string,addr string) []byte {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
         panic(err)
    }
    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) == 3 {
					   account := list[2]
						 if account == strings.TrimPrefix(addr,"0x") {
							 keyjson, err := ioutil.ReadFile(file)
							 if err != nil {
										fmt.Println("Error in read file: ", file )
										return nil
							 }
							 return keyjson
						 }
         }
    }
		return nil
}
