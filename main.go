package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"encoding/csv"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func main() {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// csvFile, _ := os.Open("en-US-taranis.csv")
	// note that the default taranis file uses ; as delimiter
	// reader.Comma = ';'

	csvFile, _ := os.Open("custom.csv")

	defer csvFile.Close()
	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range csvData {
		// voice parameters and audio file type.
		req := texttospeechpb.SynthesizeSpeechRequest{
			// Set the text input to be synthesized.
			Input: &texttospeechpb.SynthesisInput{
				InputSource: &texttospeechpb.SynthesisInput_Text{Text: each[2]},
			},
			// Build the voice request, select the language code ("en-US") and the SSML
			// voice gender ("neutral").
			Voice: &texttospeechpb.VoiceSelectionParams{
				LanguageCode: "en-US",
				Name:         "en-US-Wavenet-C",
			},
			AudioConfig: &texttospeechpb.AudioConfig{
				AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
				// Increase the pitch and speed slightly
				Pitch:        2.8,
				SpeakingRate: 1.3,
				// Taranis requires 32000hz
				SampleRateHertz: 32000,
			},
		}

		resp, err := client.SynthesizeSpeech(ctx, &req)
		if err != nil {
			log.Fatal(err)
		}

		filename := path.Join(each[0], each[1])
		os.MkdirAll(each[0], os.ModePerm)
		err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Audio content written to file: %v\n", filename)
	}

}
