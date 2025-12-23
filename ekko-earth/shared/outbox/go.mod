module github.com/ekko-earth/shared/outbox

go 1.25.4

require github.com/google/uuid v1.6.0

require (
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/ekko-earth/shared/adapters v0.0.0
	github.com/ekko-earth/shared/messaging v0.0.0
)

require golang.org/x/sys v0.39.0 // indirect

replace github.com/ekko-earth/shared/adapters => ../adapters

replace github.com/ekko-earth/shared/messaging => ../messaging
