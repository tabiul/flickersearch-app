package flicker

import (
	"com/flickersearch/test"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "jsonFlickrApi({\"photos\":{\"page\":1,\"pages\":54786,\"perpage\":5,\"total\":\"273929\",\"photo\":[{\"id\":\"27401099743\",\"owner\":\"130926299@N05\",\"secret\":\"6bbb8004c8\",\"server\":\"7439\",\"farm\":8,\"title\":\"SAM_1222\",\"ispublic\":1,\"isfriend\":0,\"isfamily\":0},{\"id\":\"27401271524\",\"owner\":\"27809489@N03\",\"secret\":\"fbf24480e3\",\"server\":\"7409\",\"farm\":8,\"title\":\"Play\",\"ispublic\":1,\"isfriend\":0,\"isfamily\":0}]},\"stat\":\"ok\"})")
	}))
	defer server.Close()
	response, err := http.Get(server.URL)
	test.AssertNotError(t, err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	test.AssertNotError(t, err)
	images, err := processResponse(body)
	test.AssertNotError(t, err)
	test.AssertEquals(t, 2, len(images.ImageURL))

}

func TestDecodeSearchJSONData(t *testing.T) {
	imageURLS, err := decodeSearchJSONData([]byte("{\"photos\":{\"page\":1,\"pages\":54786,\"perpage\":5,\"total\":\"273929\",\"photo\":[{\"id\":\"27401099743\",\"owner\":\"130926299@N05\",\"secret\":\"6bbb8004c8\",\"server\":\"7439\",\"farm\":8,\"title\":\"SAM_1222\",\"ispublic\":1,\"isfriend\":0,\"isfamily\":0},{\"id\":\"27401271524\",\"owner\":\"27809489@N03\",\"secret\":\"fbf24480e3\",\"server\":\"7409\",\"farm\":8,\"title\":\"Play\",\"ispublic\":1,\"isfriend\":0,\"isfamily\":0}]},\"stat\":\"ok\"}"))
	test.AssertNotError(t, err)
	test.AssertEquals(t, 1, imageURLS.Photos.Page)
	test.AssertEquals(t, 54786, imageURLS.Photos.Pages)
	test.AssertEquals(t, 2, len(imageURLS.Photos.Photo))
	test.AssertEquals(t, "27401099743", imageURLS.Photos.Photo[0].ID)
	test.AssertEquals(t, "6bbb8004c8", imageURLS.Photos.Photo[0].Secret)
	test.AssertEquals(t, 8, imageURLS.Photos.Photo[0].FarmID)
	test.AssertEquals(t, "7439", imageURLS.Photos.Photo[0].ServerID)
	test.AssertEquals(t, "ok", imageURLS.Status)

}

func TestBuildPhotoURL(t *testing.T) {
	images := buildPhotoURL([]flickerPhoto{
		flickerPhoto{
			ID:       "27401099743",
			FarmID:   1,
			Secret:   "secret",
			ServerID: "123",
		},
	})
	test.AssertEquals(t, 1, len(images))
	test.AssertEquals(t, "https://farm1.staticflickr.com/123/27401099743_secret_t.jpg", images[0].ThumbNail)
	test.AssertEquals(t, "https://farm1.staticflickr.com/123/27401099743_secret_b.jpg", images[0].Large)
}
