# migrate

DB migration tool.

Usage examples:

- run migrations for schema APC, with the directory `apc` found in the current path for local development: 
`migrate -db=apc -env=dev -dbuser=myuser -dbpass=Passw0rd -command=up`

- run migration for `test` schema in a specific directory
`migrate -db=apc -env=dev -dbuser=myuser -dbpass=Passw0rd -command=up -dir=/opt/dbscripts`

- you can also indicate the db server directly using the parameter `dbhost`
`migrate -dbhost=127.0.0.1:3306 -db=apc -dbuser=myuser -dbpass=Passw0rd -command=up -dir=/opt/dbscripts`

To install the binary:

- You can get the latest relased binary from [Releases page](https://github.com/astropay/migrate/releases)

- or, you can clone the project locally and install it with `go build -i -o migrate github.com/astropay/migrate/cmd`

In both cases, you should have the `migrate` command in your path in order to be able to run ir from anywhere.