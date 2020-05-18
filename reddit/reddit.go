package reddit

import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"log"
	"fmt"
	"os"
)

 const (
	 RedditTokenEndpoint = "https://www.reddit.com/api/v1/access_token"
 )

type Application struct {
	ClientId string
	ClientSecret string
	UserAgent string
}

func NewApp(clientId string, clientSecret string, userAgent string) *Application {
	return &Application{
		ClientId : clientId,
		ClientSecret : clientSecret,
		UserAgent : userAgent,
	}
}

func (app *Application) Auth() *Reddit {
	// Prepare Request to get Auth Token
	grantType := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", RedditTokenEndpoint, grantType)
	req.SetBasicAuth(app.ClientId, app.ClientSecret)
	req.Header.Set("User-Agent", app.UserAgent)

	client := &http.Client{}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Check for valid response
	if res.StatusCode != 200 {
		log.Fatal("unable to auth client")
	}

	// Extract request body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}


	var reddit Reddit
	err = json.Unmarshal(body, &reddit)
	if err != nil {
		log.Fatal(err)
	}

	//Pass along the user agent
	reddit.UserAgent = app.UserAgent


	return &reddit

}

type Reddit struct {
	AccessToken string `json:"access_token"`
	UserAgent string
}

func (r *Reddit) GetPostComments(postId string) []Comment {
	url := fmt.Sprintf("https://oauth.reddit.com/comments/%s?depth=0&sort=top&showmore=true",postId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}

	req.Header.Set("Authorization", "bearer " + r.AccessToken)
	req.Header.Set("User-Agent", r.UserAgent)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		log.Fatal(err)
	}
	comments := parseCommentJson(jsonBody)

	return comments
}

func parseCommentJson(raw interface{}) []Comment {

	comments := []Comment{}

	array := raw.([]interface{})

	for _, list := range array {
		data := list.(map[string]interface{})["data"]
		children := data.(map[string]interface{})["children"]
		for _, child := range children.([]interface{}) {
			childMap := child.(map[string]interface{})
			kind := childMap["kind"].(string)
			if kind == "t1" {
				commentData := childMap["data"].(map[string]interface{})
				commentDataString, err := json.Marshal(commentData)
				if err != nil {
					log.Fatal(err)
				}

				var comment Comment
				err = json.Unmarshal(commentDataString, &comment)
				if err != nil {
					log.Fatal(err)
				}
				comments = append(comments, comment)
			}

		}
	}

	return comments
}
