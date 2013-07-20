package twilio

import (
  "net/url"
  "encoding/json"
)

type Call struct {
  Sid string `json:"sid"`
  ParentCallSid string `json:"parent_call_sid"`
  DateCreated string `json:"date_created"`
  DateUpdated string `json:"date_updated"`
  AccountSid string `json:"account_sid"`
  To string `json:"to"`
  From string `json:"from"`
  PhoneNumberSid string `json:"phone_number_sid"`
  Status string `json:"status"`
  StartTime string `json:"start_time"`
  EndTime string `json:"end_time"`
  Duration string `json:"duration"`
  Price string `json:"price"`
  PriceUnit string `json:"price_unit"`
  Direction string `json:"direction"`
  AnsweredBy string `json:"answered_by"`
  ForwardedFrom string `json:"forwarded_from"`
  CallerName string `json:"caller_name"`
  Uri string `json:"uri"`
}

func MakeCall(client Client, from, to, callback string) (*Call, error) {
  var call *Call

  params := url.Values{}
  params.Set("From", from)
  params.Set("To", to)
  params.Set("Url", callback)

  res, err := client.post(params, client.RootUrl() + "/Calls.json")

  if err != nil {
    return nil, err
  }

  call = new(Call)
  err = json.Unmarshal(res, call)

  return call, err
}