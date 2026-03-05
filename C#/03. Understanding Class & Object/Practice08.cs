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
 *   - Multiple classes collaborating
 *   - List<T> and HashSet<T> for collections
 *   - Encapsulation with private fields and properties
 *   - Constructor design
 *   - Business logic (recommendations based on genre)
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

using System;
using System.Collections.Generic;

/** Represents a film in the library. */
class Movie
{
    private string movieID;
    private string title;
    private string genre;
    private double totalRating;
    private int ratingCount;

    public Movie(string movieID, string title, string genre)
    {
        this.movieID = movieID;
        this.title = title;
        this.genre = genre;
        this.totalRating = 0;
        this.ratingCount = 0;
    }

    public string MovieID { get { return movieID; } }
    public string Title { get { return title; } }
    public string Genre { get { return genre; } }

    public double GetAverageRating()
    {
        return ratingCount == 0 ? 0 : totalRating / ratingCount;
    }

    public void AddRating(double rating)
    {
        totalRating += rating;
        ratingCount++;
    }

    public void PrintInfo()
    {
        Console.WriteLine($"  [{movieID}] {title} | Genre: {genre} | " +
                          $"Avg Rating: {GetAverageRating():F1} ({ratingCount} ratings)");
    }
}

/** Represents a platform user. */
class User
{
    private string name;
    private List<Movie> watchlist;
    private List<Movie> watchHistory;

    public User(string name)
    {
        this.name = name;
        this.watchlist = new List<Movie>();
        this.watchHistory = new List<Movie>();
    }

    public string Name { get { return name; } }
    public List<Movie> WatchHistory { get { return watchHistory; } }

    public void AddToWatchlist(Movie movie)
    {
        watchlist.Add(movie);
        Console.WriteLine($"  [OK] {name} added '{movie.Title}' to watchlist.");
    }

    public void WatchMovie(Movie movie)
    {
        watchlist.RemoveAll(m => m.MovieID == movie.MovieID);
        watchHistory.Add(movie);
        Console.WriteLine($"  [OK] {name} watched '{movie.Title}'.");
    }

    public void RateMovie(Movie movie, double rating)
    {
        bool watched = false;
        foreach (Movie m in watchHistory)
        {
            if (m.MovieID == movie.MovieID)
            {
                watched = true;
                break;
            }
        }
        if (!watched)
        {
            Console.WriteLine($"  [Error] {name} hasn't watched '{movie.Title}' yet.");
            return;
        }
        if (rating < 1 || rating > 5)
        {
            Console.WriteLine("  [Error] Rating must be between 1 and 5.");
            return;
        }
        movie.AddRating(rating);
        Console.WriteLine($"  [OK] {name} rated '{movie.Title}' — {rating:F1}/5");
    }

    public void ShowWatchlist()
    {
        Console.WriteLine($"  {name}'s Watchlist ({watchlist.Count} movies):");
        foreach (Movie m in watchlist)
        {
            Console.WriteLine($"    - {m.Title} ({m.Genre})");
        }
        if (watchlist.Count == 0)
        {
            Console.WriteLine("    (empty)");
        }
    }

    public void ShowWatchHistory()
    {
        Console.WriteLine($"  {name}'s Watch History ({watchHistory.Count} movies):");
        foreach (Movie m in watchHistory)
        {
            Console.WriteLine($"    - {m.Title} ({m.Genre})");
        }
        if (watchHistory.Count == 0)
        {
            Console.WriteLine("    (empty)");
        }
    }
}

/** Manages the movie library and users. */
class StreamingPlatform
{
    private string name;
    private List<Movie> movies;

    public StreamingPlatform(string name)
    {
        this.name = name;
        this.movies = new List<Movie>();
    }

    public void AddMovie(Movie movie)
    {
        movies.Add(movie);
        Console.WriteLine($"  [OK] Added '{movie.Title}' to {name} library.");
    }

    public void BrowseMovies()
    {
        Console.WriteLine($"  {name} Library ({movies.Count} movies):");
        foreach (Movie m in movies)
        {
            m.PrintInfo();
        }
    }

    public void RecommendMovies(User user)
    {
        HashSet<string> genres = new HashSet<string>();
        HashSet<string> watchedIDs = new HashSet<string>();
        foreach (Movie m in user.WatchHistory)
        {
            genres.Add(m.Genre);
            watchedIDs.Add(m.MovieID);
        }

        Console.WriteLine($"  Recommendations for {user.Name}:");
        bool found = false;
        foreach (Movie m in movies)
        {
            if (genres.Contains(m.Genre) && !watchedIDs.Contains(m.MovieID))
            {
                Console.WriteLine($"    - {m.Title} ({m.Genre}) — Avg Rating: {m.GetAverageRating():F1}");
                found = true;
            }
        }
        if (!found)
        {
            Console.WriteLine("    (no recommendations available)");
        }
    }
}

class Practice08
{
    static void Main(string[] args)
    {
        Console.WriteLine("=== Practice 08: Movie Streaming Platform ===");
        Console.WriteLine();

        // Create platform
        StreamingPlatform platform = new StreamingPlatform("CineStream");

        // Add movies
        Console.WriteLine("--- Add Movies ---");
        Movie m1 = new Movie("M-001", "The Matrix", "Sci-Fi");
        Movie m2 = new Movie("M-002", "Interstellar", "Sci-Fi");
        Movie m3 = new Movie("M-003", "The Dark Knight", "Action");
        Movie m4 = new Movie("M-004", "Inception", "Sci-Fi");
        Movie m5 = new Movie("M-005", "John Wick", "Action");

        platform.AddMovie(m1);
        platform.AddMovie(m2);
        platform.AddMovie(m3);
        platform.AddMovie(m4);
        platform.AddMovie(m5);
        Console.WriteLine();

        // Browse library
        Console.WriteLine("--- Browse Movies ---");
        platform.BrowseMovies();
        Console.WriteLine();

        // Create user
        User user = new User("Imtiaz");

        // Add to watchlist
        Console.WriteLine("--- Add to Watchlist ---");
        user.AddToWatchlist(m1);
        user.AddToWatchlist(m2);
        user.AddToWatchlist(m3);
        Console.WriteLine();

        Console.WriteLine("--- Watchlist ---");
        user.ShowWatchlist();
        Console.WriteLine();

        // Watch movies
        Console.WriteLine("--- Watch Movies ---");
        user.WatchMovie(m1);
        user.WatchMovie(m3);
        Console.WriteLine();

        // Rate movies
        Console.WriteLine("--- Rate Movies ---");
        user.RateMovie(m1, 5);
        user.RateMovie(m3, 4.5);
        user.RateMovie(m2, 4); // should fail — not watched yet
        Console.WriteLine();

        // Show history and updated watchlist
        Console.WriteLine("--- Watch History ---");
        user.ShowWatchHistory();
        Console.WriteLine();

        Console.WriteLine("--- Updated Watchlist ---");
        user.ShowWatchlist();
        Console.WriteLine();

        // Recommendations
        Console.WriteLine("--- Recommendations ---");
        platform.RecommendMovies(user);
    }
}
