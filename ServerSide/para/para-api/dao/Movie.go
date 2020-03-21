package dao

import (
	"para/para-api/model"
	"sync"
)

// MovieRepository handles the basic operations of a movie entity/model.
// It's an interface in order to be testable, i.e a memory movie repository or
// a connected to an sql database.

type Query func(movie model.Movie) bool
type MovieRepository interface {
	//Exec(query Query, action Query, limit int, mode int)(ok bool)

	Select(query Query)(movie model.Movie, found bool)
	//SelectMany(query Query, limit int64)(results []model.Movie)
	//
	//InsertOrUpdate(movie model.Movie)(updateMovie model.Movie, err error)
	//Delete(query Query, limit int)(delete bool)
}

// NewMovieRepository returns a new movie memory-based repository,
// the one and only repository type in our example.
func NewMovieRepository(source map[int64]model.Movie) MovieRepository{
	return &movieMemoryRepository{source:source}
}


// movieMemoryRepository is a "MovieRepository"
// which manages the movies using the memory data source (map).
type movieMemoryRepository struct {
	source map[int64]model.Movie
	mu sync.RWMutex
}

func (r *movieMemoryRepository) Select(query Query)(movie model.Movie, found bool)  {

}