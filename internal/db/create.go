package db

import (
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/entities/image"
	"b3_ux_backend/internal/entities/storeapp"
	"b3_ux_backend/internal/entities/wywwmovie"
	"b3_ux_backend/internal/fsutils"
	"b3_ux_backend/internal/imageUtils"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/olekukonko/errors"

	_ "modernc.org/sqlite"
)

func CreateAppEntities() {
	ctx := context.Background()
	var err error

	appName := "b3_ux"
	appDir, err := fsutils.AppDir(appName)
	if err != nil {
		panic(err)
	}

	err = fsutils.EnsureDir(appDir)
	if err != nil {
		panic(err)
	}

	sqlPath := filepath.Join(appDir, "database.db")
	if fsutils.Exists(sqlPath) {
		_ = os.Remove(sqlPath)
	}

	// Client
	dsn := fmt.Sprintf("file:%s?_fk=1", filepath.ToSlash(sqlPath))
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB(dialect.SQLite, db)
	Client = entities.NewClient(entities.Driver(drv))

	if err := Client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	// Create app store entities
	err = createAppEntities(ctx, Client)
	if err != nil {
		panic(err)
	}

	// Create top movies entities
	err = createWywwEntities(ctx, Client)
	if err != nil {
		panic(err)
	}
}

func createAppEntities(ctx context.Context, client *entities.Client) error {
	//  Create a transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return errors.Wrapf(err, "createAppEntities transaction creation failed")
	}
	defer tx.Rollback()

	// Get genres from json file
	genresJsonPath := "assets/appStore/genres.json"

	genreData, err := readJsonFile[string](genresJsonPath)
	if err != nil {
		return errors.Wrapf(err, "CreateAppEntities genres.json read failed")
	}

	// Add genres to database
	var genreCreateBulk []*entities.StoreGenreCreate
	for _, v := range genreData {
		genreCreateBulk = append(genreCreateBulk, tx.StoreGenre.Create().SetName(v))
	}
	genres, err := tx.StoreGenre.CreateBulk(genreCreateBulk...).Save(ctx)
	if err != nil {
		return errors.Wrapf(err, "CreateAppEntities genre insertion to db failed")
	}

	//  Get apps from json file
	appsJsonPath := "assets/appStore/apps.json"
	appData, err := readJsonFile[appJsonData](appsJsonPath)
	if err != nil {
		return errors.Wrapf(err, "CreateAppEntities apps.json read failed")
	}

	// Add apps to database
	var appCreateBulk []*entities.StoreAppCreate
	for _, v := range appData {
		cr, err := v.TxCreate(tx, genres)
		if err != nil {
			return errors.Wrapf(err, "CreateAppEntities create genre failed")
		}

		appCreateBulk = append(appCreateBulk, cr)
	}

	err = bulkCreateBatched(
		appCreateBulk,
		100,
		func(batch []*entities.StoreAppCreate) error {
			_, err := tx.StoreApp.CreateBulk(batch...).Save(ctx)
			if err != nil {
				return errors.Wrapf(err, "CreateAppEntities insertion of apps to db failed")
			}

			return nil
		},
	)
	if err != nil {
		return errors.Wrapf(err, "CreateAppEntities batch insertion of apps to db failed")
	}

	//  Commit transaction
	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, "CreateAppEntities transaction commit failed")
	}

	return nil
}

