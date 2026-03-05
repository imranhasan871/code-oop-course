"""
Practice 08: Movie Streaming Platform (OOAD)
Task: Model a movie streaming platform where users can browse movies,
      add to watchlist, watch movies, and rate them.

How to run:
  python practice_08.py

Key Concepts:
  - Classes with list fields (collections)
  - Methods that manipulate collections
  - Average calculation with running totals
"""


class Movie:
    """Movie represents a movie in the streaming library."""

    def __init__(self, title: str, genre: str, duration: int):
        """Creates a new Movie with no ratings."""
        self.title = title
        self.genre = genre
        self.duration = duration
        self._total_rating = 0.0
        self._rating_count = 0

    def add_rating(self, score: int):
        """Adds a rating score (1-5) to the movie."""
        if score < 1 or score > 5:
            print(f"  [Error] Rating must be between 1 and 5 (given: {score}).")
            return
        self._total_rating += score
        self._rating_count += 1

    def average_rating(self) -> float:
        """Returns the average rating or 0 if no ratings."""
        if self._rating_count == 0:
            return 0
        return self._total_rating / self._rating_count

    def show_info(self):
        """Prints movie details."""
        print(f"  Title    : {self.title}")
        print(f"  Genre    : {self.genre}")
        print(f"  Duration : {self.duration} min")
        if self._rating_count > 0:
            print(f"  Rating   : {self.average_rating():.1f} / 5.0 ({self._rating_count} ratings)")
        else:
            print("  Rating   : No ratings yet")
        print()


class User:
    """User represents a streaming platform user."""

    def __init__(self, username: str):
        """Creates a new User with empty watchlist and history."""
        self.username = username
        self.watchlist = []
        self.watch_history = []

    def add_to_watchlist(self, movie: Movie):
        """Adds a movie to the user's watchlist."""
        if movie in self.watchlist:
            print(f'  [Error] "{movie.title}" is already in {self.username}\'s watchlist.')
            return
        self.watchlist.append(movie)
        print(f'  [OK] "{movie.title}" added to {self.username}\'s watchlist.')

    def watch_movie(self, movie: Movie):
        """Moves a movie from watchlist to watch history."""
        if movie in self.watch_history:
            print(f'  [Error] {self.username} has already watched "{movie.title}".')
            return
        if movie in self.watchlist:
            self.watchlist.remove(movie)
        self.watch_history.append(movie)
        print(f'  [OK] {self.username} watched "{movie.title}".')

    def rate_movie(self, movie: Movie, score: int):
        """Rates a movie the user has watched."""
        if movie not in self.watch_history:
            print(f'  [Error] {self.username} has not watched "{movie.title}" yet. Watch it first to rate.')
            return
        movie.add_rating(score)
        print(f'  [OK] {self.username} rated "{movie.title}" with {score} / 5.')

    def show_info(self):
        """Prints user details including watchlist and history."""
        print(f"  Username      : {self.username}")

        if not self.watchlist:
            print("  Watchlist     : Empty")
        else:
            titles = ", ".join(m.title for m in self.watchlist)
            print(f"  Watchlist     : {titles}")

        if not self.watch_history:
            print("  Watch History : Empty")
        else:
            titles = ", ".join(m.title for m in self.watch_history)
            print(f"  Watch History : {titles}")
        print()


def main():
    # --- Create movie library ---
    movie1 = Movie("The Shawshank Redemption", "Drama", 142)
    movie2 = Movie("Inception", "Sci-Fi", 148)
    movie3 = Movie("The Dark Knight", "Action", 152)

    print("=== Movie Library ===")
    movie1.show_info()
    movie2.show_info()
    movie3.show_info()

    # --- Create users ---
    user1 = User("Tareq")
    user2 = User("Afsana")

    # --- Add to watchlist ---
    print("=== Adding to Watchlist ===")
    user1.add_to_watchlist(movie1)
    user1.add_to_watchlist(movie2)
    user2.add_to_watchlist(movie2)
    user2.add_to_watchlist(movie3)
    print()

    print("=== User Info After Adding Watchlist ===")
    user1.show_info()
    user2.show_info()

    # --- Watch movies ---
    print("=== Watching Movies ===")
    user1.watch_movie(movie1)
    user1.watch_movie(movie2)
    user2.watch_movie(movie2)
    print()

    print("=== User Info After Watching ===")
    user1.show_info()
    user2.show_info()

    # --- Rate movies ---
    print("=== Rating Movies ===")
    user1.rate_movie(movie1, 5)
    user1.rate_movie(movie2, 4)
    user2.rate_movie(movie2, 5)
    # Try rating unwatched movie
    user2.rate_movie(movie1, 3)
    print()

    # --- Final state ---
    print("=== Final Movie Ratings ===")
    movie1.show_info()
    movie2.show_info()
    movie3.show_info()


if __name__ == "__main__":
    main()
