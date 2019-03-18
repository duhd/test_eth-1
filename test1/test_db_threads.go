package main

import (
    "github.com/jinzhu/gorm"
    // _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "sync"
    "time"
    "fmt"
    "runtime"
)

type Account struct {
  gorm.Model
  Code string
  Value uint
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)

  fmt.Printf("NumCPU: ", runtime.NumCPU())
  fmt.Println("GOMAXPROCS: ",runtime.GOMAXPROCS(runtime.NumCPU()))
  start := time.Now()
  db, err := gorm.Open("mysql", "admin:123456@/demo?charset=utf8&parseTime=True&loc=Local")
  // db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect database")
  }
  db.DB().SetMaxIdleConns(100)
  defer db.Close()

  db1, err1 := gorm.Open("mysql", "admin:123456@/demo?charset=utf8&parseTime=True&loc=Local")
  if err1 != nil {
    panic("Thread1: failed to connect database")
  }
  db1.DB().SetMaxIdleConns(100)
  defer db1.Close()

  db2, err2 := gorm.Open("mysql", "admin:123456@/demo?charset=utf8&parseTime=True&loc=Local")
  if err2 != nil {
    panic("Thread2: failed to connect database")
  }
  db2.DB().SetMaxIdleConns(100)
  defer db2.Close()

  // db3, err3 := gorm.Open("mysql", "admin:123456@/demo?charset=utf8&parseTime=True&loc=Local")
  // if err3 != nil {
  //   panic("Thread3: failed to connect database")
  // }
  // db3.DB().SetMaxIdleConns(100)
  // defer db3.Close()
  //
  // db4, err4 := gorm.Open("mysql", "admin:123456@/demo?charset=utf8&parseTime=True&loc=Local")
  // if err4 != nil {
  //   panic("Thread4: failed to connect database")
  // }
  // db4.DB().SetMaxIdleConns(100)
  // defer db4.Close()

  // Migrate the schema
  db.AutoMigrate(&Account{})
  var account Account
  db.First(&account) // find product with code l1212
  if &account != nil {
      db.Delete(&account)
  }
  // Create
  db.Create(&Account{Code: "eb80964e1567064ba810b45300fd2ce3193d1684", Value: 0})


  start_pro := time.Now()
  elapsed := start_pro.Sub(start)
  fmt.Println("Create connection: ",elapsed)
  go func(){

       // Read
       var acc Account
       for {
          db1.First(&acc) // find product with id 1
          if &acc != nil {
            //fmt.Println("Thread 1: account.value = ",acc.Value)
            db1.Model(acc).Update("Value", acc.Value + 1)
            if acc.Value >= 10000 {
              fmt.Println("End Thread 1")
              wg.Done()
              return
            }
          }else{
              fmt.Println("Thread 1: account null ")
          }
        }

  }()
  // go func(){
  //      // Read
  //      var acc Account
  //      for {
  //         db2.First(&acc) // find product with id 1
  //         if &acc != nil {
  //           //fmt.Println("Thread 2: account.value = ",acc.Value)
  //           db2.Model(acc).Update("Value", acc.Value + 1)
  //           if acc.Value >= 10000 {
  //             fmt.Println("End Thread 2")
  //             wg.Done()
  //             return
  //           }
  //         }else{
  //             fmt.Println("Thread 2: account null ")
  //         }
  //       }
  // }()
  // go func(){
  //
  //      // Read
  //      var acc Account
  //      for {
  //         db3.First(&acc) // find product with id 1
  //         if &acc != nil {
  //           //fmt.Println("Thread 3: account.value = ",acc.Value)
  //           db3.Model(acc).Update("Value", acc.Value + 1)
  //           if acc.Value >= 10000 {
  //             wg.Done()
  //             return
  //           }
  //         }else{
  //             fmt.Println("Thread 2: account null ")
  //         }
  //       }
  // }()
  // go func(){
  //
  //      // Read
  //      var acc Account
  //      for {
  //         db4.First(&acc) // find product with id 1
  //         if &acc != nil {
  //           //fmt.Println("Thread 2: account.value = ",acc.Value)
  //           db4.Model(acc).Update("Value", acc.Value + 1)
  //           if acc.Value >= 10000 {
  //             wg.Done()
  //             return
  //           }
  //         }else{
  //             fmt.Println("Thread 4: account null ")
  //         }
  //       }
  // }()
  wg.Wait()
  end_pro := time.Now()
  end_elapsed := end_pro.Sub(start_pro)
  fmt.Println("End Process : ",end_elapsed)
}
