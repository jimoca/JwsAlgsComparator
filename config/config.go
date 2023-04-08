package config

import (
	"strconv"
	"time"
)

var (
	JsonString = `{
	"sub": "member001",
	"aud": "kioxiaApp2",
	"nbf": 1661700103,
	"scope": [
	  "read",
	  "openid",
	  "write"
	],
	"iss": "http://192.168.0.100:30200",
	"exp": ` + strconv.FormatInt(time.Now().Add(time.Hour*24).Unix(), 10) + `,
	"iat": ` + strconv.FormatInt(time.Now().Unix(), 10) + `,
	"device": {
	  "deviceId": 1,
	  "employee": null,
	  "name": "Win",
	  "firewall": true,
	  "antivirus": false,
	  "os": "Windows",
	  "osVersion": "21H2",
	  "osLatestUpdate": 1516239022,
	  "deviceType": "Computer",
	  "firstLoginTime": 1661700100159,
	  "riskAnalysis": "Unevaluated",
	  "deviceActivityList": [
		{
		  "deviceActivityId": 1,
		  "ipAddress": "0:0:0:0:0:0:0:1",
		  "activityStatus": "Latest",
		  "loginTime": 1661700100165
		},
		{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },
		  {
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  },{
			"deviceActivityId": 1,
			"ipAddress": "0:0:0:0:0:0:0:1",
			"activityStatus": "Latest",
			"loginTime": 1661700100165
		  }
	  ]
	}
  }`
)
