job "cn-demo" {
  datacenters = ["dc1"]

  group "generator" {
    network {
      mode = "host"

      port "api" {
        // nomads choice
      }
    }

    service {
      name = "uuid-api"
      port = "${NOMAD_PORT_api}"

      connect {
        native = "generate"
      }
    }

    task "generate" {
      driver = "docker"

      config {
        image = "hashicorpnomad/uuid-api:v0.0.1"
      }

      env {
        BIND = "0.0.0.0"
        PORT = "${NOMAD_PORT_api}"
      }
    }
  }

  group "frontend" {
    network {
      mode = "host"

      port "http" {
        static = 9800
      }
    }

    service {
      name = "uuid-fe"
      port = "9800"

      connect {
        native = "frontend"
      }
    }

    task "frontend" {
      driver = "docker"

      config {
        image = "hashicorpnomad/uuid-fe:v0.0.1"
      }

      env {
        UPSTREAM = "uuid-api"
        BIND     = "0.0.0.0"
        PORT     = "9800"
      }
    }
  }
}
