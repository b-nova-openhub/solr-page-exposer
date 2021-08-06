#!/bin/sh

echo "$(<LICENSE)"

# Configure your sopagex app below:
SOLR=
SOURCE="http://localhost:8002/pages, http://localhost:8002/pages"
PORT=8003

./bin/sopagex serve --solr=$SOLR --source=$SOURCE --port=$PORT