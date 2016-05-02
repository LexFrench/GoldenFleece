# GoldenFleece
[![Build Status](https://travis-ci.org/LexFrench/GoldenFleece.svg?branch=master)](https://travis-ci.org/LexFrench/GoldenFleece)
[![Coverage](http://gocover.io/_badge/github.com/LexFrench/GoldenFleece)](http://gocover.io/github.com/LexFrench/GoldenFleece)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/LexFrench/GoldenFleece) [![License](https://img.shields.io/badge/license-GPL-lightgrey.svg)](https://github.com/LexFrench/GoldenFleece/blob/master/LICENSE)

GoLang json parser for unknown json data models

##Install
----------

    go get github.com/LexFrench/GoldenFleece

##To-Do List
----------
- Write benchmark tests

##Examples/How-To
----------
###Load JSON byte array
Load json byte array and parse it for use.

     data, err := GoldenFleece.Load([]byte(`{"foo": "bar"}`))

###Load JSON string
Load json string and parse it for use.

    data, err := GoldenFleece.Load("{\"foo\": 123})

### Dump byte array
Dump a byte array of the parsed json data

    result := data.Dump()

### Dump String
Dump an indented string of the parsed json data

    result := data.Dumps(4)

###Get Data
Use one of the following Get* methods to access data in json object.

    result, err := data.GetInt("foo")
    result, err := data.GetString("foo")
    result, err := data.GetFloat("foo")
    result, err := data.GetBool("foo")
    result, err := data.GetMap("foo")

###Get Array
Get array value from json object.
Note: This will return an array of interfaces so you will need to type assert the returned data in order to properly access the data.

    result, err := data.GetArray("foo")

###Get map/object from Array
Return a map/object from an array at the index specified.

    result, err := data.GetArrayMap("foo")

###Follow JSON paths
Using any of the above methods you can access data down nested JSON paths.

    result, err := data.GetInt("foo", "bar")

##FAQ:
----------
 **Q:** Why did you write JSON parser when there is already an amazing JSON encoding package built in?
 **A:** Hi Carl, you don't mind if I call you Carl, do you? So Carl, you're absolutely right; the built in encoding/JSON package is amazing -- so much so that I even use it in this library. To answer your question Carl, the main reason I created this package is because I was in the process of learning GoLang and realized that handling dynamic JSON was not as intuitive as some other scripting languages (\*cough* Python \*cough*), so I built something that would allow me to work with JSON in which I don't know what the data will look like coming from the server.

**Q:** But aren't there a bunch of other Go Libraries that will do just that?
**A:** I don't think you were listening, Carl! Like I said, I'm learning GoLang and this is a nice easy problem to solve and learn some of the basics of GO.

**Q:** Why do you keep calling me Carl?
**A:** Stop asking dumb questions Carl.
