#! /bin/bash

echo -e "Start running the script..."
cd ../

echo -e "Start building the app..."
wails build --clean

echo -e "End running the script!"

# echo -e "Start running the script..."
# # cd ..

# echo -e "Start building the app..."

# sed -i "s/src\/assets\/js\/htmx.min.js/\/assets\/htmx.min.js/g" frontend/dist/index.html

# cp frontend/src/assets/js/htmx.min.js frontend/dist/assets

# npm run build

# wails build -clean -s -tags webkit2_40

# echo -e "End running the script!"

# https://geekland.eu/uso-del-comando-sed-en-linux-y-unix-con-ejemplos/
