package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

var DEFAULT_ROUTE_URL = "http://localhost:8000"
func TestShouldNotLoginWhenBadRequest(t *testing.T) {
  resp, err := http.PostForm(fmt.Sprintf("%s/%s", DEFAULT_ROUTE_URL, "login"), nil)
  if err != nil {
    t.Errorf(fmt.Sprintf("%s: err={%s}", "Can't Request", err))
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil || resp.StatusCode != 400 {
    t.Errorf(fmt.Sprintf("%s: err={%s} resp={%s}", "Didn't get 400", body, err))
  }
}

func TestShouldNotLoginWhenNonExisting(t *testing.T) {
  urlValues := url.Values{}
  urlValues.Set("username", "nonexisting")
  urlValues.Set("password", "couldntcareless")
  resp, err := http.PostForm(fmt.Sprintf("%s/%s", DEFAULT_ROUTE_URL, "login"), urlValues)
  if err != nil {
    t.Errorf(fmt.Sprintf("%s: err={%s}", "Can't Request", err))
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil || resp.StatusCode != 401 {
    t.Errorf(fmt.Sprintf("%s: err={%s} resp={%s}", "Didn't get 401", body, err))
  }
}

func TestShouldNotLoginWhenExistingAndPasswordWrong(t *testing.T) {
  urlValues := url.Values{}
  urlValues.Set("username", "existing")
  urlValues.Set("password", "wrong")
  resp, err := http.PostForm(fmt.Sprintf("%s/%s", DEFAULT_ROUTE_URL, "login"), urlValues)
  if err != nil {
    t.Errorf(fmt.Sprintf("%s: err={%s}", "Can't Request", err))
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil || resp.StatusCode != 401 {
    t.Errorf(fmt.Sprintf("%s: err={%s} resp={%s}", "Didn't get 401", body, err))
  }
}
