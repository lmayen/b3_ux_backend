# Schema

## Color

**Table:** `colors`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `r` | `int` | `number` | — | — | — | — | ✓ |
| `g` | `int` | `number` | — | — | — | — | ✓ |
| `b` | `int` | `number` | — | — | — | — | ✓ |
| `ratio` | `float32` | `number` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `image` | `Image` | `Image` | M2O | ✓ | ✓ | — | ✓ |

---

## Image

**Table:** `images`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `filename` | `string` | `string` | — | ✓ | — | — | ✓ |
| `path` | `string` | `string` | — | ✓ | — | — | ✓ |
| `width` | `int` | `number` | — | — | — | — | ✓ |
| `height` | `int` | `number` | — | — | — | — | ✓ |
| `kind` | `image.Kind` | `Kind` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `colors` | `Color` | `Colors` | O2M | ✓ | — | — | — |
| `movie` | `WywwMovie` | `Movie` | M2O | ✓ | ✓ | — | ✓ |
| `store` | `StoreApp` | `Store` | M2O | ✓ | ✓ | — | ✓ |

---

## Session

**Table:** `sessions`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `token` | `string` | `string` | — | ✓ | — | ✓ | ✓ |
| `created_at` | `time.Time` | `string` | — | — | — | — | ✓ |
| `expires_at` | `time.Time` | `string` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `artist` | `User` | `Artist` | M2O | — | ✓ | ✓ | ✓ |

---

## StoreApp

**Table:** `store_apps`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `name` | `string` | `string` | — | — | — | — | ✓ |
| `category` | `storeapp.Category` | `Category` | — | — | — | — | ✓ |
| `rating` | `float32` | `number` | ✓ | — | — | — | ✓ |
| `reviews` | `int32` | `number` | — | — | — | — | ✓ |
| `size` | `string` | `string` | — | — | — | — | ✓ |
| `installs` | `string` | `string` | — | — | — | — | ✓ |
| `type` | `storeapp.Type` | `Type` | — | — | — | — | ✓ |
| `price` | `float32` | `number` | — | — | — | — | ✓ |
| `content_rating` | `string` | `string` | — | — | — | — | ✓ |
| `last_updated` | `string` | `string` | — | — | — | — | ✓ |
| `current_ver` | `string` | `string` | — | — | — | — | ✓ |
| `in_app_purchases` | `bool` | `boolean` | — | — | — | — | ✓ |
| `ad_supported` | `bool` | `boolean` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `genres` | `StoreGenre` | `Genres` | M2M | ✓ | — | — | — |
| `images` | `Image` | `Images` | O2M | ✓ | — | — | — |
| `user_install_list` | `User` | `UserInstallList` | M2M | ✓ | — | — | ✓ |

---

## StoreGenre

**Table:** `store_genres`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `name` | `string` | `string` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `app` | `StoreApp` | `App` | M2M | ✓ | — | — | ✓ |

---

## User

**Table:** `users`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `username` | `string` | `string` | — | ✓ | — | — | ✓ |
| `is_admin` | `bool` | `boolean` | — | — | — | — | ✓ |
| `email` | `string` | `string` | — | ✓ | — | — | ✓ |
| `password` | `string` | `string` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `recommended_movies` | `WywwMovie` | `RecommendedMovies` | M2M | ✓ | — | — | — |
| `watched_movies` | `WywwMovie` | `WatchedMovies` | M2M | ✓ | — | — | — |
| `installed_apps` | `StoreApp` | `InstalledApps` | M2M | ✓ | — | — | — |
| `sessions` | `Session` | `Sessions` | O2M | ✓ | — | — | — |

---

## WywwGenre

**Table:** `wyww_genres`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `name` | `string` | `string` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `app` | `WywwMovie` | `App` | M2M | ✓ | — | — | ✓ |

---

## WywwMovie

**Table:** `wyww_movies`

### Fields

| Field | Go | TypeScript | Nillable | Unique | Optional | Immutable | Public |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `id` | `uuid.UUID` | `string` | — | ✓ | — | ✓ | ✓ |
| `movie_title` | `string` | `string` | — | — | — | — | ✓ |
| `original_title` | `string` | `string` | — | — | — | — | ✓ |
| `content_rating` | `wywwmovie.ContentRating` | `ContentRating` | — | — | — | — | ✓ |
| `description` | `string` | `string` | — | — | — | — | ✓ |
| `released_year` | `int32` | `number` | — | — | — | — | ✓ |
| `runtime` | `int32` | `number` | — | — | — | — | ✓ |
| `rating` | `float32` | `number` | — | — | — | — | ✓ |

### Edges

| Edge | Target | Go field | Relation | Optional | Unique | Immutable | Inverse |
|---|---|---|---|:---:|:---:|:---:|:---:|
| `genres` | `WywwGenre` | `Genres` | M2M | ✓ | — | — | — |
| `images` | `Image` | `Images` | O2M | ✓ | — | — | — |
| `user_recommendations` | `User` | `UserRecommendations` | M2M | ✓ | — | — | ✓ |
| `user_watchlist` | `User` | `UserWatchlist` | M2M | ✓ | — | — | ✓ |

---

