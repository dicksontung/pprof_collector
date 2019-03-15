package profiling

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


//Download pprof profile from configured url
func Download() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Unexpected panic occurred: %+v\n", r)
		}
	}()
	log.Println("-- New polling cycle started --")
	defer log.Println("-- Polling cycle ended --")
	t := time.Now().Format("20060102T150405Z")
	if _, err := os.Stat(viper.GetString("output")); os.IsNotExist(err) {
		err = os.Mkdir(viper.GetString("output"), os.ModePerm)
		if err != nil {
			log.Printf("Unable to create folder: %+v\n", err)
			return
		}
	}
	fileName := fmt.Sprintf("%s/%s-%s-%s.pprof", viper.GetString("output"), viper.GetString("application_name"), t, uuid.New().String())

	out, err := os.Create(fileName)
	if err != nil {
		log.Printf("Unable to create file %q: %+v\n", fileName, err)
		return
	}
	defer out.Close()
	url := fmt.Sprintf("%s?seconds=%d", viper.GetString("url"), viper.GetInt("seconds"))
	http.DefaultClient.Timeout = viper.GetDuration("seconds") * 2 * time.Second
	log.Printf("Downloading profile from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Unable to get pprof profile: %+v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Unable to download pprof profile: %+v", resp)
		return
	}
	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("Unable to write file: %+v", err)
		return
	}
	log.Printf("Written to file %q of %d bytes", fileName, n)
}


