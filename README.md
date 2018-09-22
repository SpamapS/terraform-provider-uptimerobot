# terraform-provider-uptimerobot

Apache 2 licensed Uptimerobot Terraform Provider

## Installation

```bash
git clone git@github.com:SpamapS/terraform-provider-uptimerobot.git
cd terraform-provider-uptimerobot
go build
mv terraform-provider-uptimerobot $HOME/.terraform.d/plugins
```

## Usage

Add the uptimerobot provider:

```hcl
provider "uptimerobot" {}
```

Set the `UPTIMEROBOT_API_KEY` environment variable, or add it as a provider attribute:

```hcl
provider "uptimerobot" {
  api_key = "xxx"
}
```

Add a monitor resource:

```hcl
resource "uptimerobot_monitor" "my_monitor" {
  url           = "http://example.com"
  type          = "http"
  friendly_name = "Example"
}
```

... where `type` is one of [MonitorTypeNames][].

[MonitorTypeNames]: https://github.com/SpamapS/uptimerobot/blob/b95e7aed2bfe79e8eb6ea231a491a1303212bbc7/uptimerobot.go#L55-L64
