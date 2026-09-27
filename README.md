# B3 UX Backend

[About](#about-this-project) · [Installation](#installation) · [API](#api) · [Credits](#credits)

> **Educational project**
>
> This backend was created to support frontend development exercises and student projects. It is **not intended to be a production-ready application**.

## About this project

This project is an **educational backend designed to support frontend development exercises and projects**.

Its purpose is to provide students with a ready-to-use API so they can focus on building frontend applications without having to implement or maintain the backend themselves.

This is **not a production-ready application**. Some technical choices, validations, features, and security considerations have been intentionally simplified to keep the project focused on its educational purpose.

The API may evolve during the course. Its main goal is to provide a stable and understandable interface for frontend development.

---

## Installation

There are two ways to run the backend locally:

- **Compile the project yourself** from the source code.
- **Download the precompiled Windows executable** from the GitHub Releases section.

For most students, using the executable is the simplest option.

### Using the Windows executable

Download the latest `.exe` file from the project's **Releases** page and run it.

The application automatically creates the files it needs inside your Windows Documents folder:

```text
C:\Users\<username>\Documents\b3_ux
```

You do not need to manually create this directory.

### Local data

The `b3_ux` directory contains the local resources used by the backend:

```text
b3_ux/
├── data/
├── images/
└── database.db
```

#### `data/`

Contains the JSON files used to generate and populate the database.

These files provide the initial data used by the different student projects.

#### `images/`

Contains the images used by the applications, including AI-generated illustrations for movies and apps.

These files are served by the backend through the image routes.

#### `database.db`

The SQLite database used by the backend.

It contains the application's current local data, including data created or modified while using the API.

Because the database is stored locally, each student works with their own independent copy of the application data.

---

# API

[Images](#images) · [Authentication](#authentication) · [Store](#store) · [Wyww](#wyww)

## API organization

The backend is shared by **two frontend projects**. Both projects use the same API server, but their routes and data are kept separate.

The API can be divided into three main parts:

| Group | Purpose |
| --- | --- |
| **Auth** | Registration, login and authentication shared by both projects |
| **Store** | Routes used by the fake app store project |
| **Wyww** | Routes used by the movie recommendation project |

You only need to use the part of the API relevant to the project you are currently working on, together with the authentication routes when required.

### Store

`Store` contains the routes used by the **fake app store project**.

Its routes are grouped under:

```text
/store/...
```

### Wyww

`Wyww` contains the routes used by the **movie recommendation project**.

Its routes are grouped under:

```text
/wyww/...
```

### Authentication

Authentication is **shared by both projects**.

Its routes are grouped under:

```text
/auth/...
```

---

## Images

[Authentication](#authentication) · [Store](#store) · [Wyww](#wyww)

| Method | Route | Description |
| --- | --- | --- |
| `GET` | `/images/:id` | Returns an image from its ID |

### `GET /images/:id`

Returns an image file.

The image is identified by its ID, provided directly in the URL.

**Path parameters**

```ts
{
    id: string;
}
```

**Response**

An image file.

---

## Authentication

[Images](#images) · [Store](#store) · [Wyww](#wyww)

Authentication is shared between the `Store` and `Wyww` projects.

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/auth/signin` | Creates a new user |
| `POST` | `/auth/login` | Logs a user in |
| `POST` | `/auth/logout` | Logs the current user out |
| `POST` | `/auth/me` | Restores a session from authentication cookies |

### `POST /auth/signin`

Creates a new user and their authentication credentials.

**Request body**

```ts
{
    username: string;
    email: string;
    password: string;
    is_admin: boolean;
}
```

**Response**

```ts
UserData
```

### `POST /auth/login`

Authenticates a user from their username and password.

**Request body**

```ts
{
    username: string;
    password: string;
}
```

**Response**

```ts
UserData
```

The backend creates the authentication session required for subsequent authenticated requests.

### `POST /auth/logout`

Logs out the currently authenticated user.

Authentication is determined from the current session.

No request body is required.

### `POST /auth/me`

Attempts to authenticate the user using authentication cookies already stored by the browser.

This route can be used when loading the frontend application to determine whether the user already has an active session.

No request body is required.

---

## Store

[Apps](#apps) · [Store genres](#store-genres) · [Authentication](#authentication) · [Wyww](#wyww)

These routes provide access to the entities used by the fake app store project.

### Route overview

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/store/app/findOne` | Finds a single app |
| `POST` | `/store/app/findMany` | Finds multiple apps |
| `POST` | `/store/genre/findOne` | Finds a single genre |
| `POST` | `/store/genre/findMany` | Finds multiple genres |

### Apps

#### `POST /store/app/findOne`

Finds a single app matching the supplied filters.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
}
```

**Response**

```ts
StoreAppData
```

#### `POST /store/app/findMany`

Finds multiple apps matching the supplied filters.

Sorting and pagination can optionally be applied.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
    pagination?: PaginationOptions;
}
```

**Response**

```ts
StoreAppData[]
```

### Store genres

#### `POST /store/genre/findOne`

Finds a single store genre matching the supplied filters.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
}
```

**Response**

```ts
StoreGenreData
```

#### `POST /store/genre/findMany`

Finds multiple store genres matching the supplied filters.

Sorting and pagination can optionally be applied.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
    pagination?: PaginationOptions;
}
```

**Response**

```ts
StoreGenreData[]
```

---

## Wyww

[Recommendations](#recommendations) · [User movie lists](#user-movie-lists) · [Movies](#movies) · [Wyww genres](#wyww-genres)

These routes provide the data and operations required by the movie recommendation project.

### Route overview

| Method | Route | Description |
| --- | --- | --- |
| `GET` | `/wyww/movie/newRecommendation` | Generates a new recommendation set |
| `GET` | `/wyww/movie/swapRecommendation` | Changes the main recommendation |
| `GET` | `/wyww/movie/watchlist` | Gets the user's state for a movie |
| `PATCH` | `/wyww/movie/toggleWatched` | Toggles the watched state |
| `PATCH` | `/wyww/movie/toggleRecommended` | Toggles the recommended state |
| `POST` | `/wyww/movie/findOne` | Finds a single movie |
| `POST` | `/wyww/movie/findMany` | Finds multiple movies |
| `POST` | `/wyww/genre/findOne` | Finds a single genre |
| `POST` | `/wyww/genre/findMany` | Finds multiple genres |

### Recommendations

#### `GET /wyww/movie/newRecommendation`

Generates a new set of movie recommendations for the current user.

The backend selects a **main recommendation** semi-randomly based on the user's data. Additional recommendations are grouped by genre.

No request parameters are required.

**Response**

```ts
{
    main: WywwMovieData;
    genres: {
        [key: string]: WywwMovieData[];
    };
}
```

The keys inside `genres` represent genres, and each value contains the movies recommended for that genre.

#### `GET /wyww/movie/swapRecommendation`

Replaces the current main recommendation with a selected movie.

The supplied movie becomes the new main recommendation, and the backend generates a new set of secondary recommendations around it.

**Query parameters**

```ts
{
    id: string;
}
```

**Response**

```ts
{
    main: WywwMovieData;
    genres: {
        [key: string]: WywwMovieData[];
    };
}
```

### User movie lists

#### `GET /wyww/movie/watchlist`

Returns the current user's status for a particular movie.

**Query parameters**

```ts
{
    id: string;
}
```

**Response**

```ts
{
    watched: boolean;
    recommended: boolean;
}
```

- `watched` indicates whether the movie has been marked as watched.
- `recommended` indicates whether the user has added the movie to their recommendations.

#### `PATCH /wyww/movie/toggleWatched`

Toggles the watched state of a movie for the current user.

If the movie is not currently marked as watched, it is added to the user's watched movies. If it is already marked as watched, it is removed.

**Request body**

```ts
{
    id: string;
}
```

**Response**

```ts
{
    watched: boolean;
}
```

The returned value represents the new state after the operation.

#### `PATCH /wyww/movie/toggleRecommended`

Toggles the recommended state of a movie for the current user.

If the movie is not currently recommended, it is added to the user's recommended movies. If it is already recommended, it is removed.

**Request body**

```ts
{
    id: string;
}
```

**Response**

```ts
{
    recommended: boolean;
}
```

The returned value represents the new state after the operation.

### Movies

#### `POST /wyww/movie/findOne`

Finds a single movie matching the supplied filters.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
}
```

**Response**

```ts
WywwMovieData
```

#### `POST /wyww/movie/findMany`

Finds multiple movies matching the supplied filters.

Sorting and pagination can optionally be applied.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
    pagination?: PaginationOptions;
}
```

**Response**

```ts
WywwMovieData[]
```

### Wyww genres

#### `POST /wyww/genre/findOne`

Finds a single movie genre matching the supplied filters.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
}
```

**Response**

```ts
WywwGenreData
```

#### `POST /wyww/genre/findMany`

Finds multiple movie genres matching the supplied filters.

Sorting and pagination can optionally be applied.

**Request body**

```ts
{
    filters: FilterItem[];
    sorting?: SortOptions;
    pagination?: PaginationOptions;
}
```

**Response**

```ts
WywwGenreData[]
```

---

## Credits

### Built with

- Go
- Ent
- Gin
- Google UUID
- errors
- Zap
- `x/image`
- `modernc.org/sqlite`

### AI-generated assets

Images used to illustrate movies and applications were generated with:

- **FLUX.1 [dev]** — Black Forest Labs
- **ChatGPT** — OpenAI

### Third-party licenses

This project relies on third-party open-source libraries distributed under their respective licenses.

See [`THIRD_PARTY_LICENSES.md`](./THIRD_PARTY_LICENSES.md) for details.