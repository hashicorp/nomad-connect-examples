# This Consul service definition file can be used to run
# the two connect native service examples locally using
# only Consul (No Nomad or Docker required).

# Launch Consul
# $ consul agent -dev -config-file=services.hcl
#
# Launch uuid-api (in uuid-api/)
# $ PORT=2001 go run main.go
#
# Launch uuid-fe (in uuid-fe/)
# $ UPSTREAM=uuid-api PORT=2002 go run main.go

services {
  name = "uuid-api"

  port = 2001

  connect {
    native = true
  }
}

services {
  name = "uuid-fe"

  port = 2002

  connect {
    native = true
  }
}
