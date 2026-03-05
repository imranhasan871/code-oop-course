/**
 * Practice 08: Movie Streaming Platform (OOAD)
 * Task: Model a movie streaming platform where users can browse movies,
 *       add to watchlist, watch movies, and rate them.
 *
 * How to compile and run:
 *   javac Practice08.java
 *   java Practice08
 *
 * Key Concepts:
 *   - Classes with ArrayList fields (collections)
 *   - Methods that manipulate collections
 *   - Average calculation with running totals
 */

import java.util.ArrayList;
import java.util.List;

public class Practice08 {

    /** Movie represents a movie in the streaming library. */
    static class Movie {
        private String title;
        private String genre;
        private int duration;
        private double totalRating;
        private int ratingCount;

        /** Creates a new Movie with no ratings. */
        public Movie(String title, String genre, int duration) {
            this.title = title;
            this.genre = genre;
            this.duration = duration;
            this.totalRating = 0;
            this.ratingCount = 0;
        }

        /** Adds a rating score (1-5) to the movie. */
        public void addRating(int score) {
            if (score < 1 || score > 5) {
                System.out.printf("  [Error] Rating must be between 1 and 5 (given: %d).%n", score);
                return;
            }
            totalRating += score;
            ratingCount++;
        }

        /** Returns the average rating or 0 if no ratings. */
        public double averageRating() {
            if (ratingCount == 0) return 0;
            return totalRating / ratingCount;
        }

        /** Prints movie details. */
        public void showInfo() {
            System.out.printf("  Title    : %s%n", title);
            System.out.printf("  Genre    : %s%n", genre);
            System.out.printf("  Duration : %d min%n", duration);
            if (ratingCount > 0) {
                System.out.printf("  Rating   : %.1f / 5.0 (%d ratings)%n", averageRating(), ratingCount);
            } else {
                System.out.printf("  Rating   : No ratings yet%n");
            }
            System.out.println();
        }
    }

    /** User represents a streaming platform user. */
    static class User {
        private String username;
        private List<Movie> watchlist;
        private List<Movie> watchHistory;

        /** Creates a new User with empty watchlist and history. */
        public User(String username) {
            this.username = username;
            this.watchlist = new ArrayList<>();
            this.watchHistory = new ArrayList<>();
        }

        /** Adds a movie to the user's watchlist. */
        public void addToWatchlist(Movie movie) {
            if (watchlist.contains(movie)) {
                System.out.printf("  [Error] \"%s\" is already in %s's watchlist.%n", movie.title, username);
                return;
            }
            watchlist.add(movie);
            System.out.printf("  [OK] \"%s\" added to %s's watchlist.%n", movie.title, username);
        }

        /** Moves a movie from watchlist to watch history. */
        public void watchMovie(Movie movie) {
            if (watchHistory.contains(movie)) {
                System.out.printf("  [Error] %s has already watched \"%s\".%n", username, movie.title);
                return;
            }
            watchlist.remove(movie);
            watchHistory.add(movie);
            System.out.printf("  [OK] %s watched \"%s\".%n", username, movie.title);
        }

        /** Rates a movie the user has watched. */
        public void rateMovie(Movie movie, int score) {
            if (!watchHistory.contains(movie)) {
                System.out.printf("  [Error] %s has not watched \"%s\" yet. Watch it first to rate.%n",
                        username, movie.title);
                return;
            }
            movie.addRating(score);
            System.out.printf("  [OK] %s rated \"%s\" with %d / 5.%n", username, movie.title, score);
        }

        /** Prints user details including watchlist and history. */
        public void showInfo() {
            System.out.printf("  Username      : %s%n", username);

            System.out.print("  Watchlist     : ");
            if (watchlist.isEmpty()) {
                System.out.println("Empty");
            } else {
                List<String> titles = new ArrayList<>();
                for (Movie m : watchlist) titles.add(m.title);
                System.out.println(String.join(", ", titles));
            }

            System.out.print("  Watch History : ");
            if (watchHistory.isEmpty()) {
                System.out.println("Empty");
            } else {
                List<String> titles = new ArrayList<>();
                for (Movie m : watchHistory) titles.add(m.title);
                System.out.println(String.join(", ", titles));
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // --- Create movie library ---
        Movie movie1 = new Movie("The Shawshank Redemption", "Drama", 142);
        Movie movie2 = new Movie("Inception", "Sci-Fi", 148);
        Movie movie3 = new Movie("The Dark Knight", "Action", 152);

        System.out.println("=== Movie Library ===");
        movie1.showInfo();
        movie2.showInfo();
        movie3.showInfo();

        // --- Create users ---
        User user1 = new User("Tareq");
        User user2 = new User("Afsana");

        // --- Add to watchlist ---
        System.out.println("=== Adding to Watchlist ===");
        user1.addToWatchlist(movie1);
        user1.addToWatchlist(movie2);
        user2.addToWatchlist(movie2);
        user2.addToWatchlist(movie3);
        System.out.println();

        System.out.println("=== User Info After Adding Watchlist ===");
        user1.showInfo();
        user2.showInfo();

        // --- Watch movies ---
        System.out.println("=== Watching Movies ===");
        user1.watchMovie(movie1);
        user1.watchMovie(movie2);
        user2.watchMovie(movie2);
        System.out.println();

        System.out.println("=== User Info After Watching ===");
        user1.showInfo();
        user2.showInfo();

        // --- Rate movies ---
        System.out.println("=== Rating Movies ===");
        user1.rateMovie(movie1, 5);
        user1.rateMovie(movie2, 4);
        user2.rateMovie(movie2, 5);
        // Try rating unwatched movie
        user2.rateMovie(movie1, 3);
        System.out.println();

        // --- Final state ---
        System.out.println("=== Final Movie Ratings ===");
        movie1.showInfo();
        movie2.showInfo();
        movie3.showInfo();
    }
}
