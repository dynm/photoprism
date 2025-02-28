package search

import (
	"testing"

	"github.com/photoprism/photoprism/internal/form"
	"github.com/stretchr/testify/assert"
)

func TestPhotosFilterPath(t *testing.T) {
	t.Run("2790/07", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2790/07"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("2790*", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2790*"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("London", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "London"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("London pipe 2790/07", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "London|2790/07"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 2)
	})
	t.Run("StartsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "%abc/%folderx"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "ab%c/fol%de"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "abc%/folde%"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "&abc/&folde"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "tes&r/lo&c"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2020&/vacation&"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "'2020/'vacation"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "20'20/vacat'ion"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2020'/vacation'"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "*2020/*vacation"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 4)
	})
	t.Run("CenterAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "202*3/vac*ation"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2023*/vacatio*"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "|202/|vacation"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "20|22/vacat|ion"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 0, len(photos))
	})
	t.Run("EndsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2022|/vacation|"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2000/holiday"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2000/02"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 2)
	})
	t.Run("EndsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2000/02"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 2)
	})
	t.Run("StartsWithDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "\"2000/\"02"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "20\"00/0\"2"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2000\"/02\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = " 2000/ 02"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "20 00/ 0 2"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "2000 /02 "
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("OrSearch", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "%abc/%folderx|20'20/vacat'ion"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 2)
	})
	t.Run("OrSearch2", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "202*3/vac*ation|20'20/vacat'ion"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 1)
	})
	t.Run("OrSearch3", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "20'20/vacat'ion|&abc/&folde"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 1)
	})
	t.Run("OrSearch4", func(t *testing.T) {
		var f form.SearchPhotos

		f.Path = "London|1990/04"
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 2)
	})
}

func TestPhotosQueryPath(t *testing.T) {
	t.Run("2790/07", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2790/07\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("2790*", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2790*\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("London", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"London\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("London pipe 2790/07", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"London|2790/07\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 2)
	})
	t.Run("StartsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"%abc/%folderx\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"ab%c/fol%de\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithPercent", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"abc%/folde%\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"&abc/&folde\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"tes&r/lo&c\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2020&/vacation&\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"'2020/'vacation\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"20'20/vacat'ion\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2020'/vacation'\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"*2020/*vacation\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 4)
	})
	t.Run("CenterAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"202*3/vac*ation\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2023*/vacatio*\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("StartsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"|202/|vacation\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"20|22/vacat|ion\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 0, len(photos))
	})
	t.Run("EndsWithPipe", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2022|/vacation|\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2000/holiday\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2000/02\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 2)
	})
	t.Run("EndsWithNumber", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2000/02\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 2)
	})
	t.Run("StartsWithDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"\"2000/\"02\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		//TODO
		assert.Greater(t, len(photos), 1)
	})
	t.Run("CenterDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"20\"00/0\"2\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		t.Log(photos[0].PhotoPath)
		t.Log(photos[1].PhotoPath)
		//TODO
		assert.Greater(t, len(photos), 1)
	})
	t.Run("EndsWithDoubleQuotes", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2000\"/02\"\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		//TODO
		assert.Greater(t, len(photos), 1)
	})
	t.Run("StartsWithWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\" 2000/ 02\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 1)
	})
	t.Run("CenterWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"20 00/ 0 2\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("EndsWithWhitespace", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"2000 /02 \""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 1)
	})
	t.Run("OrSearch", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"%abc/%folderx|20'20/vacat'ion\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 2)
	})
	t.Run("OrSearch2", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"202*3/vac*ation|20'20/vacat'ion\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 1)
	})
	t.Run("OrSearch3", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"20'20/vacat'ion|&abc/&folde\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 1)
	})
	t.Run("OrSearch4", func(t *testing.T) {
		var f form.SearchPhotos

		f.Query = "path:\"London|1990/04\""
		f.Merged = true

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(photos), 2)
	})
}
