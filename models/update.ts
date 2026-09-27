
import { ImageKindEnum } from "./data";

import { StoreAppCategoryEnum } from "./data";

import { StoreAppTypeEnum } from "./data";

import { WywwMovieContentRatingEnum } from "./data";


export interface UpdateField<T> {
	present: boolean,
	value:   T,
}


export interface ColorUpdateParameters {
    r?: UpdateField<number>,
    g?: UpdateField<number>,
    b?: UpdateField<number>,
    ratio?: UpdateField<number>,
    }

export interface ImageUpdateParameters {
    filename?: UpdateField<string>,
    path?: UpdateField<string>,
    width?: UpdateField<number>,
    height?: UpdateField<number>,
    kind?: UpdateField<ImageKindEnum>,
    }

export interface SessionUpdateParameters {
    created_at?: UpdateField<string>,
    expires_at?: UpdateField<string>,
    }

export interface StoreAppUpdateParameters {
    name?: UpdateField<string>,
    category?: UpdateField<StoreAppCategoryEnum>,
    rating?: UpdateField<number|null>,
    reviews?: UpdateField<number>,
    size?: UpdateField<string>,
    installs?: UpdateField<string>,
    type?: UpdateField<StoreAppTypeEnum>,
    price?: UpdateField<number>,
    content_rating?: UpdateField<string>,
    last_updated?: UpdateField<string>,
    current_ver?: UpdateField<string>,
    in_app_purchases?: UpdateField<boolean>,
    ad_supported?: UpdateField<boolean>,
    }

export interface StoreGenreUpdateParameters {
    name?: UpdateField<string>,
    }

export interface UserUpdateParameters {
    username?: UpdateField<string>,
    is_admin?: UpdateField<boolean>,
    email?: UpdateField<string>,
    password?: UpdateField<string>,
    }

export interface WywwGenreUpdateParameters {
    name?: UpdateField<string>,
    }

export interface WywwMovieUpdateParameters {
    movie_title?: UpdateField<string>,
    original_title?: UpdateField<string>,
    content_rating?: UpdateField<WywwMovieContentRatingEnum>,
    description?: UpdateField<string>,
    released_year?: UpdateField<number>,
    runtime?: UpdateField<number>,
    rating?: UpdateField<number>,
    }


