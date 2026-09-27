package imageUtils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"math/rand"
	"os"
	"sort"

	_ "golang.org/x/image/webp"
)

type ImgRepr struct {
	Width  int
	Height int
	Path   string
	Colors []Color
}

type Color struct {
	R, G, B int
	Ratio   float32
}

type point struct {
	r float64
	g float64
	b float64
}

func EvalImageRepr(path string, count int) (*ImgRepr, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	size := img.Bounds().Size()

	// We do not need every pixel.
	pixels := sampleImage(img, 128)

	if len(pixels) == 0 {
		return nil, fmt.Errorf("image contains no pixels")
	}

	if count > len(pixels) {
		count = len(pixels)
	}

	clusters := kmeans(pixels, count, 20)

	resultColors := make([]Color, 0, len(clusters))

	for _, cluster := range clusters {
		resultColors = append(resultColors, Color{
			R:     int(uint8(math.Round(cluster.center.r))),
			G:     int(math.Round(cluster.center.g)),
			B:     int(math.Round(cluster.center.b)),
			Ratio: float32(cluster.count) / float32(len(pixels)),
		})
	}

	sort.Slice(resultColors, func(i, j int) bool {
		return resultColors[i].Ratio > resultColors[j].Ratio
	})

	return &ImgRepr{
		Width:  size.X,
		Height: size.Y,
		Path:   path,
		Colors: resultColors,
	}, nil
}

func sampleImage(img image.Image, targetSize int) []point {
	bounds := img.Bounds()

	width := bounds.Dx()
	height := bounds.Dy()

	stepX := max(1, width/targetSize)
	stepY := max(1, height/targetSize)

	result := make([]point, 0, targetSize*targetSize)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += stepY {
		for x := bounds.Min.X; x < bounds.Max.X; x += stepX {
			r, g, b, a := img.At(x, y).RGBA()

			// Ignore completely transparent pixels.
			if a == 0 {
				continue
			}

			result = append(result, point{
				r: float64(r >> 8),
				g: float64(g >> 8),
				b: float64(b >> 8),
			})
		}
	}

	return result
}

type cluster struct {
	center point
	count  int
}

func kmeans(points []point, count, iterations int) []cluster {
	clusters := make([]cluster, count)

	// Random initial centers.
	perm := rand.Perm(len(points))

	for i := range clusters {
		clusters[i].center = points[perm[i]]
	}

	assignments := make([]int, len(points))

	for range iterations {
		// Assignment step.
		for i, p := range points {
			bestIndex := 0
			bestDistance := math.MaxFloat64

			for j, c := range clusters {
				distance := colorDistance(p, c.center)

				if distance < bestDistance {
					bestDistance = distance
					bestIndex = j
				}
			}

			assignments[i] = bestIndex
		}

		// Update step.
		sums := make([]point, count)
		counts := make([]int, count)

		for i, p := range points {
			clusterIndex := assignments[i]

			sums[clusterIndex].r += p.r
			sums[clusterIndex].g += p.g
			sums[clusterIndex].b += p.b
			counts[clusterIndex]++
		}

		for i := range clusters {
			if counts[i] == 0 {
				continue
			}

			clusters[i].center = point{
				r: sums[i].r / float64(counts[i]),
				g: sums[i].g / float64(counts[i]),
				b: sums[i].b / float64(counts[i]),
			}

			clusters[i].count = counts[i]
		}
	}

	return clusters
}

func colorDistance(a, b point) float64 {
	r := a.r - b.r
	g := a.g - b.g
	bl := a.b - b.b

	// Squared Euclidean distance.
	// No need for sqrt() because we only compare distances.
	return r*r + g*g + bl*bl
}
