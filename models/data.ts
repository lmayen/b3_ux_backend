
export type ImageKindEnum = "poster" | "banner" | "icon" | "";

export type StoreAppCategoryEnum = "ARTIFICIAL_INTELLIGENCE" | "ART_AND_DESIGN" | "AUTO_AND_VEHICLES" | "BEAUTY" | "BOOKS_AND_REFERENCE" | "BUSINESS" | "COMICS" | "COMMUNICATION" | "DATING" | "EDUCATION" | "ENTERTAINMENT" | "EVENTS" | "FAMILY" | "FINANCE" | "FOOD_AND_DRINK" | "GAME" | "HEALTH_AND_FITNESS" | "HOUSE_AND_HOME" | "LIBRARIES_AND_DEMO" | "LIFESTYLE" | "MAPS_AND_NAVIGATION" | "MEDICAL" | "NEWS_AND_MAGAZINES" | "PARENTING" | "PERSONALIZATION" | "PHOTOGRAPHY" | "PRODUCTIVITY" | "SHOPPING" | "SOCIAL" | "SPORTS" | "TOOLS" | "TRAVEL_AND_LOCAL" | "VIDEO_PLAYERS" | "VIRTUAL_REALITY" | "WEATHER" | "WEB3_AND_CRYPTO" | "";

export type StoreAppTypeEnum = "Free" | "Paid" | "";

export type WywwMovieContentRatingEnum = "Approved" | "G" | "NC-17" | "Not Rated" | "PG" | "PG-13" | "Passed" | "R" | "";

export interface ColorData {
    id: string,r: number,
    g: number,
    b: number,
    ratio: number,
    }

export interface ImageData {
    id: string,filename: string,
    path: string,
    width: number,
    height: number,
    kind: ImageKindEnum,
    colors?: ColorData[],
    }

export interface SessionData {
    id: string,token: string,
    created_at: string,
    expires_at: string,
    }

export interface StoreAppData {
    id: string,name: string,
    category: StoreAppCategoryEnum,
    rating: number|null,
    reviews: number,
    size: string,
    installs: string,
    type: StoreAppTypeEnum,
    price: number,
    content_rating: string,
    last_updated: string,
    current_ver: string,
    in_app_purchases: boolean,
    ad_supported: boolean,
    genres?: StoreGenreData[],
    images?: ImageData[],
    }

export interface StoreGenreData {
    id: string,name: string,
    }

export interface UserData {
    id: string,username: string,
    is_admin: boolean,
    email: string,
    password: string,
    recommended_movies?: WywwMovieData[],
    watched_movies?: WywwMovieData[],
    installed_apps?: StoreAppData[],
    sessions?: SessionData[],
    }

export interface WywwGenreData {
    id: string,name: string,
    }

export interface WywwMovieData {
    id: string,movie_title: string,
    original_title: string,
    content_rating: WywwMovieContentRatingEnum,
    description: string,
    released_year: number,
    runtime: number,
    rating: number,
    genres?: WywwGenreData[],
    images?: ImageData[],
    }


