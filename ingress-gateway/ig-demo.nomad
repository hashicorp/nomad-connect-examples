job "ig-demo" {

  datacenters = ["dc1"]

  group "ingress-group" {

    network {
      mode = "host"

      # This example will enable plain HTTP traffic to access the uuid-api connect
      # native example service on port 8080.
      port "inbound" {
        static = 8080
      }

      # When running an ingress gateway in host networking mode, the underlying
      # Envoy proxy creates an admin interface listener bound to localhost that
      # requires the allocation of a port.
      port "envoy" {
        static = 19001
      }
    }

    service {
      name = "my-ingress-service"

      # The Envoy proxy admin interface listener will use the service port to
      # determine its localhost bind address.
      port = "envoy"

      connect {
        gateway {

          // Consul gateway [envoy] proxy options.
          proxy {
            # Envoy proxy options are documented at
            # https://www.nomadproject.io/docs/job-specification/gateway#proxy-parameters
            connect_timeout = "500ms"
          }

          // Consul Ingress Gateway Configuration Entry.
          ingress {
            # Nomad will automatically manage the Configuration Entry in Consul
            # given the parameters in the ingress block.
            #
            # Additional options are documented at
            # https://www.nomadproject.io/docs/job-specification/gateway#ingress-parameters
            listener {
              port     = 8080
              protocol = "tcp"
              service {
                name = "uuid-api"
              }
            }
          }
        }
      }
    }
  }

  # The UUID generator from the connect-native demo is used as an example service.
  # The ingress gateway above makes access to the service possible over normal HTTP.
  # For example,
  #
  # $ curl $(dig +short @127.0.0.1 -p 8600 uuid-api.ingress.dc1.consul. ANY):8080
  group "generator" {
    network {
      mode = "host"
      port "api" {
        to = -1
      }
    }

    service {
      name = "uuid-api"
      port = "${NOMAD_PORT_api}"

      connect {
        native = true
      }
    }

    task "generate" {
      driver = "docker"

      config {
        image        = "hashicorpnomad/uuid-api:v3"
        network_mode = "host"
      }

      env {
        BIND = "0.0.0.0"
        PORT = "${NOMAD_PORT_api}"
      }
    }
  }
}

