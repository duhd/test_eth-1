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

type BankA struct {
  gorm.Model
  Code string
  Value uint
}

type BankB struct {
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

  // Migrate the schema
  db.AutoMigrate(&BankA{})
  db.AutoMigrate(&BankB{})

  var banka BankA
  db.First(&banka) // find product with code l1212
  if &banka != nil {
      db.Delete(&banka)
  }
  // Create
  db.Create(&BankA{Code: "banka", Value: 0})


  var bankb BankB
  db.First(&bankb) // find product with code l1212
  if &bankb != nil {
      db.Delete(&bankb)
  }
  // Create
  db.Create(&BankB{Code: "bankb", Value: 0})

  start_pro := time.Now()
  elapsed := start_pro.Sub(start)
  fmt.Println("Create connection: ",elapsed)
  go func(){
       // Update bank A
       var ba BankA
       for i:=0; i<10000; i++ {
          db1.First(&ba) // find product with id 1
          if &ba != nil {
            fmt.Println("Thread 1: ba.value = ",ba.Value)
            db1.Model(ba).Update("Value", ba.Value + 1)
          }else{
              fmt.Println("Thread 1: ba null ")
          }
        }
        // Update bank b
        var bb BankB
        for i:=0; i<10000; i++ {
           db2.First(&bb) // find product with id 1
           if &bb != nil {
             fmt.Println("Thread 1: bb.value = ",bb.Value)
             db2.Model(bb).Update("Value", bb.Value + 1)
           }else{
               fmt.Println("Thread 1: bb null ")
           }
         }

        fmt.Println("Thread 1: Finished ")
        defer wg.Done()
  }()
  // go func(){
  //      // Read
  //      var bb BankB
  //      for i:=0; i<10000; i++ {
  //         db2.First(&bb) // find product with id 1
  //         if &bb != nil {
  //           fmt.Println("Thread 2: bb.value = ",bb.Value)
  //           db2.Model(bb).Update("Value", bb.Value + 1)
  //         }else{
  //             fmt.Println("Thread 2: bb null ")
  //         }
  //       }
  //       fmt.Println("Thread 2: Finished ")
  //       defer wg.Done()
  // }()

  wg.Wait()
  end_pro := time.Now()
  end_elapsed := end_pro.Sub(start_pro)
  fmt.Println("End Process : ",end_elapsed)
}
