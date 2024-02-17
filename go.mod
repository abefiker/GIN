<<<<<<< HEAD
module example.com/greetings

go 1.21.6
=======
module example/hello

go 1.21.6

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
>>>>>>> 0c0c0bd496915921c2a06784513da0585ccb6184
