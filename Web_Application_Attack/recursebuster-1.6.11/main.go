package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"runtime"
	"sync"

	"github.com/c-sto/recursebuster/librecursebuster"
	"github.com/fatih/color"
)

const version = "1.6.11"

func main() {
	if runtime.GOOS == "windows" { //lol goos
		//can't use color.Error, because *nix etc don't have that for some reason :(
		librecursebuster.InitLogger(color.Output, color.Output, color.Output, color.Output, color.Output, color.Output, color.Output, color.Output, color.Output, color.Output)
	} else {
		librecursebuster.InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	}

	//the state should probably change per different host. eventually
	globalState := librecursebuster.State{}.Init()
	globalState.Hosts.Init() //**

	globalState.Cfg.Version = version //**
	totesTested := uint64(0)
	globalState.TotalTested = &totesTested
	flag.BoolVar(&globalState.Cfg.ShowAll, "all", false, "Show, and write the result of all checks")                                                                              //Test written
	flag.BoolVar(&globalState.Cfg.AppendDir, "appendslash", false, "Append a / to all directory bruteforce requests (like extension, but slash instead of .yourthing)")           //Test Written
	flag.BoolVar(&globalState.Cfg.Ajax, "ajax", false, "Add the X-Requested-With: XMLHttpRequest header to all requests")                                                         //Test Written
	flag.StringVar(&globalState.Cfg.Auth, "auth", "", "Basic auth. Supply this with the base64 encoded portion to be placed after the word 'Basic' in the Authorization header.") //Test Written
	flag.StringVar(&globalState.Cfg.BadResponses, "bad", "404", "Responses to consider 'bad' or 'not found'. Comma-separated. This works the opposite way of gobuster!")          //Test Written
	flag.StringVar(&globalState.Cfg.GoodResponses, "good", "", "Whitelist of response codes to consider good. If this flag is used, ONLY the provided codes will be considered 'good'.")
	flag.Var(&globalState.Cfg.BadHeader, "badheader", "Check for presence of this header. If an prefix match is found, the response is considered bad.Supply as key:value. Can specify multiple - eg '-badheader Location:cats -badheader X-ATT-DeviceId:XXXXX'")                                                 //Test Written
	flag.StringVar(&globalState.Cfg.BodyContent, "body", "", "File containing content to send in the body of the request. Content-length header will be set accordingly. Note: HEAD requests will/should fail (server will reply with a 400). Set the '-nohead' option to prevent HEAD being sent before a GET.") //Test Written
	flag.StringVar(&globalState.Cfg.BlacklistLocation, "blacklist", "", "Blacklist of prefixes to not check. Will not check on exact matches.")                                                                                                                                                                   //Test Written
	flag.StringVar(&globalState.Cfg.Canary, "canary", "", "Custom value to use to check for wildcards")                                                                                                                                                                                                           //Test Written
	flag.BoolVar(&globalState.Cfg.CleanOutput, "clean", false, "Output clean URLs to the output file for easy loading into other tools and whatnot.")                                                                                                                                                             //todo: add test
	flag.StringVar(&globalState.Cfg.Cookies, "cookies", "", "Any cookies to include with requests. This is smashed into the cookies header, so copy straight from burp I guess? (-cookies 'cookie1=cookie1content; cookie2=1; cookie3=cookie3lol?!?;'.")                                                          //Test Written
	flag.BoolVar(&globalState.Cfg.Debug, "debug", false, "Enable debugging")                                                                                                                                                                                                                                      //todo: add test
	flag.StringVar(&globalState.Cfg.Extensions, "ext", "", "Extensions to append to checks. Multiple extensions can be specified, comma separate them. (-ext 'csv,exe,aspx')")                                                                                                                                    //Test Written
	flag.Var(&globalState.Cfg.Headers, "headers", "Additional headers to include with request. Supply as key:value. Can specify multiple - eg '-headers X-Forwarded-For:127.0.01 -headers X-ATT-DeviceId:XXXXX'")                                                                                                 //Test Written
	flag.BoolVar(&globalState.Cfg.HTTPS, "https", false, "Use HTTPS instead of HTTP.")                                                                                                                                                                                                                            //todo: write test
	flag.StringVar(&globalState.Cfg.InputList, "iL", "", "File to use as an input list of URL's to start from")                                                                                                                                                                                                   //todo: write test
	flag.BoolVar(&globalState.Cfg.SSLIgnore, "k", false, "Ignore SSL check")                                                                                                                                                                                                                                      //todo: write test
	flag.BoolVar(&globalState.Cfg.ShowLen, "len", false, "Show, and write the length of the response")                                                                                                                                                                                                            //todo: write test
	flag.BoolVar(&globalState.Cfg.NoBase, "nobase", false, "Don't perform a request to the base URL")                                                                                                                                                                                                             //todo: write test
	flag.BoolVar(&globalState.Cfg.NoGet, "noget", false, "Do not perform a GET request (only use HEAD request/response)")                                                                                                                                                                                         //Test Written
	flag.BoolVar(&globalState.Cfg.NoEncode, "nodencode", false, "Don't encode non-url safe words in the wordlist")
	flag.BoolVar(&globalState.Cfg.NoHead, "nohead", false, "Don't optimize GET requests with a HEAD (only send the GET)")                                                                                                       //Test Written
	flag.BoolVar(&globalState.Cfg.NoRecursion, "norecursion", false, "Disable recursion, just work on the specified directory. Also disables spider function.")                                                                 //Test Written
	flag.BoolVar(&globalState.Cfg.NoSpider, "nospider", false, "Don't search the page body for links and directories to add to the spider queue.")                                                                              //Test Written
	flag.BoolVar(&globalState.Cfg.NoStatus, "nostatus", false, "Don't print status info (for if it messes with the terminal)")                                                                                                  //todo: write test
	flag.BoolVar(&globalState.Cfg.NoStartStop, "nostartstop", false, "Don't show start/stop info messages")                                                                                                                     //todo: write test
	flag.BoolVar(&globalState.Cfg.NoWildcardChecks, "nowildcard", false, "Don't perform wildcard checks for soft 404 detection (or in plain english, don't do soft404)")                                                        //Test Written
	flag.BoolVar(&globalState.Cfg.NoUI, "noui", false, "Don't use sexy ui")                                                                                                                                                     //todo: write test
	flag.StringVar(&globalState.Cfg.Localpath, "o", "."+string(os.PathSeparator)+"busted.txt", "Local file to dump into")                                                                                                       //todo: write test
	flag.StringVar(&globalState.Cfg.Methods, "methods", "GET", "Methods to use for checks. Multiple methods can be specified, comma separate them. Requests will be sent with an empty body (unless body is specified)")        //Test Written
	flag.StringVar(&globalState.Cfg.ProxyAddr, "proxy", "", "Proxy configuration options in the form ip:port eg: 127.0.0.1:9050. Note! If you want this to work with burp/use it with a HTTP proxy, specify as http://ip:port") //todo: write test
	flag.Float64Var(&globalState.Cfg.Ratio404, "ratio", 0.95, "Similarity ratio to the 404 canary page.")                                                                                                                       //todo: write test
	flag.BoolVar(&globalState.Cfg.NoRobots, "norobots", false, "Don't query and add robots.txt values to checks")
	flag.BoolVar(&globalState.Cfg.FollowRedirects, "redirect", false, "Follow redirects")                                                                                                                                                                       //todo: write test
	flag.BoolVar(&globalState.Cfg.BurpMode, "sitemap", false, "Send 'good' requests to the configured proxy. Requires the proxy flag to be set. ***NOTE: with this option, the proxy is ONLY used for good requests - all other requests go out as normal!***") //todo: write test
	flag.IntVar(&globalState.Cfg.Threads, "t", 1, "Number of concurrent threads")                                                                                                                                                                               //todo: write test
	flag.IntVar(&globalState.Cfg.Timeout, "timeout", 20, "Timeout (seconds) for HTTP/TCP connections")                                                                                                                                                          //todo: write test
	flag.StringVar(&globalState.Cfg.URL, "u", "", "Url to spider")                                                                                                                                                                                              //todo: write test
	flag.StringVar(&globalState.Cfg.Agent, "ua", "RecurseBuster/"+version, "User agent to use when sending requests.")                                                                                                                                          //todo: write test
	flag.IntVar(&globalState.Cfg.VerboseLevel, "v", 0, "Verbosity level for output messages.")                                                                                                                                                                  //todo: write test
	flag.StringVar(&globalState.Cfg.Vhost, "vhost", "", "Vhost to send")                                                                                                                                                                                        //test written
	flag.BoolVar(&globalState.Cfg.ShowVersion, "version", false, "Show version number and exit")                                                                                                                                                                //todo: write test
	flag.StringVar(&globalState.Cfg.Wordlist, "w", "", "Wordlist to use for bruteforce. Blank for spider only")                                                                                                                                                 //todo: write test
	flag.StringVar(&globalState.Cfg.WhitelistLocation, "whitelist", "", "Whitelist of domains to include in brute-force")                                                                                                                                       //todo: write test

	flag.Parse()

	if globalState.Cfg.ShowVersion {
		globalState.PrintBanner()
		os.Exit(0)
	}

	if globalState.Cfg.URL == "" && globalState.Cfg.InputList == "" {
		flag.Usage()
		os.Exit(1)
	}

	urlSlice := getURLSlice(globalState)

	setupConfig(globalState, urlSlice[0])

	globalState.SetupState()

	//do first load of urls (send canary requests to make sure we can dirbust them)
	quitChan := make(chan struct{})
	if !globalState.Cfg.NoUI {
		uiWG := &sync.WaitGroup{}
		uiWG.Add(1)
		go uiQuit(quitChan)
		go globalState.StartUI(uiWG, quitChan)
		uiWG.Wait()
	}

	globalState.StartManagers()

	globalState.PrintOutput("Starting recursebuster...", librecursebuster.Info, 0)

	//seed the workers
	for _, s := range urlSlice {
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}

		if u.Scheme == "" {
			if globalState.Cfg.HTTPS {
				u, err = url.Parse("https://" + s)
			} else {
				u, err = url.Parse("http://" + s)
			}
		}
		if err != nil {
			//this should never actually happen
			panic(err)
		}

		//do canary etc

		prefix := u.String()
		if len(prefix) > 0 && string(prefix[len(prefix)-1]) != "/" {
			prefix = prefix + "/"
		}
		randURL := fmt.Sprintf("%s%s", prefix, globalState.Cfg.Canary)
		//globalState.Chans.GetWorkers() <- struct{}{}
		globalState.AddWG()
		go globalState.StartBusting(randURL, *u)

	}

	//wait for completion
	globalState.Wait()

}

