module fail2go/fail2go

go 1.19

replace fail2go/winfw => ../winfw

require (
	fail2go/SSHLogParsing v0.0.0-00010101000000-000000000000
	fail2go/winfw v0.0.0-00010101000000-000000000000
)

replace fail2go/SSHLogParsing => ../SSHLogParsing
