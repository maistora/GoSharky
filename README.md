
![Bilby Stampede](http://maistora.badowl.net/blog-tech/wp-content/uploads/2014/02/sharky.gif)
Sharky
-------

Package Sharky will offer support for [GrooveShark](http://grooveshark.com/) implemented in [Go](http:/golang.org).
There are still many methods waiting to be implemented, but the most important once are good to go.

Installation
-------

    go get github.com/maistora/sharky

Current status
-------
The library is ready to use apart from the methods that I did not have access to. These methods panic with NO_ACCESS_ERR.


Example
-------  
  
```go
    func ExampleSharky_GetSongSearchResults() {
        shrky := ExampleSetUp()
        country := shrky.GetCountry("")
        song := shrky.GetSongSearchResults("counting stars", country, 10, 0)[0]
        fmt.Println(song.SongID)
        fmt.Println(song.SongName)
        fmt.Println(song.ArtistName)
        // Output:
        // 38377063
        // Counting Stars
        // OneRepublic
    }
```


License
-------

Copyright 2013 Nikolay Goergiev

Licensed under the BSD-2-Clause license (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://opensource.org/licenses/BSD-2-Clause

