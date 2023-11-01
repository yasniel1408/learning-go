# correr los test
go test

# correr los test y ver mas datos
go test -v

# correr los test y ver coverage
go test -cover

ó

go test -coverprofile=coverage.out && go tool cover -html=coverage.out

# este ultimo es muy recomendado ya que
# me levanta una pagina y me da mucha mas info