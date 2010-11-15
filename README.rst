go-json
=======

TODO


Overview
--------

TODO: blal adasdads

* Doing this
* Doing that


Installation
------------

1. Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). web.go targets the 'release' tag. To get the release tag, simply run `hg update -r release`. If you're running an outdated version of Go, or a version near the , it likely won't compile. 
2. Make sure that $GOROOT is set. web.go installs itself as a Go package, so it requires $GOROOT. See Go's install instructions for more information about $GOROOT. 
2. git clone git://github.com/hoisie/web.go.git
3. cd web.go && make install

You can also install using `goinstall github.com/hoisie/web.go`, but if you do this, the import statement in your go programs will be `import github.com/hoisie/web.go` instead of just `import web`.  

Example
-------
    
    package main
    
    import (
        "json"
    )
    
    func main() {
        
    }


To run the application, put the code in a file called hello.go and run:

    8g hello.go && 8l -o hello hello.8 && ./hello


Documentation
-------------

TODO


About
-----

go-json was written by Sascha Peilicke <saschpe@gmx.de>.
