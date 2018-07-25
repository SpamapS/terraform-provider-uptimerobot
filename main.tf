variable "api_key" {}

provider uptimerobot {
    api_key = "${var.api_key}"
}
