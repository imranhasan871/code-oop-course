/**
 * Practice 08: Movie Streaming Platform (OOAD)
 *
 * Scenario: An online movie streaming platform offers a large library of
 * films. Users can browse the available movies, add them to their watchlist,
 * and rate the ones they've seen. The system keeps track of each user's
 * watch history and recommends similar films based on their preferences.
 *
 * Main Objects:
 *   - Movie    — knows: id, title, genre, rating
 *   - User     — knows: name; does: add to watchlist, watch, rate
 *   - Platform — knows: movie library, users; does: add movie, recommend
 *
 * Key Concepts:
 *   - Multiple structs collaborating
 *   - Slices and maps for collections
 *   - Encapsulation with unexported fields
 *   - Constructor functions
 *   - Business logic (recommendations based on genre)
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

package main

import "fmt"

// Movie represents a film in the library.
type Movie struct {
	movieID     string
	title       string
	genre       string
	totalRating float64
	ratingCount int
}

/** NewMovie creates a new movie with no ratings yet. */
func NewMovie(movieID, title, genre string) Movie {
	return Movie{movieID: movieID, title: title, genre: genre}
}

/** GetAverageRating returns the average user rating, or 0 if no ratings. */
func (m *Movie) GetAverageRating() float64 {
	if m.ratingCount == 0 {
		return 0
	}
	return m.totalRating / float64(m.ratingCount)
}

/** AddRating adds a rating (1-5) to the movie. */
func (m *Movie) AddRating(rating float64) {
	m.totalRating += rating
	m.ratingCount++
}

/** PrintInfo displays movie details. */
func (m *Movie) PrintInfo() {
	fmt.Printf("  [%s] %s | Genre: %s | Avg Rating: %.1f (%d ratings)\n",
		m.movieID, m.title, m.genre, m.GetAverageRating(), m.ratingCount)
}

// User represents a platform user.
type User struct {
	name         string
	watchlist    []*Movie
	watchHistory []*Movie
}

/** NewUser creates a new user with empty watchlist and history. */
func NewUser(name string) User {
	return User{name: name, watchlist: []*Movie{}, watchHistory: []*Movie{}}
}

/** AddToWatchlist adds a movie to the user's watchlist. */
func (u *User) AddToWatchlist(movie *Movie) {
	u.watchlist = append(u.watchlist, movie)
	fmt.Printf("  [OK] %s added '%s' to watchlist.\n", u.name, movie.title)
}

/** WatchMovie moves a movie from watchlist to watch history. */
func (u *User) WatchMovie(movie *Movie) {
	// Remove from watchlist if present
	for i, m := range u.watchlist {
		if m.movieID == movie.movieID {
			u.watchlist = append(u.watchlist[:i], u.watchlist[i+1:]...)
			break
		}
	}
	u.watchHistory = append(u.watchHistory, movie)
	fmt.Printf("  [OK] %s watched '%s'.\n", u.name, movie.title)
}

/** RateMovie rates a movie the user has watched. */
func (u *User) RateMovie(movie *Movie, rating float64) {
	// Check if user has watched this movie
	watched := false
	for _, m := range u.watchHistory {
		if m.movieID == movie.movieID {
			watched = true
			break
		}
	}
	if !watched {
		fmt.Printf("  [Error] %s hasn't watched '%s' yet.\n", u.name, movie.title)
		return
	}
	if rating < 1 || rating > 5 {
		fmt.Println("  [Error] Rating must be between 1 and 5.")
		return
	}
	movie.AddRating(rating)
	fmt.Printf("  [OK] %s rated '%s' — %.1f/5\n", u.name, movie.title, rating)
}

/** ShowWatchlist displays the user's watchlist. */
func (u *User) ShowWatchlist() {
	fmt.Printf("  %s's Watchlist (%d movies):\n", u.name, len(u.watchlist))
	for _, m := range u.watchlist {
		fmt.Printf("    - %s (%s)\n", m.title, m.genre)
	}
	if len(u.watchlist) == 0 {
		fmt.Println("    (empty)")
	}
}

