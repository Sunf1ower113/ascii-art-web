# Ascii-art-web-stylize
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.<br /> 
Ascii-art Web - is the version of the Ascii-art that runs it on the server.<br />
Ascii-art-web stylize adds styles to the Ascii-art-web frontend.
## Usage
Use `go run .` inside root directory or build docker image.<br /> 
For build docker image use `docker image build -f Dockerfile -t <name of image> .`<br /> and run: `docker run -p <target port>:3000 -d --name <name of container> <name of image>`<br /> or use `docker run -P -d --name <name of container> <name of image>` for run on random target port (use `docker ps -a` to find out target port)<br />
Than open web browser and go to `localhost:<target port>`

## Contributors
@sunf1ower113

## Notes
This is a project from Alem School GO Backend developer branch.
The project was added in the form and with the errors that were in the code review to show progress since graduation.

