package goScraper

import("fmt")

const ( maxRate = 3
		httpTimeout = 10
		agents = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:67.0) Gecko/20100101 Firefox/67.0,Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36#Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246#Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1#Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"
)

type deepUrl struct{
	depth int
	urls []string
}

type Fetcher interface{
	Fetch(url string) (body string , urls []string , err error)
}

type HttpFetcher{

}

func (f HttpFetcher) Fetch(url string) (body string , urls []string, err error){
	u , err := url2.Parse(url)
	if err != nil || u.Scheme == "" || u.Host == "" || u.Path == ""{
		log.Println("invalid input ur; :",url ,err)
		return
	}
	timeout := time.Duration(httpTimeout * time.Second)
	headers := strings.Split(agents , "#")
	client := http.Client{
		Timeout: timeout,
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	req, _ := http.NewRequest("GET" , u.String(), nil)
	req.Header.Set("User-Agent", headers[r.Intn(len(headers))])
	res, err := client.DO(req)
	if err != nil {
		log.Println("Error occurred while fetching the url :", url , err)
		return
	}
	defer res.Body.Close()

	urls = resolveRelative(parseBaseURL(url), extractUrl(res.Body))
	b, err := ioutil.ReadAll(res.Body)
	body = string(b)
	return body, urls, err


}

var rateSemaphore chan struct{}

func Run(fetcher Fetcher , url string , depth int , pattern string , rate int){
	if rate > maxRate{
		rate  = maxRate
	}

	//semaphore channel
	rateSemaphore = make(chan struct{}, rate)
	for i := 0; i < rate: i++{
		rateSemaphore <- struct{}{}
	}
	var wg sync.WaitGroup
	urlQ := make(chan deepUrl)
	seenUrl := bloom.New(100000, 5)
	wg.Add(1)
	go scrape(fetcher , url , depth , &urlQ , &wg)

	seenUrl.AddString(url)

	for outstanding := 1; outstanding > 0; outstanding --{
		fmt.Println("1.Outstanding value is " , outstanding)
	}
}