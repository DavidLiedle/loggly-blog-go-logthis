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
	"github.com/golang/glog" // An external package, from Google
	//"log" // The internal package (just to show it in action). Un-comment to play with the isolated snippet.
 	"os"
 	"flag"
	"fmt"
)

func usage(){
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
} // End of func usage()

func init(){

	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line 
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()

	// Here's an isolated snippet to show the internal log package in action:
		/*
		f, err := os.Open("config.json") // We're catching both return types: the file, AND the error
		if err != nil { // Conditionally do something IF there's anything of note in err
		    log.Fatal(err) // For example, kill it with fire before it lays eggs
		}
		*/
	// ...and then you'd do something with the open config *File f, since err == nil and you're good to Go!

	// THIS EXAMPLE WILL KILL YOUR CODE EVERYTIME,
	// so comment it out, or create the "config.json" file for great success! :D
	// Because it is a breaking example, it's up to you to un-comment the region and configure as desired if
	// you want to play with it.

} // End of func init()

func main() {

	fmt.Println("Examples of what to log (and how).")
	glog.V(2).Infoln("We've begun our program...")

	fmt.Println("Here's what we'll log:")
	glog.V(2).Infoln("We're up to the part where we talk about what we're going to log...")

	//glog.V(2).Flush() // cuz that's a thing

	number_of_lines := 100
	for i := 0; i < number_of_lines; i++ {
		glog.V(2).Infof("LINE: %d", i)
		message := fmt.Sprintf("TEST LINE: %d", i)
		glog.Error(message)
	}

} // End of func main()
