job "cn-bridge-demo" {
  datacenters = ["dc1"]

  // Generator uses bridge network unlike cn-demo.nomad
  group "generator" {
    network {
      mode = "bridge"
      port "api" {
        // nomads choice
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
        image = "hashicorpnomad/uuid-api:v3"
      }

      env {
        BIND = "0.0.0.0"
        PORT = "${NOMAD_PORT_api}"
      }
    }
  }

  // Frontend continues to use host networking like cn-demo.nomad
  group "frontend" {
    network {
      port "http" {
        static = 9800
      }
    }

    service {
      name = "uuid-fe"
      port = "9800"

      connect {
        native = true
      }
    }

    task "frontend" {
      driver = "docker"

      config {
        image        = "hashicorpnomad/uuid-fe:v3"
        network_mode = "host"
      }

      env {
        UPSTREAM = "uuid-api"
        BIND     = "0.0.0.0"
        PORT     = "9800"
      }
    }
  }
}
