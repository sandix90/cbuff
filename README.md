create keyspace if not exists test with replication={'class':'SimpleStrategy', 'replication_factor': 1};

cd migrations && go-bindata -pkg migrations ./revisions