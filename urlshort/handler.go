package urlshort

import (
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, path, http.StatusFound)
		}
	}
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJSON)
	return MapHandler(pathMap, fallback), nil
}

func BoltDBHandler(db *bolt.DB, fallback http.Handler) (http.HandlerFunc, error) {
	bucketName := []byte("paths")
	boltDB := make(map[string]string)

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			boltDB[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return MapHandler(boltDB, fallback), nil
}
