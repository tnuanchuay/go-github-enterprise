package github

import (
	"fmt"
	"encoding/json"
)

func (g *Github) GetPullRequest(number int) (*PullRequest, []error, *Error){
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d?access_token=%s", g.Base_url, g.Organize, g.Repository, number, g.Access_token)
	fmt.Println(url)
	response, body, errs := g.Agent.Get(url).End()

	if response.StatusCode >= 300{
		var githubError Error
		json.Unmarshal([]byte(body), &githubError)
		return nil, nil, &githubError
	}

	if errs != nil{
		return nil, errs, nil
	}

	var pr PullRequest
	err := json.Unmarshal([]byte(body), &pr)
	if err != nil {
		errs = append(errs, err)
		return nil, errs, nil
	}
	return &pr, nil, nil
}

func (g *Github) RemoveLabel(pull_request_number int, label string) (string, []error, *Error){
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels/%s?access_token=%s",
		g.Base_url,
		g.Organize,
		g.Repository,
		pull_request_number,
		label,
		g.Access_token,
	)

	response, body, errs := g.Agent.Delete(url).End()

	if response.StatusCode >= 300{
		var githubError Error
		json.Unmarshal([]byte(body), &githubError)
		return "", nil, &githubError
	}

	if errs != nil{
		return "", errs, nil
	}

	return response.Status, nil, nil
}

func (g *Github) AddLabel(pull_request_number int, label string) ([]error, *Error){
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d/labels?access_token=%s",
		g.Base_url,
		g.Organize,
		g.Repository,
		pull_request_number,
		g.Access_token,
	)

	response, body, errs := g.Agent.Post(url).Send([]string{label}).End()

	if response.StatusCode >= 300 {
		var githubError Error
		json.Unmarshal([]byte(body), &githubError)
		return nil, &githubError
	}

	if errs != nil{
		return errs, nil
	}

	return nil, nil
}


