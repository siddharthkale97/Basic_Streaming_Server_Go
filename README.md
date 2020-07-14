# Basic_Streaming_Server_Go

This is a Basic Video Streaming Server created in GO

# Brief Description

I have created a basic streaming server using GO and ffmpeg. This server host m3u8 files which are streamed on demand by using the API. Users also have the option to upload and stream their video files on demand as well. Currently the server is only supporting m4 files for upload. 

*Note-* To read the though process behind this project, open ThoughtProcess.md file

# Installation and Setup --

### Installation

1. Install Go on your local machine, link to installation guide [here](https://golang.org/doc/install)
2. Get the gorilla/mux dependency by running the command 
```sh
go get -u github.com/gorilla/mux
```
3. Install ffmpeg tool, link to installation guide [here](https://ffmpeg.org/download.html)

### Setup

1. Move into the root folder of the repository by `cd Basic_Streaming_Server_Go-master`.
2. Run the command `go run test.go`. If prompted for any Firewall access, please click on allow.
3. Now the server should be up and running and you should see a output `Starting server on 8080`, 8080 is the default port here.

# Using the API

This API has 2 endpoints, one to view the transcoded m3u8 file, either in raw format or in a HLS player. The other endpoint is to upload a file (max size 10 MB) and get it streamed on demand.

### API endpoint to view

Users have the ability to steam the uploaded files which are present in the server.
Currently, a person has to know the exact filename to access it.
One could goto a web browser and paste the following URL to access a m3u8 file in raw format.
`http://localhost:8080/media/upload-009132623.m3u8`
This is a test file that I had uploaded for testing purpose. As you can see after /media , one has to put up a file name. This whole URL is provided to you when you are uploading your file. As of now there is no method to index the files. :(

You can also paste the same URL in Postman and execute a Get request, you will get the same output as the browser.

The File can be streamed on demand, that's why we have created this API, if you have HLS player extension enabled in Google Chrome or follow the steps below for VLC player-

* Open VLC player and click on Media.
* In the Dropdown menu click on Open Network Stream.
* A dialouge box will appear where you will find an input field to enter a URL.
* Paste the URL of the file you want to stream or for testing paste the below URL 
    `localhost:8080/media/upload-009132623.m3u8`
* Click on Play, the stream should start in few seconds...
* Please Note that the server should be running during the above procedure, else streaming wont be possible.

### API endpoint to upload

To upload your own videos to the server, you will need a mp4 file with size of less than 10 MB, and to follow the procedures explained below.

* The File you upload will be downgraded to *640x360* resolution.
* Use curl or Postman to upload the file to the server, I will demonstrate how to do it with Postman
* Create a New Request in Postman app.
* Set the Method to POST
* Enter the following URL in the Input Field titled *Enter request url* 
    `http://127.0.0.1:8080/upload`
* Below this input field you will find a button named body.
* Click on that button, and check form data
* In the key add myFile, hovering over the key myFile will present you with a dropdown expander with a sort of downward facing arrow
* Expand the dropdown menu and select file
* In the field Value, directly opposite to the key myFile field, there will be an option to upload a file. Please upload a mp4 file with size less than 10 MB.
* Click on Send, and wait for few seconds.
* If successful you will get a Json response of the following type
```sh
{
    "Staus": "Success",
    "UploadName": "your-upload-file-name",
    "UploadSize": "your-upload-file-size-in-KB(approx)",
    "URL": "url-from-which-you-can-stream-the-file-on-demand"
}
```

Example -- 

```sh
{
    "Staus": "Success",
    "UploadName": "cat.mp4",
    "UploadSize": 7501,
    "URL": "localhost:8080/media/upload-009132623.m3u8"
}
```
* Now you can follow the steps in API end point to view, to literally view/stream your file

# Demo

Demo of this project is uploaded [here](www.link.here), Ops its currently unavailable!!

# Demo Video Link

A video link giving demo will be uploaded shortly!!

## To read the though process behind this project, open ThoughtProcess.md file