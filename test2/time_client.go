package main

import (
        "fmt"
        "github.com/beevik/ntp"
        "time"
        "os/exec"
)
func UpdateDate(time_server string ){
      response, err := ntp.Query(time_server)
      networkTime := time.Now().Add(response.ClockOffset)
      err = SetSystemDate(networkTime)
      if err != nil {
         fmt.Println("Can not update system time.")
      }

      // ntpTime, err := ntp.Time(time_server)
      // if err != nil {
      //     fmt.Println(err)
      // }

      ntpTimeFormatted := networkTime.Format(time.RFC3339Nano)
      //
      // cmd := exec.Command("date", "--set",ntpTimeFormatted)
      // err = cmd.Run()
      // if err != nil {
      //     fmt.Println("Command finished with error: ", err)
      // }
      timeFormatted := time.Now().Local().Format(time.RFC3339Nano)
      fmt.Println("Unix Date Network time: ", ntpTimeFormatted)
      fmt.Println("Unix Date System time:", timeFormatted)
}
func SetSystemDate(newTime time.Time) error {
    _, lookErr := exec.LookPath("date")
    if lookErr != nil {
        fmt.Printf("Date binary not found, cannot set system date: %s\n", lookErr.Error())
        return lookErr
    } else {
      //dateString := newTime.Format("2006-01-2 15:4:5")
      dateString := newTime.Format("2 Jan 2006 15:04:05")
      fmt.Printf("Setting system date to: %s\n", dateString)
      args := []string{"--set", dateString}
      return exec.Command("date", args...).Run()
    }
}

func main() {
        time_server := "time.apple.com"
      	tick := time.Tick(5 * time.Second)
      	// Keep trying until we're timed out or got a result or got an error
      	for {
      		select {
          		case <-tick:
          			UpdateDate(time_server)
      		   }
      	}
}
