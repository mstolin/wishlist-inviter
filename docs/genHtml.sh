#! /bin/bash

search_dir=$(pwd)
for file in "$search_dir"/*
do
  extension="${file##*.}"
  filename="${file%.*}"
  if [ "$extension" = "apib" ]; then
    aglio --theme-variables flatly -i "$file" -o "$filename.html"
  fi;
done

