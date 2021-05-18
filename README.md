# Gomedo

Gomedo allows for easy monitoring of a [tomedo](https://tomedo.de) medical appointment calendar.

*This project is not affiliated, associated, authorized, endorsed by, or in any way officially connected with
zollsoft GmbH or any of its subsidiaries or its affiliates.*

## Usage

Using the provided [docker-compose configuration](deployments/docker-compose.yaml) is probably the fastest way to get started.
Running the binary directly works just as well, provided all required environment variables are set correctly.
See the [Dockerfile](build/docker/go/Dockerfile) for detailed build instructions.

### docker-compose

```shell
docker-compose -f ./deployments/docker-compose.yaml up -d
```

## Configuration

The following environment variables should be used to monitor a specific calendar.

| Environment          | Description                         | Required |
| -------------------- | ----------------------------------- | -------- |
| UNIQUE_IDENTIFIER    | Location specific unique identifier | Yes      |
| SCRAPE_ENDPOINT      | HTTP endpoint used for scraping     | Yes      |
| SCRAPE_INTERVAL      | Interval used for scraping          | Yes      |
| APPOINTMENT_KEYWORDS | Comma separated list of keywords    | No       |
| NOTIFICATION_HOOKS   | Comma separated list of webhooks    | No       |

### Example

```dotenv
UNIQUE_IDENTIFIER=610befd11b2f8
SCRAPE_ENDPOINT=https://onlinetermine.zollsoft.de/includes/searchTermine_app_feature.php
SCRAPE_INTERVAL=30s
APPOINTMENT_KEYWORDS=impftermin,covid-19,biontech,astrazeneca
NOTIFICATION_HOOKS=https://example.org/webhook/68e11060-89dc-4031-9a63-a7a2d7e29927
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
