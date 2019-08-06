# migrate

DB migration tool.

Usage examples:

- run migrations for schema APC, with the directory `apc` found in the current path for local development: 
`migrate -db=apc -env=dev -dbuser=myuser -dbpass=Passw0rd -command=up`

- run migration for `test` schema in a specific directory
`migrate -db=apc -env=dev -dbuser=myuser -dbpass=Passw0rd -command=up -dir=/opt/dbscripts`