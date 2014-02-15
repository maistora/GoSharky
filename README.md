
![Bilby Stampede](https://raw2.github.com/maistora/sharky/master/sharky.png)

Package Sharky will offer support for [GrooveShark](http://grooveshark.com/) implemented in [Go](http:/golang.org).
There are still many methods waiting to be implemented, but the most important once are good to go.

Installation
-------

    go get github.com/maistora/sharky

TODO
-------

* Implement the rest of the methods
* Refactoring

Example
-------  
  
```go
    func findSongAndGetStream() string {  
        sharky := setUp()  
        country := sharky.GetCountry("") // returns country data for the requestor's IP  
        song := sharky.GetSongSearchResults("counting stars", country, 10, 0)[0]  
        fmt.Println(song)  
        // Output:  
        // &{38377063 Counting Stars 401901 OneRepublic 8545065 Native 8545065-20140206135006.jpg  true false 0 }  
            
        streamDetails := sharky.GetStreamKeyStreamServer(song.SongID, country, false)  
        return streamDetails.Url  
    }  
```


License
-------

Copyright 2013 Nikolay Goergiev

Licensed under the BSD-2-Clause license (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://opensource.org/licenses/BSD-2-Clause