func createWywwEntities(ctx context.Context, client *entities.Client) error {
	//  Create a transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities transaction creation failed")
	}
	defer tx.Rollback()

	// Get genres from json file
	genresJsonPath := "assets/top250/genre.json"

	genreData, err := readJsonFile[string](genresJsonPath)
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities genre.json read failed")
	}

	// Add genres to database
	var genreCreateBulk []*entities.WywwGenreCreate
	for _, v := range genreData {
		genreCreateBulk = append(genreCreateBulk, tx.WywwGenre.Create().SetName(v))
	}
	genres, err := tx.WywwGenre.CreateBulk(genreCreateBulk...).Save(ctx)
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities genre insertion to db failed")
	}

	//  Get apps from json file
	moviesJsonPath := "assets/top250/movie.json"
	appData, err := readJsonFile[movieJsonData](moviesJsonPath)
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities movie.json read failed")
	}

	// Add apps to database
	var movieCreateBulk []*entities.WywwMovieCreate
	var imageCreateBulk []*entities.ImageCreate
	var colorCreateBulk []*entities.ColorCreate
	for _, v := range appData {
		movCreate, imgCreates, colorCreates, err := v.TxCreate(tx, genres)
		if err != nil {
			return errors.Wrapf(err, "createWywwEntities parse movie data failed")
		}

		movieCreateBulk = append(movieCreateBulk, movCreate)
		imageCreateBulk = append(imageCreateBulk, imgCreates...)
		colorCreateBulk = append(colorCreateBulk, colorCreates...)
	}

	err = bulkCreateBatched(
		movieCreateBulk,
		100,
		func(batch []*entities.WywwMovieCreate) error {
			_, err := tx.WywwMovie.CreateBulk(batch...).Save(ctx)
			if err != nil {
				return errors.Wrapf(err, "createWywwEntities insertion of movies to db failed")
			}

			return nil
		},
	)
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities batch insertion of movies to db failed")
	}

	err = bulkCreateBatched(
		imageCreateBulk,
		100,
		func(batch []*entities.ImageCreate) error {
			_, err := tx.Image.CreateBulk(batch...).Save(ctx)
			if err != nil {
				return errors.Wrapf(err, "image insertion of movies to db failed")
			}

			return nil
		},
	)
	if err != nil {
		return errors.Wrapf(err, "image batch insertion of movies to db failed")
	}

	err = bulkCreateBatched(
		colorCreateBulk,
		100,
		func(batch []*entities.ColorCreate) error {
			_, err := tx.Color.CreateBulk(batch...).Save(ctx)
			if err != nil {
				return errors.Wrapf(err, "color insertion of movies to db failed")
			}

			return nil
		},
	)
	if err != nil {
		return errors.Wrapf(err, "color batch insertion of movies to db failed")
	}

	//  Commit transaction
	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, "createWywwEntities transaction commit failed")
	}

	return nil
}

func readJsonFile[T any](filePath string) ([]T, error) {
	err := fsutils.CheckFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "readJsonFile failed to check file")
	}

	data, err := fsutils.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "readJsonFile failed to read file")
	}

	var result []T

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Wrapf(err, "readJsonFile failed to unmarshal file")
	}

	return result, nil
}

type appJsonData struct {
	App            string   `json:"app"`
	Category       string   `json:"category"`
	Rating         float32  `json:"rating"`
	Reviews        int32    `json:"reviews"`
	Size           string   `json:"size"`
	Installs       string   `json:"installs"`
	Type           string   `json:"type"`
	Price          float32  `json:"price"`
	ContentRating  string   `json:"content_rating"`
	Genres         []string `json:"genres"`
	LastUpdated    string   `json:"last_updated"`
	CurrentVer     string   `json:"current_ver"`
	AndroidVer     string   `json:"android_ver"`
	InAppPurchases bool     `json:"in_app_purchases"`
	AdSupported    bool     `json:"ad_supported"`
}

func (a *appJsonData) TxCreate(tx *entities.Tx, genres []*entities.StoreGenre) (*entities.StoreAppCreate, error) {
	cat, err := appCategoryFromString(a.Category)
	if err != nil {
		return nil, errors.Wrapf(err, "TxCreate could not parse category")
	}

	typ, err := appTypeFromString(a.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "TxCreate could not parse type")
	}

	var genreIds []uuid.UUID
	for _, entity := range genres {
		if slices.Contains(a.Genres, entity.Name) {
			genreIds = append(genreIds, entity.ID)
		}
	}

	cr := tx.StoreApp.Create().
		SetName(a.App).
		SetCategory(cat).
		SetRating(a.Rating).
		SetReviews(a.Reviews).
		SetSize(a.Size).
		SetInstalls(a.Installs).
		SetType(typ).
		SetPrice(a.Price).
		SetContentRating(a.ContentRating).
		SetLastUpdated(a.LastUpdated).
		SetCurrentVer(a.CurrentVer).
		SetInAppPurchases(a.InAppPurchases).
		SetAdSupported(a.AdSupported)

	if len(genreIds) > 0 {
		cr.AddGenreIDs(genreIds...)
	}

	return cr, nil
}

