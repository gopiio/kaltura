# Kaltura API Client for Go
This is a very simple kaltura client written in Go.  

## Import package
```
go get github.com/gopiio/kaltura
```
Include the import as 
```
import("github.com/gopiio/kaltura/kaltura")
```


## Definition
Kaltura is defined as struct as follows
```
	kalturaClient := &kaltura.Kaltura{
		Secret:     REPLACE_SECRET,
		PartnerID:  REPLACE_PARTNER_ID,
		ServiceURL: REPLACE_SERVICE_URL,
		Format:     kaltura.KalturaFormatJSON,
		Session: &kaltura.Session{
			UserID:     REPLACE_USER_ID,
			Type:       kaltura.KalturaAdminSession,
			Duration:   REPLACE_WITH_DURATION,
			Privileges: "",
		},
	}
```

## How to use
- Start the session by kalturaClient.SessionStart()
- To get the resource kalturaClient.DoRequest("GET",SERVICE_TO_USE,ACTION,PAYLOAD as map[string]interface{})