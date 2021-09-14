[![codecov](https://codecov.io/gh/ph-piment/onion-scraper/branch/main/graph/badge.svg?token=A8HSHCRM5X)](https://codecov.io/gh/ph-piment/onion-scraper)

# onion-scraper
 Scraping with onion arch.

# Usage

```bash
git clone git@github.com:ph-piment/onion-scraper.git
cd onion-scraper
docker-compose up -d
# view http://localhost:16543
```

```bash
docker exec -i -t golang sh
migrate -database 'postgres://root:root@postgres:5432/os?sslmode=disable' -path ./migrations/ up
migrate -database 'postgres://root:root@postgres:5432/os?sslmode=disable' -path ./migrations/ down
```

```bash
xo schema "pgsql://root:root@localhost:5432/os?sslmode=disable" -o ./app/infrastructure/dao --src templates
```

```bash
cobra init --pkg-name github.com/ph-piment/onion-scraper --viper=false
cobra add import
```

```bash
wire ./cmd/di
```
