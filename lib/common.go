package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

var remoteHTTPClient = func() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryWaitMin = 5 * time.Minute
	client.RetryWaitMax = 5 * time.Minute
	client.RetryMax = 10
	return client
}()

// GetRemoteURLResponse performs a GET with the shared retry policy.
// The caller must close the response body.
func GetRemoteURLResponse(url string) (*http.Response, error) {
	return remoteHTTPClient.Get(url)
}

func GetRemoteURLContent(url string) ([]byte, error) {
	resp, err := GetRemoteURLResponse(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get remote content -> %s: %s", url, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func GetRemoteURLReader(url string) (io.ReadCloser, error) {
	resp, err := GetRemoteURLResponse(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get remote content -> %s: %s", url, resp.Status)
	}

	return resp.Body, nil
}

func GetIgnoreIPType(onlyIPType IPType) IgnoreIPOption {
	switch onlyIPType {
	case IPv4:
		return IgnoreIPv6
	case IPv6:
		return IgnoreIPv4
	}

	return nil
}

type WantedListExtended struct {
	TypeSlice []string
	TypeMap   map[string][]string
}

func (w *WantedListExtended) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	slice := make([]string, 0)
	mapMap := make(map[string][]string, 0)

	err := json.Unmarshal(data, &slice)
	if err != nil {
		err2 := json.Unmarshal(data, &mapMap)
		if err2 != nil {
			return err2
		}
	}

	w.TypeSlice = slice
	w.TypeMap = mapMap

	return nil
}
