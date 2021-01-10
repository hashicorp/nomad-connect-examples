job "cn-demo" {
  datacenters = ["dc1"]

  group "generator" {
    network {
      port "api" {
        // nomads choice
      }
    }

    service {
      name = "uuid-api"
      port = "api"

      connect {
        native = true
      }
    }

    task "generate" {
      driver = "docker"

      config {
        image        = "hashicorpnomad/uuid-api:v5"
        network_mode = "host"
      }

      env {
        BIND = "0.0.0.0"
        PORT = "${NOMAD_PORT_api}"
      }
    }
  }

  group "frontend" {
    network {
      port "http" {
        static = 9800
      }
    }

    service {
      name = "uuid-fe"
      port = "http"

      connect {
        native = true
      }
    }

    task "frontend" {
      driver = "docker"

      config {
        image        = "hashicorpnomad/uuid-fe:v5"
        network_mode = "host"
      }

      env {
        UPSTREAM = "uuid-api"
        BIND     = "0.0.0.0"
        PORT     = "${NOMAD_PORT_http}"
      }
    }
  }
}
