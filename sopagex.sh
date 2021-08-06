#!/bin/sh

echo "$(<LICENSE)"

# Configure your sopagex app below:
PORT=8080

./bin/sopagex serve --port=$PORT