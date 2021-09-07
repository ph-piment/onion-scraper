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