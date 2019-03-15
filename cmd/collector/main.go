package main

import (
	"fmt"
	"github.com/pprof_collector/pkg"
	"github.com/pprof_collector/pkg/awsS3"
	"github.com/pprof_collector/pkg/server"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	configureFlags()
	log.Printf("Start polling pprof profile every %d seconds\n", viper.GetInt("seconds"))
	log.Printf("Transferring to %q \n", viper.GetString("bucket"))
	go server.Profiling()
	for {
		profiling.Download()
		uploadAll()
		cooldown()
	}

}

func configureFlags(){
	viper.AutomaticEnv()
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Printf("\nDescription:\n")
		fmt.Printf("  Continuously collects the pprof profile from `net/http/pprof` endpoints and sends it to a AWS Bucket\n")
		fmt.Printf("\nOptions:\n")
		pflag.PrintDefaults()
	}
	pflag.IntP("seconds", "s", 30, "Polling interval in seconds")
	pflag.IntP("cooldown", "", 5, "Cool down period in seconds")
	pflag.StringP("aws_access_key", "u", "", "[required] AWS Access Key")
	pflag.StringP("aws_secret_key", "p", "", "[required] AWS Secret Key")
	pflag.StringP("region", "r", "", "[required] Region for the S3 Bucket")
	pflag.StringP("bucket", "b", "", "[required] S3 Bucket name")
	pflag.StringP("application_name", "", "undefined", "Application name to identify the current profile")
	pflag.StringP("url", "", "http://localhost/debug/pprof/profile", "Url for pprof profile")
	pflag.StringP("output", "o", "./tmp/pprof_collector", "Output directory")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalf("Unable to parse env variable: %+v\n", err)
	}
}

func cooldown() {
	time.Sleep(viper.GetDuration("cooldown") * time.Second)
}

func uploadAll() {
	files, err := ioutil.ReadDir(viper.GetString("output"))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s3.Upload(viper.GetString("output") + "/" + f.Name())
	}
}


