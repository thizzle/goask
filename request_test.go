package goask

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestIntentRequest(t *testing.T) {
	f, err := os.Open("./testassets/http-echo-intent-request")
	if err != nil {
		t.Error(err)
		return
	}

	buf := bufio.NewReader(f)
	req, err := http.ReadRequest(buf)
	if err != nil {
		t.Error(err)
		return
	}

	ar := new(AlexaRequestContainer)
	err = json.NewDecoder(req.Body).Decode(ar)
	if err != nil {
		t.Error(err)
		return
	}

	if ar.Version != "1.0" {
		t.Error("version is incorrect")
	}

	if ar.Session.New != false {
		t.Error("session.new must be false")
	}
	if ar.Session.SessionID != "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000" {
		t.Error("session.sessionId is incorrect")
	}
	if ar.Session.Application.ApplicationID != "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe" {
		t.Error("session.application.applicationId is incorrect")
	}
	if ar.Session.User.UserID != "amzn1.account.AM3B00000000000000000000000" {
		t.Error("session.user.userId is incorrect")
	}

	if ar.Request.Type != "IntentRequest" {
		t.Error("request.type is incorrect")
	}
	if ar.Request.RequestID != " amzn1.echo-api.request.0000000-0000-0000-0000-00000000000" {
		t.Error("request.requestId is incorrect")
	}
	actualTimestamp, _ := time.Parse(time.RFC3339, "2015-05-13T12:34:56Z")
	if !actualTimestamp.Equal(ar.Request.Timestamp) {
		t.Error("request.timestamp is incorrect")
	}

	if ar.Request.Intent.Name != "GetZodiacHoroscopeIntent" {
		t.Error("request.intent.name is incorrect")
	}
	if _, ok := ar.Request.Intent.Slots["ZodiacSign"]; !ok {
		t.Error("request.intent.slots is incorrect")
	}
	if ar.Request.Intent.Slots["ZodiacSign"].Name != "ZodiacSign" {
		t.Error("request.intent.slots.ZodiacSign.name is incorrect")
	}
	if ar.Request.Intent.Slots["ZodiacSign"].Value != "virgo" {
		t.Error("request.intent.slots.ZodiacSign.value is incorrect")
	}
}
