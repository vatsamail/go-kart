module example.com/hellop

go 1.17

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	example.com/adv_greetings v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace example.com/adv_greetings => ../adv_greetings
