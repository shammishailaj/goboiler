package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type ApiRequest struct {
	Log                *log.Logger
	Url                string            `json:"url"`
	Data               string            `json:"data"`
	Headers            map[string]string `json:"headers"`
	Async              bool              `json:"async"`
	Timeout            int               `json:"timeout"`
	RequestMethod      string            `json:"request_method"`
	ScheduledTimeStamp int32             `json:"scheduled_time_stamp"`
	DataFormat         string            `json:"data_format"`
	ExtraInformation   string            `json:"extra_information"`
	AmazonTraceId      string            `json:"amzn_trace_id"`
}

type ApiResponse struct {
	UUID             string     `json:"uuid"`
	TimeTaken        string     `json:"time_taken"`
	ResponseHeaders  string     `json:"response_headers"`
	Response         string     `json:"response"`
	ApiDataExtraInfo string     `json:"extra_information"`
	ApiRequestInfo   ApiRequest `json:"api_request_info"`
}

func (u *Utils) SendRequest(p *ApiRequest, myChan chan ApiResponse, result ApiResponse) {

	client := resty.New()
	var resp *resty.Response
	var err error
	if p.RequestMethod == "GET" {
		resp, err = client.R().
			SetHeaders(p.Headers).
			EnableTrace().
			Get(p.Url)

	} else {
		resp, err = client.R().
			SetHeaders(p.Headers).
			EnableTrace().
			SetBody(p.Data).
			Post(p.Url)

	}
	//Save result in response struct
	if resp != nil {
		result.TimeTaken = resp.Time().String()

		jsonString, jsonErr := json.Marshal(resp.Header())
		if jsonErr != nil {
			result.ResponseHeaders = fmt.Sprint(resp.Header())
		} else {
			result.ResponseHeaders = string(jsonString)
		}

		result.Response = fmt.Sprint(resp)
		if myChan != nil {
			myChan <- result
		}
		//connectCouchBase(result.UUID, result)

		u.Log.Printf("[%s] Response Info:", result.UUID)
		u.Log.Printf("[%s]  Error      : %v", result.UUID, err)
		u.Log.Printf("[%s]  Status Code: %d", result.UUID, resp.StatusCode())
		u.Log.Printf("[%s]  Status     :%s", result.UUID, resp.Status())
		u.Log.Printf("[%s]  Proto      :%s", result.UUID, resp.Proto())
		u.Log.Printf("[%s]  Time       :%s", result.UUID, resp.Time())
		u.Log.Printf("[%s]  raw header       :%v", result.UUID, resp.RawResponse)
		u.Log.Printf("[%s]  raw header       :%s", result.UUID, resp.Header())
		u.Log.Printf("[%s]  Received At:%s", result.UUID, resp.ReceivedAt())
		u.Log.Printf("[%s]  Body       :\n %s", result.UUID, resp)

		for key, element := range resp.Header() {
			u.Log.Printf("[%s]response header  Key:%s => Element: %s", result.UUID, key, element)
		}

		// Explore trace info
		u.Log.Printf("[%s]Request Trace Info:", result.UUID)
		ti := resp.Request.TraceInfo()
		u.Log.Printf("[%s]  DNSLookup     :%s", result.UUID, ti.DNSLookup)
		u.Log.Printf("[%s]  ConnTime      :%s", result.UUID, ti.ConnTime)
		u.Log.Printf("[%s]  TCPConnTime   :%s", result.UUID, ti.TCPConnTime)
		u.Log.Printf("[%s]  TLSHandshake  :%s", result.UUID, ti.TLSHandshake)
		u.Log.Printf("[%s]  ServerTime    :%s", result.UUID, ti.ServerTime)
		u.Log.Printf("[%s]  ResponseTime  :%s", result.UUID, ti.ResponseTime)
		u.Log.Printf("[%s]  TotalTime     :%s", result.UUID, ti.TotalTime)
		u.Log.Printf("[%s]  IsConnReused  :%t", result.UUID, ti.IsConnReused)
		u.Log.Printf("[%s]  IsConnWasIdle :%t", result.UUID, ti.IsConnWasIdle)
		u.Log.Printf("[%s]  ConnIdleTime  :%s", result.UUID, ti.ConnIdleTime)
		u.Log.Printf("[%s]  RequestAttempt:%d", result.UUID, ti.RequestAttempt)
		u.Log.Printf("[%s]  RemoteAddr    :%s", result.UUID, ti.RemoteAddr.String())
	}

}

func (p *ApiRequest) Send() ApiResponse {
	u := NewUtils(p.Log)
	var uuid string

	u.Log.Printf("[%s]p.AmazonTraceId : ", p.AmazonTraceId)
	if len(p.AmazonTraceId) > 0 {
		uuid = p.AmazonTraceId
	} else {
		uuid = u.NewUUID()
	}

	u.Log.Printf("[%s]uuid : ", uuid)
	u.Log.Printf("[%s] data in request %s", uuid, p.Url)
	responseResult := ApiResponse{UUID: uuid}

	myChan := make(chan ApiResponse)

	go u.SendRequest(p, myChan, responseResult)
	responseResult = <-myChan
	return responseResult
}

func (p *ApiRequest) AsyncSend() ApiResponse {
	u := NewUtils(p.Log)
	var uuid string
	u.Log.Printf("[%s]p.AmazonTraceId : ", p.AmazonTraceId)
	if len(p.AmazonTraceId) > 0 {
		uuid = p.AmazonTraceId
	} else {
		uuid = u.NewUUID()
	}
	u.Log.Printf("[%s]uuid : ", uuid)
	u.Log.Printf("[%s] data in request %s", uuid, p.Url)
	responseResult := ApiResponse{UUID: uuid}

	return responseResult
}

func (p *ApiRequest) CallApi(responseResult ApiResponse) ApiResponse {
	u := NewUtils(p.Log)
	myChan := make(chan ApiResponse)
	go u.SendRequest(p, myChan, responseResult)
	responseResult = <-myChan

	return responseResult
}
