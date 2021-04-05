create keyspace if not exists frames with replication={'class':'SimpleStrategy', 'replication_factor': 1};

cd migrations && go-bindata -pkg migrations ./revisions