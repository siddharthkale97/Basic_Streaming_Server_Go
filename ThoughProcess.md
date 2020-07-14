# Though Process behind the implementation

From the beginning, I was clear in mind what I had to do to achieve the goal of creating a simple yet efficient Basic Streaming Server.

First I gathered some info about file server in Go, transcoding files into m3u8 format, and basic API development(in Go). I had to go through all this reading material as this was my first time trying out a project in Go and video streaming.
After reading quite a bit and building a roadmap in my mind, I proceeded to follow Rohit Mundra’s Blog, link [here]( https://www.rohitmundra.com/video-streaming-server), to set up a server as soon as I could.

I encountered 2 major errors following his tutorial, the problem was the final URL was somehow not playable in VLC and there was no API endpoint for any kind of video transcoding and upload to the server.

So in my mind, I decided that I would need to create an API endpoint that can fetch files from users as well as transcode the files before uploading them to the server. Other than that I needed to fix the VLC problem.

So I discarded the code and started looking on the internet about Files servers in Go. From various sources, I was able to figure out how to serve the m3u8 files with the help of Http handlers in Go. The only difference from Rohit Mundra's code was that the streaming service didn’t need to have 2 URLs, one to serve the initial manifest file, and another to serve the .ts segments. 

I refined the new code a bit more by turning it into an API endpoint. Now I was able to view the raw m3u8 file in Postman, and any web browser. Moreover, the VLC problem was solved. I even tested it for different devices by setting up the server in a LAN environment. 

The next and most important step was to create the upload system. To figure out the upload system, I broke it down in the following steps

#### Creating a File Upload System

I went to the official documentation of Go to get an idea of quickly setting up a file upload system. As I felt I needed a little bit more idea of the system, I went to youtube and found out this awesome tutorial by TutorialEdge – link [here]( https://www.youtube.com/watch?time_continue=262&v=0sRjYzL_oYs&feature=emb_logo)

#### Manipulation of File Names

I had to do a certain manipulation of file names to rename and generate various strings for various functions. Again the official document provided me with enough information and an answer on Stackoverflow solved my problems.

#### Transcoding the Uploaded File

Rohit Mundra’s Blog gave an idea of using FFmpeg to transcode the file. The only problem was, it was a manual method and my method had to be automated to run on the server every time a file was uploaded. Reading about executing commands via Go, I stumbled upon the OS library of Go, which had the exact function of my need. Using the Os module, I was successful in transcoding my videos with FFmpeg.

#### Sending Back Json Response

This was the last and simplest task that I had to perform during the whole project. I quickly created a structure of the response I needed to send and used Json Encoder to convert the structure to a JSON response and send it back to the user.
At this point my basic API was ready and I quickly moved on to host the project and create the required documentation.

## Features/Improvement in Future

Due to the time constraint, I could not create some add ons to the server, but I will surely love to have them. In a few days, I will attempt to at least add one of the features.

All the features that came across my mind building this project –

* ABR Streaming Configuration
* Indexing all the uploaded files
* Better Error Handling
* In-Built HLV Player to view the stream on the server itself.

Thanks for reading through the whole document! I hope you like this project and will attempt to duplicate this project if you are new to creating a streaming server and Go just like me. Have a nice day!
