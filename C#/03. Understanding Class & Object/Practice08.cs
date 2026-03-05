/**
 * Practice 08: Movie Streaming Platform (OOAD)
 * Task: Model a movie streaming platform where users can browse movies,
 *       add to watchlist, watch movies, and rate them.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice08.cs
 *
 * Or create a Console App and copy contents into Program.cs:
 *   dotnet new console -n Practice08
 *   dotnet run --project Practice08
 *
 * Key Concepts:
 *   - Classes with List fields (collections)
 *   - Methods that manipulate collections
 *   - Average calculation with running totals
 */

using System;
using System.Collections.Generic;

class Practice08
{
    /** Movie represents a movie in the streaming library. */
    class Movie
    {
        public string Title;
        public string Genre;
        public int Duration;
        private double totalRating;
        private int ratingCount;

        /** Creates a new Movie with no ratings. */
        public Movie(string title, string genre, int duration)
        {
            Title = title;
            Genre = genre;
            Duration = duration;
            totalRating = 0;
            ratingCount = 0;
        }

        /** Adds a rating score (1-5) to the movie. */
        public void AddRating(int score)
        {
            if (score < 1 || score > 5)
            {
                Console.WriteLine($"  [Error] Rating must be between 1 and 5 (given: {score}).");
                return;
            }
            totalRating += score;
            ratingCount++;
        }

        /** Returns the average rating or 0 if no ratings. */
        public double AverageRating()
        {
            if (ratingCount == 0) return 0;
            return totalRating / ratingCount;
        }

        /** Prints movie details. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Title    : {Title}");
            Console.WriteLine($"  Genre    : {Genre}");
            Console.WriteLine($"  Duration : {Duration} min");
            if (ratingCount > 0)
                Console.WriteLine($"  Rating   : {AverageRating():F1} / 5.0 ({ratingCount} ratings)");
            else
                Console.WriteLine("  Rating   : No ratings yet");
            Console.WriteLine();
        }
    }

    /** User represents a streaming platform user. */
    class User
    {
        public string Username;
        private List<Movie> watchlist;
        private List<Movie> watchHistory;

        /** Creates a new User with empty watchlist and history. */
        public User(string username)
        {
            Username = username;
            watchlist = new List<Movie>();
            watchHistory = new List<Movie>();
        }

        /** Adds a movie to the user's watchlist. */
        public void AddToWatchlist(Movie movie)
        {
            if (watchlist.Contains(movie))
            {
                Console.WriteLine($"  [Error] \"{movie.Title}\" is already in {Username}'s watchlist.");
                return;
            }
            watchlist.Add(movie);
            Console.WriteLine($"  [OK] \"{movie.Title}\" added to {Username}'s watchlist.");
        }

        /** Moves a movie from watchlist to watch history. */
        public void WatchMovie(Movie movie)
        {
            if (watchHistory.Contains(movie))
            {
                Console.WriteLine($"  [Error] {Username} has already watched \"{movie.Title}\".");
                return;
            }
            watchlist.Remove(movie);
            watchHistory.Add(movie);
            Console.WriteLine($"  [OK] {Username} watched \"{movie.Title}\".");
        }

        /** Rates a movie the user has watched. */
        public void RateMovie(Movie movie, int score)
        {
            if (!watchHistory.Contains(movie))
            {
                Console.WriteLine($"  [Error] {Username} has not watched \"{movie.Title}\" yet. Watch it first to rate.");
                return;
            }
            movie.AddRating(score);
            Console.WriteLine($"  [OK] {Username} rated \"{movie.Title}\" with {score} / 5.");
        }

        /** Prints user details including watchlist and history. */
        public void ShowInfo()
        {
            Console.WriteLine($"  Username      : {Username}");

            Console.Write("  Watchlist     : ");
            if (watchlist.Count == 0)
            {
                Console.WriteLine("Empty");
            }
            else
            {
                List<string> titles = new List<string>();
                foreach (Movie m in watchlist) titles.Add(m.Title);
                Console.WriteLine(string.Join(", ", titles));
            }

            Console.Write("  Watch History : ");
            if (watchHistory.Count == 0)
            {
                Console.WriteLine("Empty");
            }
            else
            {
                List<string> titles = new List<string>();
                foreach (Movie m in watchHistory) titles.Add(m.Title);
                Console.WriteLine(string.Join(", ", titles));
            }
            Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        // --- Create movie library ---
        Movie movie1 = new Movie("The Shawshank Redemption", "Drama", 142);
        Movie movie2 = new Movie("Inception", "Sci-Fi", 148);
        Movie movie3 = new Movie("The Dark Knight", "Action", 152);

        Console.WriteLine("=== Movie Library ===");
        movie1.ShowInfo();
        movie2.ShowInfo();
        movie3.ShowInfo();

        // --- Create users ---
        User user1 = new User("Tareq");
        User user2 = new User("Afsana");

        // --- Add to watchlist ---
        Console.WriteLine("=== Adding to Watchlist ===");
        user1.AddToWatchlist(movie1);
        user1.AddToWatchlist(movie2);
        user2.AddToWatchlist(movie2);
        user2.AddToWatchlist(movie3);
        Console.WriteLine();

        Console.WriteLine("=== User Info After Adding Watchlist ===");
        user1.ShowInfo();
        user2.ShowInfo();

        // --- Watch movies ---
        Console.WriteLine("=== Watching Movies ===");
        user1.WatchMovie(movie1);
        user1.WatchMovie(movie2);
        user2.WatchMovie(movie2);
        Console.WriteLine();

        Console.WriteLine("=== User Info After Watching ===");
        user1.ShowInfo();
        user2.ShowInfo();

        // --- Rate movies ---
        Console.WriteLine("=== Rating Movies ===");
        user1.RateMovie(movie1, 5);
        user1.RateMovie(movie2, 4);
        user2.RateMovie(movie2, 5);
        // Try rating unwatched movie
        user2.RateMovie(movie1, 3);
        Console.WriteLine();

        // --- Final state ---
        Console.WriteLine("=== Final Movie Ratings ===");
        movie1.ShowInfo();
        movie2.ShowInfo();
        movie3.ShowInfo();
    }
}
