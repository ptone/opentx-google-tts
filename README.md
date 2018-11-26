# opentx-google-tts

This is a small golang program that takes a CSV file specifying text and output file, and uses [Google Cloud's text-to-speech (TTS) API](https://cloud.google.com/text-to-speech/) and [WaveNet](https://deepmind.com/blog/wavenet-generative-model-raw-audio/) voices to generate sounds for use in [OpenTX](https://www.open-tx.org/) transmitter platforms.

If you just want to download the already generated english files, download this repo as a [zip file](https://github.com/ptone/opentx-google-tts/archive/master.zip)
You need to install [golang](https://golang.org/)

Follow the steps for the [golang speech quickstart](https://cloud.google.com/text-to-speech/docs/quickstart-client-libraries)

You will need to set the GOOGLE_APPLICATION_CREDENTIALS environment variable to a credential per the quickstart. If you are not otherwise using the API - you should be able to generate many sound clips in the free tier of the service.

If you do not have a development environment set up, you should be able to perform these steps in [Cloud Shell](https://cloud.google.com/shell/docs/). You will need to zip the output files, and download via one of Cloud Shell's options, either download file from cloud-shell menu, or export from code editor.