func getURLSlice(globalState *librecursebuster.State) []string {
	urlSlice := []string{}
	if globalState.Cfg.URL != "" {
		urlSlice = append(urlSlice, globalState.Cfg.URL)
	}

	if globalState.Cfg.InputList != "" { //can have both -u flag and -iL flag
		//must be using an input list
		URLList := make(chan string, 10)
		go librecursebuster.LoadWords(globalState.Cfg.InputList, URLList)
		for x := range URLList {
			//ensure all urls will parse good
			_, err := url.Parse(x)
			if err != nil {
				panic("URL parse fail: " + err.Error())
			}
			urlSlice = append(urlSlice, x)
			//globalState.Whitelist[u.Host] = true
		}
	}

	return urlSlice
}

func uiQuit(quitChan chan struct{}) {
	<-quitChan
	os.Exit(0)
}

func setupConfig(globalState *librecursebuster.State, urlSliceZero string) {
	if globalState.Cfg.Debug {
		go func() {
			http.ListenAndServe("localhost:6061", http.DefaultServeMux)
		}()
	}

	var h *url.URL
	var err error
	h, err = url.Parse(urlSliceZero)
	if err != nil {
		panic(err)
	}

	if h.Scheme == "" {
		if globalState.Cfg.HTTPS {
			h, err = url.Parse("https://" + urlSliceZero)
		} else {
			h, err = url.Parse("http://" + urlSliceZero)
		}
	}
	if err != nil {
		panic(err)
	}
	globalState.Hosts.AddHost(h)

	if globalState.Cfg.Canary == "" {
		globalState.Cfg.Canary = librecursebuster.RandString()
	}

}
