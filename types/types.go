package types

import (
	"strconv"
	"time"
)

type (
	// Languages !
	Languages struct {
		Popular []struct {
			URLParam string `json:"urlParam"`
			Name     string `json:"name"`
		} `json:"popular"`

		All []struct {
			URLParam string `json:"urlParam"`
			Name     string `json:"name"`
		} `json:"all"`
	}

	// Repos !
	Repos []struct {
		Author             string        `json:"author"`
		Name               string        `json:"name"`
		URL                string        `json:"url"`
		Description        string        `json:"description"`
		Language           string        `json:"language"`
		LanguageColor      string        `json:"languageColor"`
		Stars              int           `json:"stars"`
		Forks              int           `json:"forks"`
		CurrentPeriodStars int           `json:"currentPeriodStars"`
		BuiltBy            []interface{} `json:"builtBy"`
	}

	// Records !
	Records map[string]Repos
)

func currentDate() string {
	now := time.Now()
	return strconv.Itoa(now.Day()) + "-" +
		strconv.Itoa(int(now.Month())) + "-" +
		strconv.Itoa(now.Year())
}

func (r *Records) urlExists(url *string) bool {
	for _, repos := range *r {
		for _, repo := range repos {
			if *url == repo.URL {
				return true
			}
		}
	}
	return false
}

// Consume write data in database
// skip duplicate urls in entry
// should this method be in db ?
func (r *Records) Consume(repos Repos) {
	index := currentDate()
	var items Repos
	if oldItems, ok := (*r)[index]; ok {
		items = oldItems
	}

	for _, repo := range repos {
		if !r.urlExists(&repo.URL) {
			// skip urls that exists
			// create new record set
			items = append(items, repo)
		}
	}

	(*r)[index] = items
}
