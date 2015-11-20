/**
 * logthis.go - GitHub.com/DavidLiedle/loggly-blog-go-logthis
 * Supporting example code for the Loggly blog post on ""What to Log in Go"
 *
 * @author David Christian Liedle
 * @link   http://davidcanhelp.me/
 *
 * Thanks to http://dave.cheney.net/2015/11/05/lets-talk-about-logging
 * and https://gist.github.com/heatxsink/7221ebe499b0767d4784
 * for the inspiring thoughts and code snippets.
 *
 * USAGE: `go run logthis.go -logtostderr=true -v=2`
 *
 */

package main

import (
	"github.com/golang/glog"
 	"os"
 	"flag"
	"fmt"
)

func usage(){
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init(){
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line 
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {

	fmt.Println("Examples of what to log (and how).")
	glog.V(2).Infoln("We've begun our program...")

	fmt.Println("Here's what we'll log:")
	glog.V(2).Infoln("We're up to the part where we talk about what we're going to log...")

	//glog.V(2).Flush()

	number_of_lines := 100
	for i := 0; i < number_of_lines; i++ {
		glog.V(2).Infof("LINE: %d", i)
		message := fmt.Sprintf("TEST LINE: %d", i)
		glog.Error(message)
	}

}
