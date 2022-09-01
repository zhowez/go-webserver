# GO-Webserver

## Description
This is a small webserver written in go. 
It will serve up an index.* file along with and other file needed for the webpage.
The webserver will work with client side routing and allows users to access any page from its url.
It will forward any /api/ request through the reverse proxy.


## Flags
- p : specifies the port for the webserver to listen on. Default 8000
- pp : specifies the port to forward /api/ request. Default 8001
- d : specifies the sub-directory location of the index.* in /html

## Files 
webserver.go - the webserver
/html - the folder that holds all static file
    - A react app with client side routeing for testing
    -  /test - a sub-directory to test the "-d" flag
      - index.html a test html page

