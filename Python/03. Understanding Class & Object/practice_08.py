"""
Practice 08: Movie Streaming Platform (OOAD)

Scenario: An online movie streaming platform offers a large library of
films. Users can browse the available movies, add them to their watchlist,
and rate the ones they've seen. The system keeps track of each user's
watch history and recommends similar films based on their preferences.

Main Objects:
  - Movie    — knows: id, title, genre, rating
  - User     — knows: name; does: add to watchlist, watch, rate
  - Platform — knows: movie library, users; does: add movie, recommend

Key Concepts:
  - Multiple classes collaborating
  - Lists and sets for collections
  - Encapsulation with private attributes
  - Constructor design (__init__)
  - Business logic (recommendations based on genre)

Course: Professional OOP — by Zohirul Alam Tiemoon
"""


class Movie:
    """Represents a film in the library."""

    def __init__(self, movie_id: str, title: str, genre: str):
        self._movie_id = movie_id
        self._title = title
        self._genre = genre
        self._total_rating = 0.0
        self._rating_count = 0

    @property
    def movie_id(self) -> str:
        return self._movie_id

    @property
    def title(self) -> str:
        return self._title

    @property
    def genre(self) -> str:
        return self._genre

    def get_average_rating(self) -> float:
        return 0 if self._rating_count == 0 else self._total_rating / self._rating_count

    def add_rating(self, rating: float):
        self._total_rating += rating
        self._rating_count += 1

    def print_info(self):
        print(f"  [{self._movie_id}] {self._title} | Genre: {self._genre} | "
              f"Avg Rating: {self.get_average_rating():.1f} ({self._rating_count} ratings)")


class User:
    """Represents a platform user."""

    def __init__(self, name: str):
        self._name = name
        self._watchlist = []
        self._watch_history = []

    @property
    def name(self) -> str:
        return self._name

    @property
    def watch_history(self) -> list:
        return self._watch_history

    def add_to_watchlist(self, movie: Movie):
        self._watchlist.append(movie)
        print(f"  [OK] {self._name} added '{movie.title}' to watchlist.")

    def watch_movie(self, movie: Movie):
        self._watchlist = [m for m in self._watchlist if m.movie_id != movie.movie_id]
        self._watch_history.append(movie)
        print(f"  [OK] {self._name} watched '{movie.title}'.")

    def rate_movie(self, movie: Movie, rating: float):
        watched = any(m.movie_id == movie.movie_id for m in self._watch_history)
        if not watched:
            print(f"  [Error] {self._name} hasn't watched '{movie.title}' yet.")
            return
        if rating < 1 or rating > 5:
            print("  [Error] Rating must be between 1 and 5.")
            return
        movie.add_rating(rating)
        print(f"  [OK] {self._name} rated '{movie.title}' — {rating:.1f}/5")

    def show_watchlist(self):
        print(f"  {self._name}'s Watchlist ({len(self._watchlist)} movies):")
        for m in self._watchlist:
            print(f"    - {m.title} ({m.genre})")
        if not self._watchlist:
            print("    (empty)")

    def show_watch_history(self):
        print(f"  {self._name}'s Watch History ({len(self._watch_history)} movies):")
        for m in self._watch_history:
            print(f"    - {m.title} ({m.genre})")
        if not self._watch_history:
            print("    (empty)")


class StreamingPlatform:
    """Manages the movie library and users."""

    def __init__(self, name: str):
        self._name = name
        self._movies = []

    def add_movie(self, movie: Movie):
        self._movies.append(movie)
        print(f"  [OK] Added '{movie.title}' to {self._name} library.")

    def browse_movies(self):
        print(f"  {self._name} Library ({len(self._movies)} movies):")
        for m in self._movies:
            m.print_info()

    def recommend_movies(self, user: User):
        genres = {m.genre for m in user.watch_history}
        watched_ids = {m.movie_id for m in user.watch_history}

        print(f"  Recommendations for {user.name}:")
        found = False
        for m in self._movies:
            if m.genre in genres and m.movie_id not in watched_ids:
                print(f"    - {m.title} ({m.genre}) — "
                      f"Avg Rating: {m.get_average_rating():.1f}")
                found = True
        if not found:
            print("    (no recommendations available)")


if __name__ == "__main__":
    print("=== Practice 08: Movie Streaming Platform ===")
    print()

    # Create platform
    platform = StreamingPlatform("CineStream")

    # Add movies
    print("--- Add Movies ---")
    m1 = Movie("M-001", "The Matrix", "Sci-Fi")
    m2 = Movie("M-002", "Interstellar", "Sci-Fi")
    m3 = Movie("M-003", "The Dark Knight", "Action")
    m4 = Movie("M-004", "Inception", "Sci-Fi")
    m5 = Movie("M-005", "John Wick", "Action")

    platform.add_movie(m1)
    platform.add_movie(m2)
    platform.add_movie(m3)
    platform.add_movie(m4)
    platform.add_movie(m5)
    print()

    # Browse library
    print("--- Browse Movies ---")
    platform.browse_movies()
    print()

    # Create user
    user = User("Imtiaz")

    # Add to watchlist
    print("--- Add to Watchlist ---")
    user.add_to_watchlist(m1)
    user.add_to_watchlist(m2)
    user.add_to_watchlist(m3)
    print()

    print("--- Watchlist ---")
    user.show_watchlist()
    print()

    # Watch movies
    print("--- Watch Movies ---")
    user.watch_movie(m1)
    user.watch_movie(m3)
    print()

    # Rate movies
    print("--- Rate Movies ---")
    user.rate_movie(m1, 5)
    user.rate_movie(m3, 4.5)
    user.rate_movie(m2, 4)  # should fail — not watched yet
    print()

    # Show history and updated watchlist
    print("--- Watch History ---")
    user.show_watch_history()
    print()

    print("--- Updated Watchlist ---")
    user.show_watchlist()
    print()

    # Recommendations
    print("--- Recommendations ---")
    platform.recommend_movies(user)
