### Database setup for development:

- Run:

```bash
docker run --detach --name pricetracker-mariadb --env MYSQL_ROOT_PASSWORD=<password> -p 3306:3306 mariadb:latest
```

- Attach:

```bash
docker exec -it pricetracker-mariadb mysql -u root -p
```
