package main

import (
       "bufio"
       "flag"
       "github.com/whosonfirst/go-writer-tts"
       "io"
       "log"
       "os"
       "strings"
)

func main() {

	var engine = flag.String("engine", "", "Valid options are: osx")
	var stdout = flag.Bool("stdout", false, "")
	
	flag.Parse()
	args := flag.Args()

	if *engine == "" {
		log.Fatal("You forgot to specify a text-to-speech engine")
	}

	speaker, err := tts.NewSpeakerForEngine(*engine)

	if err != nil {
	   log.Fatal(err)
	}

	writers := []io.Writer{
		speaker,
	}

	if *stdout {
		writers = append(writers, os.Stdout)
	}
	
	multi := io.MultiWriter(writers...)
	writer := bufio.NewWriter(multi)

	if len(args) == 1 && args[0] == "-" {

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			msg := scanner.Text()
			writer.WriteString(msg + "\n")
			writer.Flush()					       
  		}
	} else {

		msg := strings.Join(args, " ")
		writer.WriteString(msg + "\n")
		writer.Flush()					       
	}

	os.Exit(0)
}
