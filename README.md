# Backend service

## Development

There is a VSCode devcontainer configuration that provides all dependencies for the backend including a PostgreSQL DB instance.

## Initial Startup

When you first start the devcontainer, you will need to first deploy the database schema to our local instance.

```bash
# This will deploy our DB schema and tables to our local instance
$ make db-deploy
```

Once that is complete, we can start the service:

```bash
# This will start our server running on port 8080
$ make start
```

### Postman

There is a Postman script included that can be used to test our endpoints [here](./integra-partners-backend.postman_collection.json)

### Swagger

When our server is up and running, we can also use the Swagger page that is located at http://localhost:8080/docs/index.html

## Database

Tech Stack:

- [PostgreSQL](https://www.postgresql.org/)
- [Sqitch](https://sqitch.org/) - Database change management framework for RDS deployment

## Making changes

>Note: all direct sqitch commands need to be made from the `./sqitch` directory

Sqitch creates plans for any changes made to the DB schema. In order to make any modifications you can follow below:

```bash
// Add a new plan
$ sqitch add newplan -n "Adding a new plan to adjust the DB"
```

This will create three SQL files for us, one under the `deploy`, `revert`, and `verify` directories nested in the `/sqitch` directory.

We will want to ensure any additions to the DB have an appropriate revert script (to revert only the changes made in your deploy script) and a verify script (to assert the changes made in the deploy script were fully functional).

Deploying the new changes can be done via:

```bash
# Deploys changes to the DB
$ sqitch deploy newplan

# Reverts the new changes from the DB
$ sqitch revert newplan

# Runs the verify script
$ sqitch verify newplan

# There are also Make commands that can be run from the root directory
# Deploys changes to the DB
$ make db-deploy

# Reverts the new changes from the DB
$ make db-revert

# Runs the verify script
$ make db-verify
```

