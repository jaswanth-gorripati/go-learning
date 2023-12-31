package concurrency

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var WEB_URLs = []string{"https://go.dev/tour/methods/9", "https://google.com", "http://examle.com", "https://aciana.com", "https://archents.com"}

type WebLinks interface {
	StatusCheck(context.Context, string)
}

type WebURLS struct {
	Urls []string
}

type HttpResponse struct {
	URL        string
	StatusCode int
	Message    string
	Error      error
	Latency    time.Duration
}

func (w *WebURLS) DoHttpCall(ctx context.Context, result chan HttpResponse, URL string) {
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		result <- HttpResponse{
			URL:        URL,
			StatusCode: 0,
			Message:    "Failed",
			Error:      fmt.Errorf("failed to create request for url %v , due to :%v ", URL, err),
			Latency:    time.Since(startTime),
		}
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		result <- HttpResponse{
			URL:        URL,
			StatusCode: 0,
			Message:    "Failed",
			Error:      fmt.Errorf("URL calling failed for %v , due to :%v", URL, err),
			Latency:    time.Since(startTime),
		}
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result <- HttpResponse{
			URL:        URL,
			StatusCode: resp.StatusCode,
			Message:    "Success",
			Error:      nil,
			Latency:    time.Since(startTime),
		}
		return
	} else {
		err = fmt.Errorf("invalid status %v for URL:%v", resp.StatusCode, URL)
		result <- HttpResponse{
			URL:        URL,
			StatusCode: resp.StatusCode,
			Message:    "Failed",
			Error:      err,
			Latency:    time.Since(startTime),
		}
	}
}

func (w *WebURLS) GetAllStatus() {
	httpResp := make(chan HttpResponse, len(w.Urls))
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2)
	// defer cancel()

	for _, url := range w.Urls {

		go func(url string) {
			w.DoHttpCall(context.Background(), httpResp, url)
		}(url)
	}

	for i := 0; i < len(w.Urls); i++ {
		resp := <-httpResp

		log.Debugf("URL: %v , Status: %v ,Latency : %v ,  Error: %v", resp.URL, resp.Message, resp.Latency, resp.Error)
	}
	close(httpResp)
}
