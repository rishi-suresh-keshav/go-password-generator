# go-password-generator
A simple app to generate password containing random characters. It contains a mixture of uppercase, lowercase, digits and special characters.
1. First character is always an uppercase letter
2. 3 random digits included
3. 3 special characters included


## App can be run in 2 ways
1. Using cli app
2. Using rest api call

## Using cli app to generate password
Build the app using the following command first
    
    make build-gpg-cli

Run the cli app to generate the password 

    ./gpg generate
    ./gpg generate --length 20

Note: Run `./gpg --help` for help in using command


## Using api to generate password
call `http://localhost:8000/go-password-generator/generate` endpoint to generate

## Building and running docker image
    docker build -t go-password-generator:latest .
    docker run -it -p 8000:8000 "go-password-generator:latest"