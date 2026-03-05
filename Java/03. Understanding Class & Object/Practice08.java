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
 *   - Multiple inner classes collaborating
 *   - ArrayList for collections
 *   - Encapsulation with private fields
 *   - Constructor and method design
 *   - Business logic (recommendations based on genre)
 *
 * Course: Professional OOP — by Zohirul Alam Tiemoon
 */

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Practice08 {

    /** Represents a film in the library. */
    static class Movie {
        private String movieID;
        private String title;
        private String genre;
        private double totalRating;
        private int ratingCount;

        public Movie(String movieID, String title, String genre) {
            this.movieID = movieID;
            this.title = title;
            this.genre = genre;
            this.totalRating = 0;
            this.ratingCount = 0;
        }

        public String getMovieID() { return movieID; }
        public String getTitle() { return title; }
        public String getGenre() { return genre; }

        public double getAverageRating() {
            return ratingCount == 0 ? 0 : totalRating / ratingCount;
        }

        public void addRating(double rating) {
            totalRating += rating;
            ratingCount++;
        }

        public void printInfo() {
            System.out.printf("  [%s] %s | Genre: %s | Avg Rating: %.1f (%d ratings)%n",
                    movieID, title, genre, getAverageRating(), ratingCount);
        }
    }

    /** Represents a platform user. */
    static class User {
        private String name;
        private List<Movie> watchlist;
        private List<Movie> watchHistory;

        public User(String name) {
            this.name = name;
            this.watchlist = new ArrayList<>();
            this.watchHistory = new ArrayList<>();
        }

        public String getName() { return name; }
        public List<Movie> getWatchHistory() { return watchHistory; }

        public void addToWatchlist(Movie movie) {
            watchlist.add(movie);
            System.out.printf("  [OK] %s added '%s' to watchlist.%n", name, movie.getTitle());
        }

        public void watchMovie(Movie movie) {
            watchlist.removeIf(m -> m.getMovieID().equals(movie.getMovieID()));
            watchHistory.add(movie);
            System.out.printf("  [OK] %s watched '%s'.%n", name, movie.getTitle());
        }

        public void rateMovie(Movie movie, double rating) {
            boolean watched = false;
            for (Movie m : watchHistory) {
                if (m.getMovieID().equals(movie.getMovieID())) {
                    watched = true;
                    break;
                }
            }
            if (!watched) {
                System.out.printf("  [Error] %s hasn't watched '%s' yet.%n", name, movie.getTitle());
                return;
            }
            if (rating < 1 || rating > 5) {
                System.out.println("  [Error] Rating must be between 1 and 5.");
                return;
            }
            movie.addRating(rating);
            System.out.printf("  [OK] %s rated '%s' — %.1f/5%n", name, movie.getTitle(), rating);
        }

        public void showWatchlist() {
            System.out.printf("  %s's Watchlist (%d movies):%n", name, watchlist.size());
            for (Movie m : watchlist) {
                System.out.printf("    - %s (%s)%n", m.getTitle(), m.getGenre());
            }
            if (watchlist.isEmpty()) {
                System.out.println("    (empty)");
            }
        }

        public void showWatchHistory() {
            System.out.printf("  %s's Watch History (%d movies):%n", name, watchHistory.size());
            for (Movie m : watchHistory) {
                System.out.printf("    - %s (%s)%n", m.getTitle(), m.getGenre());
            }
            if (watchHistory.isEmpty()) {
                System.out.println("    (empty)");
            }
        }
    }

    /** Manages the movie library and users. */
    static class StreamingPlatform {
        private String name;
        private List<Movie> movies;

        public StreamingPlatform(String name) {
            this.name = name;
            this.movies = new ArrayList<>();
        }

        public void addMovie(Movie movie) {
            movies.add(movie);
            System.out.printf("  [OK] Added '%s' to %s library.%n", movie.getTitle(), name);
        }

        public void browseMovies() {
            System.out.printf("  %s Library (%d movies):%n", name, movies.size());
            for (Movie m : movies) {
                m.printInfo();
            }
        }

        public void recommendMovies(User user) {
            Set<String> genres = new HashSet<>();
            Set<String> watchedIDs = new HashSet<>();
            for (Movie m : user.getWatchHistory()) {
                genres.add(m.getGenre());
                watchedIDs.add(m.getMovieID());
            }

            System.out.printf("  Recommendations for %s:%n", user.getName());
            boolean found = false;
            for (Movie m : movies) {
                if (genres.contains(m.getGenre()) && !watchedIDs.contains(m.getMovieID())) {
                    System.out.printf("    - %s (%s) — Avg Rating: %.1f%n",
                            m.getTitle(), m.getGenre(), m.getAverageRating());
                    found = true;
                }
            }
            if (!found) {
                System.out.println("    (no recommendations available)");
            }
        }
    }

    public static void main(String[] args) {
        System.out.println("=== Practice 08: Movie Streaming Platform ===");
        System.out.println();

        // Create platform
        StreamingPlatform platform = new StreamingPlatform("CineStream");

        // Add movies
        System.out.println("--- Add Movies ---");
        Movie m1 = new Movie("M-001", "The Matrix", "Sci-Fi");
        Movie m2 = new Movie("M-002", "Interstellar", "Sci-Fi");
        Movie m3 = new Movie("M-003", "The Dark Knight", "Action");
        Movie m4 = new Movie("M-004", "Inception", "Sci-Fi");
        Movie m5 = new Movie("M-005", "John Wick", "Action");

        platform.addMovie(m1);
        platform.addMovie(m2);
        platform.addMovie(m3);
        platform.addMovie(m4);
        platform.addMovie(m5);
        System.out.println();

        // Browse library
        System.out.println("--- Browse Movies ---");
        platform.browseMovies();
        System.out.println();

        // Create user
        User user = new User("Imtiaz");

        // Add to watchlist
        System.out.println("--- Add to Watchlist ---");
        user.addToWatchlist(m1);
        user.addToWatchlist(m2);
        user.addToWatchlist(m3);
        System.out.println();

        System.out.println("--- Watchlist ---");
        user.showWatchlist();
        System.out.println();

        // Watch movies
        System.out.println("--- Watch Movies ---");
        user.watchMovie(m1);
        user.watchMovie(m3);
        System.out.println();

        // Rate movies
        System.out.println("--- Rate Movies ---");
        user.rateMovie(m1, 5);
        user.rateMovie(m3, 4.5);
        user.rateMovie(m2, 4); // should fail — not watched yet
        System.out.println();

        // Show history and updated watchlist
        System.out.println("--- Watch History ---");
        user.showWatchHistory();
        System.out.println();

        System.out.println("--- Updated Watchlist ---");
        user.showWatchlist();
        System.out.println();

        // Recommendations
        System.out.println("--- Recommendations ---");
        platform.recommendMovies(user);
    }
}
