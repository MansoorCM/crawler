# Web Crawler

## Description

This is a CLI app that accepts a URL, recursively crawls all linked URLs on each webpage, and generates a comprehensive report of all discovered URLs.

The app implements concurrency using goroutines, significantly enhancing performance. Users can also specify the maximum number of goroutines and the number of pages to crawl, allowing for customizable and efficient operations.

## Quick Start

Clone the repo and run the following command.

**go build -o out && ./out 'url' 'max_goroutines' 'max_pages'**

for eg. **go build -o out && ./out `https://www.chess.com/` 6 250**

the parameters 'max_goroutines' and 'max_pages' are optional and if not provided they will default to 5 and 100 respectively. 

![](https://github.com/MansoorCM/crawler/blob/main/crawler_demo.gif)