// The MIT License (MIT)

// Copyright (c) 2016 Maciej Borzecki

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	defaultListenPort = 9090

	uriAddMatch = "/api/matches/new"
	uriStatic   = "/static/"
	uriIndex    = "/"
)

var (
	defaultListenAddr = fmt.Sprintf("localhost:%d",
		defaultListenPort)
)

func apiAddMatch(w http.ResponseWriter, req *http.Request) {
	// add new match
	log.Printf("add new match")

	matchdata, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("failed to receive match data: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(matchdata) == 0 {
		log.Printf("no data provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("match data: %s", matchdata)

	match, err := LoadMatchData(matchdata)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("match hash: %s", match.DataHash)
	w.Write([]byte(match.DataHash))
}

func setupRouting() {
	r := mux.NewRouter()

	r.HandleFunc(uriIndex, siteHomeHandler).
		Methods("GET")

	// matches only come through POST
	r.HandleFunc(uriAddMatch, apiAddMatch).
		Methods("POST")

	// static files
	staticroot := path.Join(C.webroot, "static")
	log.Printf("serving static files from %s", staticroot)

	filehandler := http.FileServer(http.Dir(staticroot))
	r.PathPrefix(uriStatic).
		Handler(http.StripPrefix(uriStatic, filehandler))

	// setup logging for all handlers
	lr := handlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", lr)
}

func daemonMain() error {

	setupSite()
	setupRouting()

	return http.ListenAndServe(fmt.Sprintf(":%d", C.port), nil)
}

func runDaemon() error {
	err := LoadConfig()
	if err != nil {
		return errors.Wrap(err, "daemon startup failed")
	}

	log.Printf("listen port: %d", C.port)

	return daemonMain()
}