func appCategoryFromString(value string) (storeapp.Category, error) {
	switch value {
	case "ARTIFICIAL_INTELLIGENCE":
		return storeapp.CategoryARTIFICIAL_INTELLIGENCE, nil

	case "ART_AND_DESIGN":
		return storeapp.CategoryART_AND_DESIGN, nil

	case "AUTO_AND_VEHICLES":
		return storeapp.CategoryAUTO_AND_VEHICLES, nil

	case "BEAUTY":
		return storeapp.CategoryBEAUTY, nil

	case "BOOKS_AND_REFERENCE":
		return storeapp.CategoryBOOKS_AND_REFERENCE, nil

	case "BUSINESS":
		return storeapp.CategoryBUSINESS, nil

	case "COMICS":
		return storeapp.CategoryCOMICS, nil

	case "COMMUNICATION":
		return storeapp.CategoryCOMMUNICATION, nil

	case "DATING":
		return storeapp.CategoryDATING, nil

	case "EDUCATION":
		return storeapp.CategoryEDUCATION, nil

	case "ENTERTAINMENT":
		return storeapp.CategoryENTERTAINMENT, nil

	case "EVENTS":
		return storeapp.CategoryEVENTS, nil

	case "FAMILY":
		return storeapp.CategoryFAMILY, nil

	case "FINANCE":
		return storeapp.CategoryFINANCE, nil

	case "FOOD_AND_DRINK":
		return storeapp.CategoryFOOD_AND_DRINK, nil

	case "GAME":
		return storeapp.CategoryGAME, nil

	case "HEALTH_AND_FITNESS":
		return storeapp.CategoryHEALTH_AND_FITNESS, nil

	case "HOUSE_AND_HOME":
		return storeapp.CategoryHOUSE_AND_HOME, nil

	case "LIBRARIES_AND_DEMO":
		return storeapp.CategoryLIBRARIES_AND_DEMO, nil

	case "LIFESTYLE":
		return storeapp.CategoryLIFESTYLE, nil

	case "MAPS_AND_NAVIGATION":
		return storeapp.CategoryMAPS_AND_NAVIGATION, nil

	case "MEDICAL":
		return storeapp.CategoryMEDICAL, nil

	case "NEWS_AND_MAGAZINES":
		return storeapp.CategoryNEWS_AND_MAGAZINES, nil

	case "PARENTING":
		return storeapp.CategoryPARENTING, nil

	case "PERSONALIZATION":
		return storeapp.CategoryPERSONALIZATION, nil

	case "PHOTOGRAPHY":
		return storeapp.CategoryPHOTOGRAPHY, nil

	case "PRODUCTIVITY":
		return storeapp.CategoryPRODUCTIVITY, nil

	case "SHOPPING":
		return storeapp.CategorySHOPPING, nil

	case "SOCIAL":
		return storeapp.CategorySOCIAL, nil

	case "SPORTS":
		return storeapp.CategorySPORTS, nil

	case "TOOLS":
		return storeapp.CategoryTOOLS, nil

	case "TRAVEL_AND_LOCAL":
		return storeapp.CategoryTRAVEL_AND_LOCAL, nil

	case "VIDEO_PLAYERS":
		return storeapp.CategoryVIDEO_PLAYERS, nil

	case "VIRTUAL_REALITY":
		return storeapp.CategoryVIRTUAL_REALITY, nil

	case "WEATHER":
		return storeapp.CategoryWEATHER, nil

	case "WEB3_AND_CRYPTO":
		return storeapp.CategoryWEB3_AND_CRYPTO, nil

	}

	return "", errors.New("unrecognised category")
}

func appTypeFromString(value string) (storeapp.Type, error) {
	switch value {
	case "Free":
		return storeapp.TypeFree, nil

	case "Paid":
		return storeapp.TypePaid, nil
	}

	return "", errors.New("unrecognised type")
}

func bulkCreateBatched[T any](items []T, batchSize int, save func([]T) error) error {
	for start := 0; start < len(items); start += batchSize {
		end := min(start+batchSize, len(items))

		if err := save(items[start:end]); err != nil {
			return err
		}
	}

	return nil
}