/** ShowWatchHistory displays movies the user has watched. */
func (u *User) ShowWatchHistory() {
	fmt.Printf("  %s's Watch History (%d movies):\n", u.name, len(u.watchHistory))
	for _, m := range u.watchHistory {
		fmt.Printf("    - %s (%s)\n", m.title, m.genre)
	}
	if len(u.watchHistory) == 0 {
		fmt.Println("    (empty)")
	}
}

// StreamingPlatform manages the movie library and users.
type StreamingPlatform struct {
	name    string
	movies  []*Movie
}

/** NewStreamingPlatform creates a new platform with an empty library. */
func NewStreamingPlatform(name string) StreamingPlatform {
	return StreamingPlatform{name: name, movies: []*Movie{}}
}

/** AddMovie adds a movie to the library. */
func (p *StreamingPlatform) AddMovie(movie *Movie) {
	p.movies = append(p.movies, movie)
	fmt.Printf("  [OK] Added '%s' to %s library.\n", movie.title, p.name)
}

/** BrowseMovies displays all movies in the library. */
func (p *StreamingPlatform) BrowseMovies() {
	fmt.Printf("  %s Library (%d movies):\n", p.name, len(p.movies))
	for _, m := range p.movies {
		m.PrintInfo()
	}
}

/** RecommendMovies suggests movies based on the user's watch history genres. */
func (p *StreamingPlatform) RecommendMovies(user *User) {
	// Collect genres from watch history
	genreSet := map[string]bool{}
	watchedIDs := map[string]bool{}
	for _, m := range user.watchHistory {
		genreSet[m.genre] = true
		watchedIDs[m.movieID] = true
	}

	fmt.Printf("  Recommendations for %s:\n", user.name)
	found := false
	for _, m := range p.movies {
		if genreSet[m.genre] && !watchedIDs[m.movieID] {
			fmt.Printf("    - %s (%s) — Avg Rating: %.1f\n", m.title, m.genre, m.GetAverageRating())
			found = true
		}
	}
	if !found {
		fmt.Println("    (no recommendations available)")
	}
}

func main() {
	fmt.Println("=== Practice 08: Movie Streaming Platform ===")
	fmt.Println()

	// Create platform
	platform := NewStreamingPlatform("CineStream")

	// Add movies
	fmt.Println("--- Add Movies ---")
	m1 := NewMovie("M-001", "The Matrix", "Sci-Fi")
	m2 := NewMovie("M-002", "Interstellar", "Sci-Fi")
	m3 := NewMovie("M-003", "The Dark Knight", "Action")
	m4 := NewMovie("M-004", "Inception", "Sci-Fi")
	m5 := NewMovie("M-005", "John Wick", "Action")

	platform.AddMovie(&m1)
	platform.AddMovie(&m2)
	platform.AddMovie(&m3)
	platform.AddMovie(&m4)
	platform.AddMovie(&m5)
	fmt.Println()

	// Browse library
	fmt.Println("--- Browse Movies ---")
	platform.BrowseMovies()
	fmt.Println()

	// Create user
	user := NewUser("Imtiaz")

	// Add to watchlist
	fmt.Println("--- Add to Watchlist ---")
	user.AddToWatchlist(&m1)
	user.AddToWatchlist(&m2)
	user.AddToWatchlist(&m3)
	fmt.Println()

	fmt.Println("--- Watchlist ---")
	user.ShowWatchlist()
	fmt.Println()

	// Watch movies
	fmt.Println("--- Watch Movies ---")
	user.WatchMovie(&m1)
	user.WatchMovie(&m3)
	fmt.Println()

	// Rate movies
	fmt.Println("--- Rate Movies ---")
	user.RateMovie(&m1, 5)
	user.RateMovie(&m3, 4.5)
	user.RateMovie(&m2, 4) // should fail — not watched yet
	fmt.Println()

	// Show history and updated watchlist
	fmt.Println("--- Watch History ---")
	user.ShowWatchHistory()
	fmt.Println()

	fmt.Println("--- Updated Watchlist ---")
	user.ShowWatchlist()
	fmt.Println()

	// Recommendations
	fmt.Println("--- Recommendations ---")
	platform.RecommendMovies(&user)
}
