module example

go 1.16

require "gmyst/gmyst" v0.0.0
replace "gmyst/gmyst" => ./gmyst
//require gmyst/gmyst => ./gmyst

require "gmyst/gmystcache" v0.0.0
replace "gmyst/gmystcache" => ./gmystcache
require github.com/gin-gonic/gin v1.7.1 // indirect
