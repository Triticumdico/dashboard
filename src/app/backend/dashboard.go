package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Triticumdico/dashboard/src/app/backend/args"
	"github.com/Triticumdico/dashboard/src/app/backend/client"
	"github.com/Triticumdico/dashboard/src/app/backend/handler"
	"github.com/spf13/pflag"
)

var (
	argConfigYamlPath      = pflag.String("config-yaml-path", "./config.yaml", "path to read yaml config file")
	argInsecurePort        = pflag.Int("insecure-port", 9090, "port to listen to for incoming HTTP requests")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "IP address on which to serve the --insecure-port, set to 127.0.0.1 for all interfaces")
)

func main() {

	// Set logging output to standardconsol out
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = flag.CommandLine.Parse(make([]string, 0)) // Init for glog calls in kubernetes packages

	// Initializes dashboard arguments holder so we can read them in other packages
	initArgHolder()
	// Initializes dashboard configurationq holder so we can read them in other packages
	initConfig()
	// Initializes dashboard Datatabase connection so we can read them in other packages
	initDatabase()

	apiHandler, err := handler.CreateHTTPAPIHandler()
	if err != nil {
		handleFatalInitError(err)
		log.Fatal(err)
	}

	// Run a HTTP server that serves static public files from './public' and handles API calls.
	http.Handle("/api/", apiHandler)

	// Listen for http or https
	log.Printf("Serving insecurely on HTTP port: %d", args.Holder.GetInsecurePort())
	addr := fmt.Sprintf("%s:%d", args.Holder.GetInsecureBindAddress(), args.Holder.GetInsecurePort())
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()

	select {}

}

func initArgHolder() {
	builder := args.GetHolderBuilder()
	builder.SetInsecurePort(*argInsecurePort)
	builder.SetInsecureBindAddress(*argInsecureBindAddress)
	builder.SetConfigYamlPath(*argConfigYamlPath)
}

func initConfig() {
	builderConf := args.GetConfigBuilder()
	builderConf.SetYamlConfig()
}

func initDatabase() {
	ClientDb := client.NewClientDb()
	ClientDb.OpenDbConnection("postgres")
}

/**
 * Handles fatal init error that prevents server from doing any work. Prints verbose error
 * message and quits the server.
 */
func handleFatalInitError(err error) {
	log.Fatalf("Error while initializing connection to Kubernetes apiserver. "+
		"This most likely means that the cluster is misconfigured (e.g., it has "+
		"invalid apiserver certificates or service account's configuration) or the "+
		"--apiserver-host param points to a server that does not exist. Reason: %s\n", err)
}
