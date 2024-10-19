# Just runs the first recipe when called without one, defaulting printing out the other recipes
default:
	@# Printing out without echoing the command (@)
	@just --list

alias b := build

# Builds the golang binary and installs it to your ~/go/bin/ directory
build:
	go install .

# Runs the id server, this generates unqiue ids
id-server:
	go run ./id-server/main.go

# Runs the test that is defined by the challenge.
[confirm("Is the required id server running? y/n")]
basic-test: build
	just id-server
	maelstrom test -w unique-ids --bin ~/go/bin/unique-ids-gen --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

# Runs a more intense test that is likely to fail unless you have a good implementation.
[confirm("Is the required id server running? y/n")]
heavy-test: build id-server
	maelstrom test -w unique-ids --bin ~/go/bin/unique-ids-gen --time-limit 60 --rate 100000 --node-count 10 --availability total --nemesis partition

# Starts up the maelstrom server which shows detailed results.
results:
	maelstrom serve
