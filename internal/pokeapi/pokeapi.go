package pokeapi

import (
	"errors"
	"fmt"
	"net/http"
	"io"

	"github.com/daniilgaltsev/pokedex-cli/internal/cache"
)

func Request(url string, cache cache.Cache) ([]byte, error) {
	var body []byte
	body, ok := cache.Get(url)
	if !ok {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, err = io.ReadAll(response.Body)
		response.Body.Close()
		if response.StatusCode > 299 {
			return nil, errors.New(fmt.Sprintf("Response code %d: %s", response.StatusCode, body))
		}
		if err != nil {
			return nil, err
		}

		cache.Add(url, body)
	}

	return body, nil
}
