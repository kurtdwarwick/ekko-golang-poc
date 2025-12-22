module github.com/ekko-earth/shared/http

go 1.25.4

require github.com/gorilla/mux v1.8.1

require github.com/ekko-earth/shared/messaging v0.0.0

require github.com/google/uuid v1.6.0 // indirect

replace github.com/ekko-earth/shared/messaging => ../messaging
