#!/bin/bash

for i in $(find . -regex '.*.md' -print); do
  grep -q "^---" $i
  if [ $? -gt 0 ]; then continue; fi
  cat $i | tr -d '\n' | grep -lv "^---.*permalink: .*---" &> /dev/null
  if [ $? -eq 0 ]; then
    # permalink is absent, add permalink
    PERMALINK="/$(echo $i | sed -E 's#(modules_)(en|ru)/#\2/modules/#' | sed 's#docs/##g'| tr '[:upper:]' '[:lower:]' | sed 's#\.md$#.html#' | sed 's#^\.\/##' | sed 's#readme\.html$##' )"
    sed -i "1apermalink: $PERMALINK" $i
  fi
done
