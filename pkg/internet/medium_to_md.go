/*
Copyright © 2021 Renato Torres <renato.torres@pm.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package internet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/renato0307/canivete-core/interface/internet"
)

type mediumQuery struct {
	Query string `json:"query"`
}

type mediumPostResponse struct {
	Data mediumData `json:"data"`
}

type mediumData struct {
	Post mediumPost `json:"post"`
}

type mediumPost struct {
	Title   string            `json:"title"`
	Creator mediumPostCreator `json:"creator"`
	Content mediumPostContent `json:"content"`
}

type mediumPostCreator struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type mediumPostContent struct {
	BodyModel mediumPostContentBodyModel `json:"bodyModel"`
}

type mediumPostContentBodyModel struct {
	Paragraphs []mediumPostParagraph `json:"paragraphs"`
}

type mediumPostParagraph struct {
	Text     string                      `json:"text"`
	Type     string                      `json:"type"`
	HRef     string                      `json:"href"`
	IFrame   mediumPostParagraphIFrame   `json:"iframe"`
	Layout   string                      `json:"layout"`
	Markups  []mediumPostParagraphMarkup `json:"markups"`
	Metadata mediumPostParagraphMetadata `json:"metadata"`
}

type mediumPostParagraphMarkup struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	HRef       string `json:"href"`
	Start      int    `json:"start"`
	End        int    `json:"end"`
	Rel        string `json:"rel"`
	AnchorType string `json:"anchorType"`
}

type mediumPostParagraphMetadata struct {
	TypeName       string `json:"__typename"`
	Id             string `json:"id"`
	OriginalWidth  int    `json:"originalWidth"`
	OriginalHeight int    `json:"originalHeight"`
}

type mediumPostParagraphIFrame struct {
	MediaResource mediumPostParagraphIFrameMediaResource `json:"mediaResource"`
}

type mediumPostParagraphIFrameMediaResource struct {
	HRef         string `json:"href"`
	IFrameSrc    string `json:"iframeSrc"`
	IFrameWidth  int    `json:"iframeWidth"`
	IFrameHeight int    `json:"iframeHeight"`
}

func (s *Service) ConvertMediumToMd(postId string) (internet.ConvertMediumToMdOutput, error) {
	output := internet.ConvertMediumToMdOutput{}

	result, err := getPostData(postId)
	if err != nil {
		return output, fmt.Errorf("error getting post data: %s", err.Error())
	}

	output.Markdown = postToMarkdown(result)
	output.PostId = postId

	return output, nil
}

func getPostData(postId string) (mediumPostResponse, error) {

	post := mediumPostResponse{}

	url := "https://medium.com/_/graphql"

	query := fmt.Sprintf(
		`
		query {
			post(id: "%s") {
			  title
			  createdAt
			  creator {
				id
				name
			  }
			  content {
				bodyModel {
				  paragraphs {
					text
					type
					href
					layout
					markups {
					  title
					  type
					  href
					  userId
					  start
					  end
					  anchorType
					}
					iframe {
					  mediaResource {
						href
						iframeSrc
						iframeWidth
						iframeHeight
					  }
					}
					metadata {
					  id
					  originalWidth
					  originalHeight
					}
				  }
				}
			  }
			}
		  }
		`,
		postId)

	queryStruct := mediumQuery{Query: query}

	data, err := json.Marshal(queryStruct)
	if err != nil {
		return post, fmt.Errorf("error marshling request: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return post, fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return post, fmt.Errorf("error executing request: %s", err.Error())
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return post, fmt.Errorf("error reading response: %s", err.Error())
	}

	// Convert response to the struct
	err = json.Unmarshal(body, &post)
	if err != nil {
		return post, fmt.Errorf("error un-marshalling medium response: %s", err.Error())
	}

	return post, nil
}

func postToMarkdown(post mediumPostResponse) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("# %s\n", post.Data.Post.Title))
	buffer.WriteString(fmt.Sprintf("By %s\n", post.Data.Post.Creator.Name))

	for _, paragraph := range post.Data.Post.Content.BodyModel.Paragraphs {
		if paragraph.Type == "H3" {
			buffer.WriteString(fmt.Sprintf("\n## %s\n", paragraph.Text))
		} else if paragraph.Type == "H4" {
			buffer.WriteString(fmt.Sprintf("\n### _%s_\n", paragraph.Text))
		} else if paragraph.Type == "P" {
			buffer.WriteString(fmt.Sprintf("\n%s\n", paragraph.Text))
		} else if paragraph.Type == "IMG" {
			buffer.WriteString(fmt.Sprintf("\n![%s](https://miro.medium.com/max/1400/%s)\n", paragraph.Text, paragraph.Metadata.Id))

		}
		if len(paragraph.Markups) > 0 {
			textParts := []string{}
			lastStartIndex := 0
			for _, markup := range paragraph.Markups {
				if markup.Type != "A" {
					continue
				}
				textParts = append(textParts, paragraph.Text[lastStartIndex:markup.Start])
				textParts = append(textParts, fmt.Sprintf("[%s](%s)",
					paragraph.Text[markup.Start:markup.End],
					markup.HRef))
				lastStartIndex = markup.End
			}
			buffer.WriteString(fmt.Sprintf("\n%s\n", strings.Join(textParts, "")))
		}
	}

	return buffer.String()
}
