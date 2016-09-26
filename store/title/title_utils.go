package title

import (
        "github.com/go-errors/errors"
)

func validateTitle(title string, titles []Title) error{
	// Check if title is empty
    if(len(title) == 0){
        return errors.Wrap("Title name is mandatory", 0)
    }
    // Check if title already exists
    for _, v := range titles {
        if v.Title == title {
            return errors.Wrap("Title name already exists", 0)
        }
    }
    return nil
}