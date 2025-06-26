# HTTP routes segments in tree form

## Overview
###
This Go library provides an efficient HTTP router implementation that stores URL segments in a tree structure for optimal path matching. The tree structure allows for fast route lookups and supports RESTful API routing patterns.
###


## Features
###
* Tree-based routing: URL segments are stored in a tree structure for efficient matching

* REST API support: Designed with RESTful endpoints in mind

* High performance: Optimized path matching algorithm

* High customization: As it provides minimal functionality and allowed high customization, who want to write their http server from scratch with http package provided by golang.


## Installation
```
go get github.com/mddfaisal/route_tree
```


###

# Sample tree structure created by this library.
![tree](https://github.com/mddfaisal/route_tree/blob/develop/tree.png)


# Performance

## The tree-based approach provides:

* O(n) lookup time where n is the number of segments in the URL

* Minimal memory overhead

* No regular expression compilation for standard routes


Pull requests are welcome. For major changes, please open an issue.
License

[MIT](https://choosealicense.com/licenses/mit/)
