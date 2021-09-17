[![codecov](https://codecov.io/gh/ph-piment/onion-scraper/branch/main/graph/badge.svg?token=A8HSHCRM5X)](https://codecov.io/gh/ph-piment/onion-scraper)

# onion-scraper
 Scraping with onion arch.

# Usage

```bash
git clone git@github.com:ph-piment/onion-scraper.git
cd onion-scraper
make up-docker
# view http://localhost:16543

# migrate
docker exec -i -t golang sh
migrate -database 'postgres://root:root@postgres:5432/os?sslmode=disable' -path ./migrations/ up
migrate -database 'postgres://root:root@postgres:5432/os?sslmode=disable' -path ./migrations/ down

# generate xo
make gen-xo

# add command
make add-command hoge

# generate wire
make gen-wire

# run
make run
```
