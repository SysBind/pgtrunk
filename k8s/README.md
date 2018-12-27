
- bootstrap postgres cluster
``kubectl apply -f postgres.yaml``

- allow first pod to continue:
``kubectl exec postgres-0 -c pginit rm /stop``

- allow replication connections from subsequent pods:
``kubectl exec -ti postgres-0 -c pgsql bash``
``vi /var/lib/postgresql/data/pgdata/pg_hba.conf``
``su postgres -c "pg_ctl -D /var/lib/postgresql/data/pgdata reload"``

- for each subsequent pod, sync for replication from first pod, and allow continue:
``kubectl exec postgres-1 -c pginit -- pg_basebackup -h postgres-0.postgres -Upostgres -D /var/lib/postgresql/data/pgdata``
``kubectl exec postgres-1 -c pginit -- rm /stop``

- edit the statefulset definition to remove the init containers:
``kubectl edit sts postgres``

- create pgpool config
``kubectl create configmap pgpool-config --from-file=pool_passwd --from-file=pgpool.conf --from-file=pcp.conf``

- deploy pgpool cluster
``kubectl apply -f postgres.yaml``
