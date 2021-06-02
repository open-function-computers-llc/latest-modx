package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Version struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID  string `json:"node_id"`
	Version string `json:"version"`
}

func (v *Version) MarshalJSON() ([]byte, error) {
	// mirrored struct with extra goodies
	return json.Marshal(&struct {
		Name       string `json:"name"`
		ZipballURL string `json:"zipball_url"`
		TarballURL string `json:"tarball_url"`
		CommitSha  string `json:"sha"`
		CommitURL  string `json:"url"`
		NodeID     string `json:"node_id"`
		Version    string `json:"version"`
	}{
		Name:       v.Name,
		ZipballURL: v.ZipballURL,
		TarballURL: v.TarballURL,
		CommitSha:  v.Commit.Sha,
		CommitURL:  v.Commit.URL,
		NodeID:     v.NodeID,
		Version:    parseVersionFromName(v.Name),
	})
}

func (s *Server) GetMODXVersionsFromGithub(filter string) ([]Version, error) {
	v := []Version{}
	validFilters := []string{
		"stable",
		"unstable",
		"rc",
		"alpha",
		"all",
	}
	validFilter := false
	for _, f := range validFilters {
		if f == filter {
			validFilter = true
		}
	}
	if !validFilter {
		return v, errors.New("Invalid filter type")
	}

	s.log.Info("Checking cache for raw github data")
	cachedValue, err := s.cache.Get("githubInfo")
	if err != nil {
		if err.Error() != "Key not found." {
			return v, err
		}
		s.log.Info("Getting data from github")

		resp, err := http.Get("https://api.github.com/repos/modxcms/revolution/tags")
		if err != nil {
			return v, err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &v)

		s.cache.SetWithExpire("githubInfo", body, time.Minute*15)
		s.log.Info("Updated cache for githubInfo")
	}
	cachedValue, _ = s.cache.Get("githubInfo")
	err = json.Unmarshal([]byte(cachedValue.([]uint8)), &v)

	// filter the slice down to the requested subset
	if filter == "all" {
		return v, nil
	}

	v = filterReleases(v, filter)

	return v, nil
}

func filterReleases(set []Version, f string) []Version {
	subset := []Version{}
	for _, v := range set {
		if is(v.Name, f) {
			subset = append(subset, v)
		}
	}
	return subset
}

func is(n string, f string) bool {
	t := parseTypeFromName(n)
	return t == f
}
