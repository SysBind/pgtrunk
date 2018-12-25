# Pgtrunk
![pgtrunk logo](https://raw.githubusercontent.com/SysBind/pgtrunk/master/logo.png)
Pgtrunk is A Postgres / Pgpool-II cluster wire-up.

Cloud friendly postgres cluster to acheive HA and Load Balancing (Read-Only)
Using streaming replication and Pgpool 2.

The pgtrunk's binary main goal is to prepare the ground before lanuching postgres,
after preparation (mainly syncing before replication on standby) it executes the postgres binary using the 'exec' function to completly replace itself.

Other goal is to provide an implementations for pgpool-failover, pgpool-recovery and others.
