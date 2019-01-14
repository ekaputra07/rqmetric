package main

import (
  "fmt"
)

type Request struct{
  Url string
  MinTime int
  MaxTime int
  AvgTime float64
  Count int
  OkResponseCount int
  RedirectResponseCount int
  ClientErrorCount int
  ServerErrorCount int
}

func RequestCsvHeader() []string {
  return []string{
    "Url",
    "Count",
    "MinTime",
    "MaxTime",
    "AvgTime",
    "OkResponseCount",
    "RedirectResponseCount",
    "ClientErrorCount",
    "ServerErrorCount",
  }
}

func (req *Request) ToCsvData() []string {
  var data []string
  data = append(data,
    req.Url, 
    fmt.Sprintf("%v", req.Count),
    fmt.Sprintf("%v", req.MinTime),
    fmt.Sprintf("%v", req.MaxTime),
    fmt.Sprintf("%.2f", req.AvgTime),
    fmt.Sprintf("%v", req.OkResponseCount),
    fmt.Sprintf("%v", req.RedirectResponseCount),
    fmt.Sprintf("%v", req.ClientErrorCount),
    fmt.Sprintf("%v", req.ServerErrorCount),
  )
  return data
}

func (req *Request) Add(time int, statusCode int) {
  req.setStatusCodeCount(statusCode)
  req.setAvgTime(time)
  req.Count++

  switch {
    case time < req.MinTime:
      req.MinTime = time
    case time > req.MaxTime:
      req.MaxTime = time
  }
}

/*
 * Calculate new average response time incrementally.
 * AVG_new = ((AVG_old * COUNT_old) + TIME_new) / COUNT_new
 */
func (req *Request) setAvgTime(time int){
  currentAvg := req.AvgTime
  currentCount := float64(req.Count)
  newCount := float64(currentCount + 1)

  req.AvgTime = ((currentAvg * currentCount) + float64(time)) / newCount
}

func (req *Request) setStatusCodeCount(statusCode int){
  switch {
    case statusCode >= 200 && statusCode < 300:
      req.OkResponseCount++
    case statusCode >= 300 && statusCode < 400:
      req.RedirectResponseCount++
    case statusCode >= 400 && statusCode < 500:
      req.ClientErrorCount++
    case statusCode >= 500:
      req.ServerErrorCount++
  }
}

func NewRequest(url string, time int, statusCode int) *Request {
  req := &Request{Url: url, MinTime: time, MaxTime: time, AvgTime: float64(time), Count: 1}
  req.setStatusCodeCount(statusCode)
  return req
}