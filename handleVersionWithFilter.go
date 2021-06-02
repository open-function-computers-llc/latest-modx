package main

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleVersionWithFilter(f string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		versions, err := s.GetMODXVersionsFromGithub(f)
		if err != nil {
			errOutput := map[string]string{
				"error": err.Error(),
			}
			o, _ := json.Marshal(errOutput)
			w.Write(o)
			return
		}

		jsonOutput, err := json.Marshal(&versions)
		if err != nil {
			errOutput := map[string]string{
				"error": err.Error(),
			}
			jsonOutput, _ = json.Marshal(errOutput)
		}
		w.Write(jsonOutput)
	}
}
