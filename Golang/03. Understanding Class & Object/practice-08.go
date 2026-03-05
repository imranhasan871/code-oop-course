/**
 * Practice 08: Movie Streaming Platform (OOAD)
 * Task: Model a movie streaming platform where users can browse movies,
 *       add to watchlist, watch movies, and rate them.
 *
 * How to run:
 *   go run practice-08.go
 *
 * Key Concepts:
 *   - Structs with slices (lists) as fields
 *   - Methods that manipulate collections
 *   - Average calculation with running totals
 */

package main

import "fmt"

/** Movie represents a movie in the streaming library. */
type Movie struct {
	Title       string
	Genre       string
	Duration    int // in minutes
	TotalRating float64
	RatingCount int
}

/** NewMovie creates a new Movie with no ratings. */
func NewMovie(title, genre string, duration int) *Movie {
	return &Movie{
		Title:    title,
		Genre:    genre,
		Duration: duration,
	}
}

/** AddRating adds a rating score (1-5) to the movie. */
func (m *Movie) AddRating(score int) {
	if score < 1 || score > 5 {
		fmt.Printf("  [Error] Rating must be between 1 and 5 (given: %d).\n", score)
		return
	}
	m.TotalRating += float64(score)
	m.RatingCount++
}

/** AverageRating returns the average rating or 0 if no ratings. */
func (m *Movie) AverageRating() float64 {
	if m.RatingCount == 0 {
		return 0
	}
	return m.TotalRating / float64(m.RatingCount)
}

/** ShowInfo prints movie details. */
func (m *Movie) ShowInfo() {
	fmt.Printf("  Title    : %s\n", m.Title)
	fmt.Printf("  Genre    : %s\n", m.Genre)
	fmt.Printf("  Duration : %d min\n", m.Duration)
	if m.RatingCount > 0 {
		fmt.Printf("  Rating   : %.1f / 5.0 (%d ratings)\n", m.AverageRating(), m.RatingCount)
	} else {
		fmt.Printf("  Rating   : No ratings yet\n")
	}
	fmt.Println()
}

/** User represents a streaming platform user. */
type User struct {
	Username     string
	Watchlist    []*Movie
	WatchHistory []*Movie
}

/** NewUser creates a new User with empty watchlist and history. */
func NewUser(username string) *User {
	return &User{Username: username}
}

/** AddToWatchlist adds a movie to the user's watchlist. */
func (u *User) AddToWatchlist(movie *Movie) {
	// Check if already in watchlist
	for _, m := range u.Watchlist {
		if m == movie {
			fmt.Printf("  [Error] \"%s\" is already in %s's watchlist.\n", movie.Title, u.Username)
			return
		}
	}
	u.Watchlist = append(u.Watchlist, movie)
	fmt.Printf("  [OK] \"%s\" added to %s's watchlist.\n", movie.Title, u.Username)
}

/** WatchMovie moves a movie from watchlist to watch history. */
func (u *User) WatchMovie(movie *Movie) {
	// Check if already watched
	for _, m := range u.WatchHistory {
		if m == movie {
			fmt.Printf("  [Error] %s has already watched \"%s\".\n", u.Username, movie.Title)
			return
		}
	}
	// Remove from watchlist if present
	for i, m := range u.Watchlist {
		if m == movie {
			u.Watchlist = append(u.Watchlist[:i], u.Watchlist[i+1:]...)
			break
		}
	}
	u.WatchHistory = append(u.WatchHistory, movie)
	fmt.Printf("  [OK] %s watched \"%s\".\n", u.Username, movie.Title)
}

/** RateMovie rates a movie the user has watched. */
func (u *User) RateMovie(movie *Movie, score int) {
	// Check if user has watched it
	watched := false
	for _, m := range u.WatchHistory {
		if m == movie {
			watched = true
			break
		}
	}
	if !watched {
		fmt.Printf("  [Error] %s has not watched \"%s\" yet. Watch it first to rate.\n", u.Username, movie.Title)
		return
	}
	movie.AddRating(score)
	fmt.Printf("  [OK] %s rated \"%s\" with %d / 5.\n", u.Username, movie.Title, score)
}

/** ShowInfo prints user details including watchlist and history. */
func (u *User) ShowInfo() {
	fmt.Printf("  Username      : %s\n", u.Username)

	fmt.Print("  Watchlist     : ")
	if len(u.Watchlist) == 0 {
		fmt.Println("Empty")
	} else {
		for i, m := range u.Watchlist {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(m.Title)
		}
		fmt.Println()
	}

	fmt.Print("  Watch History : ")
	if len(u.WatchHistory) == 0 {
		fmt.Println("Empty")
	} else {
		for i, m := range u.WatchHistory {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(m.Title)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// --- Create movie library ---
	movie1 := NewMovie("The Shawshank Redemption", "Drama", 142)
	movie2 := NewMovie("Inception", "Sci-Fi", 148)
	movie3 := NewMovie("The Dark Knight", "Action", 152)

	fmt.Println("=== Movie Library ===")
	movie1.ShowInfo()
	movie2.ShowInfo()
	movie3.ShowInfo()

	// --- Create users ---
	user1 := NewUser("Tareq")
	user2 := NewUser("Afsana")

	// --- Add to watchlist ---
	fmt.Println("=== Adding to Watchlist ===")
	user1.AddToWatchlist(movie1)
	user1.AddToWatchlist(movie2)
	user2.AddToWatchlist(movie2)
	user2.AddToWatchlist(movie3)
	fmt.Println()

	fmt.Println("=== User Info After Adding Watchlist ===")
	user1.ShowInfo()
	user2.ShowInfo()

	// --- Watch movies ---
	fmt.Println("=== Watching Movies ===")
	user1.WatchMovie(movie1)
	user1.WatchMovie(movie2)
	user2.WatchMovie(movie2)
	fmt.Println()

	fmt.Println("=== User Info After Watching ===")
	user1.ShowInfo()
	user2.ShowInfo()

	// --- Rate movies ---
	fmt.Println("=== Rating Movies ===")
	user1.RateMovie(movie1, 5)
	user1.RateMovie(movie2, 4)
	user2.RateMovie(movie2, 5)
	// Try rating unwatched movie
	user2.RateMovie(movie1, 3)
	fmt.Println()

	// --- Final state ---
	fmt.Println("=== Final Movie Ratings ===")
	movie1.ShowInfo()
	movie2.ShowInfo()
	movie3.ShowInfo()
}