type movieJsonData struct {
	ContentRating string   `json:"content_rating"`
	Description   string   `json:"description"`
	MovieTitle    string   `json:"movie_title"`
	OriginalTitle string   `json:"original_title"`
	Type          string   `json:"type"`
	Genres        []string `json:"genres"`
	PosterPath    *string  `json:"poster_path"`
	BannerPath    *string  `json:"banner_path"`
	ReleasedYear  string   `json:"released_year"`
	Runtime       int32    `json:"runtime"`
	Rating        float32  `json:"rating"`
}

func (m *movieJsonData) TxCreate(tx *entities.Tx, genres []*entities.WywwGenre) (*entities.WywwMovieCreate, []*entities.ImageCreate, []*entities.ColorCreate, error) {
	rating, err := contentRatingFromString(m.ContentRating)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "invalid content rating")
	}

	i64, err := strconv.ParseInt(m.ReleasedYear, 10, 32)
	if err != nil {
		return nil, nil, nil, errors.New("invalid released_year")
	}
	year := int32(i64)

	movieId := uuid.New()
	bannerImagerId := uuid.New()
	posterImagerId := uuid.New()

	var imageResults []*entities.ImageCreate
	var colorResults []*entities.ColorCreate

	if m.PosterPath != nil && *m.PosterPath != "" {
		poster, err := imageUtils.EvalImageRepr(*m.PosterPath, 3)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "could not evaluate poster image representation: %s", *m.PosterPath)
		}

		imageResults = append(
			imageResults,
			tx.Image.Create().
				SetID(posterImagerId).
				SetMovieID(movieId).
				SetKind(image.KindPoster).
				SetFilename(filepath.Base(*m.PosterPath)).
				SetPath(fmt.Sprintf("./images/%s", filepath.Base(*m.PosterPath))).
				SetWidth(poster.Width).
				SetHeight(poster.Height),
		)

		for _, color := range poster.Colors {
			colorResults = append(
				colorResults,
				tx.Color.Create().
					SetImageID(posterImagerId).
					SetR(color.R).
					SetG(color.G).
					SetB(color.B).
					SetRatio(color.Ratio),
			)
		}
	}

	if m.BannerPath != nil && *m.BannerPath != "" {
		banner, err := imageUtils.EvalImageRepr(*m.BannerPath, 3)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "could not evaluate banner image representation: %s", *m.BannerPath)
		}

		imageResults = append(
			imageResults,
			tx.Image.Create().
				SetID(bannerImagerId).
				SetMovieID(movieId).
				SetKind(image.KindBanner).
				SetFilename(filepath.Base(*m.BannerPath)).
				SetPath(fmt.Sprintf("./images/%s", filepath.Base(*m.BannerPath))).
				SetWidth(banner.Width).
				SetHeight(banner.Height),
		)

		for _, color := range banner.Colors {
			colorResults = append(
				colorResults,
				tx.Color.Create().
					SetImageID(bannerImagerId).
					SetR(color.R).
					SetG(color.G).
					SetB(color.B).
					SetRatio(color.Ratio),
			)
		}
	}

	movieResult := tx.WywwMovie.Create().
		SetID(movieId).
		SetMovieTitle(m.MovieTitle).
		SetOriginalTitle(m.OriginalTitle).
		SetContentRating(rating).
		SetDescription(m.Description).
		SetReleasedYear(year).
		SetRuntime(m.Runtime).
		SetRating(m.Rating)

	var genreIds []uuid.UUID
	for _, entity := range genres {
		if slices.Contains(m.Genres, entity.Name) {
			genreIds = append(genreIds, entity.ID)
		}
	}

	if len(genreIds) > 0 {
		movieResult.AddGenreIDs(genreIds...)
	}

	return movieResult, imageResults, colorResults, nil
}

func contentRatingFromString(value string) (wywwmovie.ContentRating, error) {
	switch value {
	case "Approved":
		return wywwmovie.ContentRatingApproved, nil

	case "G":
		return wywwmovie.ContentRatingG, nil

	case "NC-17":
		return wywwmovie.ContentRatingNC17, nil

	case "Not Rated":
		return wywwmovie.ContentRatingNotRated, nil

	case "PG":
		return wywwmovie.ContentRatingPG, nil

	case "PG-13":
		return wywwmovie.ContentRatingPG13, nil

	case "Passed":
		return wywwmovie.ContentRatingPassed, nil

	case "R":
		return wywwmovie.ContentRatingR, nil

	}

	return "", errors.New("unrecognised content rating")
}
