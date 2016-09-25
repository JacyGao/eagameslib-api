package title

import (
        "time"
        "io/ioutil"
        "encoding/json"
        "path/filepath"
        "github.com/go-errors/errors"
)

type Title struct {
	Id string       `json:"id"`
    Title string    `json:"title"`
    Created time.Time `json:"created"`
}

const Store = "title/title.json"

func InsertTitle(title Title) error {

    var titles []Title
    
    // TODO: remove hardcoded path
    absPath, _ := filepath.Abs("src/github.com/eagames/store/title/title.json")

    // Maybe marshal and unmarshal should be in seperate file
    contents, err := ioutil.ReadFile(absPath)
    if err := json.Unmarshal(contents, &titles); err != nil {
        panic(err)
    }

    // check if title already exists
    if(len(title.Title) == 0){
        return errors.Wrap("Title name is mandatory", 0)
    }
    // TODO: make it a seperate function
    for _, v := range titles {
        if v.Title == title.Title {
            return errors.Wrap("Title name already exists", 0)
        }
    }

    titles = append(titles, title);
    data, err := json.Marshal(titles); 
    if err != nil {
        panic(err)
    }
	if err := ioutil.WriteFile(absPath, data, 0644); err != nil {
		return errors.Wrap(err, 0)
	}
    return nil
}

func FindTitles() ([]Title, error) {
    absPath, _ := filepath.Abs("src/github.com/eagames/store/title/title.json")
	contents, err := ioutil.ReadFile(absPath)
    if err != nil {
        panic(err)
    }
    if len(contents) == 0 {
        return []Title{}, nil
    }
    var titles []Title
    if err := json.Unmarshal(contents, &titles); err != nil {
        panic(err)
    }
    return titles, nil
}