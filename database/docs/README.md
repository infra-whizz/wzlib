# Cockroach DB

Whizz is using SQL database behind the scenes for storing all the data
about cluster, clients etc. The database used is CockroachDB.

## TL;DR Setup

Quick answer:

Cockroach DB binary is expected in the same directory as current
scripts, but `etc` and `node[1-3]` directories one level up:

```
bin/cockroach
etc/
node1/
node2/
node3/
```

Once you've got that somewhere, do the following:

1. Run `roach_setup.sh` (needed only once)

2. Run all three `start_node_*.sh` scripts

3. Run `roach_init.sh` (needed only once)

4. Rock-n-roll

5. Run `stop_cluster.sh` if you've finished.


Long answer (the same as above):
https://www.cockroachlabs.com/docs/v19.2/secure-a-cluster.html
