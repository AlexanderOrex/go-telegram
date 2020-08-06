package telegram

import (
	"bytes"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"os"
	"strconv"
)

// Send PNG plot to Telegram chat using BotConfig
func SendPlot(config BotConfig, pointsX map[string][]float64, pointsY map[string][]float64, sizeX float64, sizeY float64, title string) string {
	tmpFile := createTmpFile()
	defer os.Remove(tmpFile.Name())

	generatePlot(tmpFile, pointsX, pointsY, sizeX, sizeY, title)
	sendPlotToTelegram(tmpFile, config)

	return "ok"
}

// Generate temporary plot file
func generatePlot(tmpFile *os.File, pointsX map[string][]float64, pointsY map[string][]float64, sizeX float64, sizeY float64, title string) {
	graph, err := plot.New()
	if err != nil {
		log.Fatal("Failed to create new plot", err)
	}

	graph.Title.Text = title
	graph.X.Label.Text = "X"
	graph.Y.Label.Text = "Y"
	graph.BackgroundColor = color.RGBA{A: 0}

	iteration := 1

	for plotKey := range pointsX {
		if len(pointsX[plotKey]) == len(pointsY[plotKey]) {
			line, err := plotter.NewLine(formatPoints(pointsX[plotKey], pointsY[plotKey]))
			if err != nil {
				log.Fatal("Failed to create new line", err)
			}

			line.LineStyle.Color = generateColors(iteration, len(pointsX))

			graph.Add(line)

			graph.Legend.Add(plotKey, line)
		} else {
			log.Fatal("Error: Len of pointsX: " + strconv.Itoa(len(pointsX)) + "; Len of pointsX: " + strconv.Itoa(len(pointsY)))
		}

		iteration++
	}

	if err := graph.Save(vg.Points(sizeX)*vg.Inch, vg.Points(sizeY)*vg.Inch, tmpFile.Name()); err != nil {
		log.Fatal("Failed to save graph to tmp file", err)
	}
}

// Create temporary file
func createTmpFile() *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "graph.*.png")

	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	return tmpFile
}

// Translate point from []float64 to plotter.XYs
func formatPoints(pointsX []float64, pointsY []float64) plotter.XYs {
	pts := make(plotter.XYs, len(pointsX))

	for key, _ := range pointsX {
		pts[key].X = pointsX[key]
		pts[key].Y = pointsY[key]
	}

	return pts
}

// Generate different colors for all graph lines
func generateColors(i int, numOfIterations int) color.RGBA {
	sum := 255 * 3
	r := math.Max(0, math.Min(255, float64(sum*i/(numOfIterations+1)-255*((i+0)%3))))
	g := math.Max(0, math.Min(255, float64(sum*i/(numOfIterations+1)-255*((i+1)%3))))
	b := math.Max(0, math.Min(255, float64(sum*i/(numOfIterations+1)-255*((i+2)%3))))

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

// Send plot file to telegram
func sendPlotToTelegram(tmpFile *os.File, config BotConfig) string {
	apiMethod := "sendPhoto"

	uri := getTelegramUri(config, apiMethod)

	fileContents, fileName := getPlotFileContent(tmpFile)

	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	part, err := writer.CreateFormFile("photo", fileName)
	if err != nil {
		log.Fatal("Cannot create form file", err)
	}

	part.Write(fileContents)

	_ = writer.WriteField("chat_id", strconv.Itoa(config.ChatId))

	err = writer.Close()
	if err != nil {
		log.Fatal("Cannot close writer", err)
	}

	resData := doRequest(uri, buffer, writer.FormDataContentType())

	return string(resData)
}

// Read content from file
func getPlotFileContent(tmpFile *os.File) ([]byte, string) {
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Fatal("Cannot open tmp file", err)
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Cannot read tmp file", err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal("Cannot get file stat", err)
	}
	file.Close()

	return fileContents, fi.Name()
}
