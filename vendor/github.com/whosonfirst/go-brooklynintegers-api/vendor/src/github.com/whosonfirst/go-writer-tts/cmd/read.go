package main

import (
       "flag"
       "github.com/whosonfirst/go-writer-tts"
       "log"
       "os"
)

func main() {

	var engine = flag.String("engine", "", "Valid options are: osx")
	// var stdout = flag.Bool("stdout", false, "")
	
	flag.Parse()
	args := flag.Args()

	if *engine == "" {
		log.Fatal("You forgot to specify a text-to-speech engine")
	}

	speaker, err := tts.NewSpeakerForEngine(*engine)

	if err != nil {
	   log.Fatal(err)
	}

	for _, path := range args {

		file, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		speaker.Read(file)
	}

	os.Exit(0)
}
