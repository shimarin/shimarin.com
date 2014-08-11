package main

import (
	"fmt"
	"flag"
	"net"
	"net/http"
	"net/http/fcgi"
	"sync"
	"html/template"
	"encoding/json"
	"os"
	"bufio"
	"strings"
	"regexp"
	"path/filepath"
)

var file_server = http.FileServer(http.Dir("."))
var html_re, _ = regexp.Compile("\\.html$")

func html(filename string, w http.ResponseWriter, r *http.Request) {
	page_name := html_re.ReplaceAllString(filepath.Base(filename), "")
	json_filename := filepath.Join(filepath.Dir(filename), "index.json")
	var json_data interface{}
	fp, err := os.Open(json_filename)
	if err == nil {
		dec := json.NewDecoder(bufio.NewReader(fp))
		dec.Decode(&json_data)
		defer fp.Close()
	}
	pages := func(data interface{}) []interface{} {
		if data == nil {
			return make([]interface{},0)
		} else {
			return data.([]interface{})
		}
	}(json_data)
	page_by_name := make(map[string] interface{})
	for _, v := range pages {
		page_by_name[v.(map[string] interface{})["name"].(string)] = v.(interface{})
	}

	data := func(page_name string) map[string] interface{} {
		if page_by_name[page_name] == nil {
			return make(map[string] interface{})
		} else {
			return page_by_name[page_name].(map[string] interface{})
		}
	}(page_name)

	funcs := template.FuncMap {
		"pages":func() []interface{} {
			return pages
		},
		"get_page":func(name string) map[string] interface{} {
			return page_by_name[name].(map[string] interface{})
		},
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 not found")
		return
	}
	t := template.Must(template.ParseFiles("root.html"))
	t.Funcs(funcs)

	if page_name != "index" {
		base_filename := filepath.Join(filepath.Dir(filename), "base.html")
		if _, err := os.Stat(base_filename); err == nil {
			t.ParseFiles(base_filename)
		}
	}

	template.Must(t.ParseFiles(filename))
	t.Execute(w, data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	filename := func(path string)(ret string) {
		if strings.HasSuffix(path, "/") {
			ret = path[1:] + "index.html"
		} else {
			ret = path[1:]
		}
		return
	}(r.URL.Path)
	
	if html_re.FindStringIndex(filename) != nil {
		html(filename, w, r)
		return
	}
	// else
	file_server.ServeHTTP(w, r)
}

func main() {
	var fcgi_enabled = flag.Bool("fcgi", false, "Enable fastcgi alongside http")
	flag.Parse()

	js := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	http.Handle("/favicon.ico", file_server)
	http.Handle("/robots.txt", file_server)
	http.HandleFunc("/", handler)

	wg := &sync.WaitGroup{}

	http_listener,_ := net.Listen("tcp", ":5000")
	wg.Add(1)
	go func() {
		http.Serve(http_listener, nil)
	}()

	if *fcgi_enabled {
		fcgi_listener,_ := net.Listen("tcp", ":9000")
		wg.Add(1)
		go func() {
			fcgi.Serve(fcgi_listener, nil)
		}()
	}

	wg.Wait()
}
