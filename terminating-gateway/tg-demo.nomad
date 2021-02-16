job "term-cd" {
  datacenters = ["dc1"]

  group "api" {
    network {
      mode = "host"
      port "port" {
        static = "9001"
      }
    }

    task "web" {
      driver = "docker"

      config {
        image        = "hashicorpnomad/counter-api:v3"
        network_mode = "host"
      }
    }

    service {
      name = "count-api"
      port = "port"
    }
  }

  group "gateway" {
    network {
      mode = "bridge"
    }

    service {
      name = "t1000"

      connect {
        gateway {
          proxy {}
          terminating {
            service {
              name = "count-api"
            }
          }
        }
      }
    }
  }

  group "dashboard" {
    network {
      mode = "bridge"

      port "http" {
        static = 9002
        to     = 9002
      }
    }

    service {
      name = "count-dashboard"
      port = "9002"

      connect {
        sidecar_service {
          proxy {
            upstreams {
              destination_name = "count-api"
              local_bind_port  = 8080
            }
          }
        }
      }
    }

    task "dashboard" {
      driver = "docker"

      env {
        COUNTING_SERVICE_URL = "http://${NOMAD_UPSTREAM_ADDR_count_api}"
      }

      config {
        image = "hashicorpnomad/counter-dashboard:v3"
      }
    }
  }
}
