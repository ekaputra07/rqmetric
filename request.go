package main

import (
  "fmt"
)

type Request struct{
  url string
  minTime int
  maxTime int
  avgTime float64
  count int
  okResponseCount int
  redirectResponseCount int
  clientErrorCount int
  serverErrorCount int
}

func RequestCsvHeader() []string {
  return []string{
    "url",
    "count",
    "minTime",
    "maxTime",
    "avgTime",
    "okResponseCount",
    "redirectResponseCount",
    "clientErrorCount",
    "serverErrorCount",
  }
}

func (req *Request) ToCsvData() []string {
  var data []string
  data = append(data,
    req.url, 
    fmt.Sprintf("%v", req.count),
    fmt.Sprintf("%v", req.minTime),
    fmt.Sprintf("%v", req.maxTime),
    fmt.Sprintf("%.2f", req.avgTime),
    fmt.Sprintf("%v", req.okResponseCount),
    fmt.Sprintf("%v", req.redirectResponseCount),
    fmt.Sprintf("%v", req.clientErrorCount),
    fmt.Sprintf("%v", req.serverErrorCount),
  )
  return data
}

func (req *Request) Add(time int, statusCode int) {
  req.setStatusCodeCount(statusCode)
  req.setAvgTime(time)
  req.count++

  switch {
    case time < req.minTime:
      req.minTime = time
    case time > req.maxTime:
      req.maxTime = time
  }
}

/*
 * Calculate new average response time incrementally.
 * AVG_new = ((AVG_old * COUNT_old) + TIME_new) / COUNT_new
 */
func (req *Request) setAvgTime(time int){
  currentAvg := req.avgTime
  currentCount := float64(req.count)
  newCount := float64(currentCount + 1)

  req.avgTime = ((currentAvg * currentCount) + float64(time)) / newCount
}

func (req *Request) setStatusCodeCount(statusCode int){
  switch {
    case statusCode >= 200 && statusCode < 300:
      req.okResponseCount++
    case statusCode >= 300 && statusCode < 400:
      req.redirectResponseCount++
    case statusCode >= 400 && statusCode < 500:
      req.clientErrorCount++
    case statusCode >= 500:
      req.serverErrorCount++
  }
}

func NewRequest(url string, time int, statusCode int) *Request {
  req := &Request{url: url, minTime: time, maxTime: time, avgTime: float64(time), count: 1}
  req.setStatusCodeCount(statusCode)
  return req
}