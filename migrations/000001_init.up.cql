create table if not exists frames.picts (
    camera_id        int,
    created_at       timestamp,
    pict_id          uuid,
    content          blob,
    previous_pict_id uuid,
    session          uuid,
    primary key (camera_id, created_at, pict_id)
) with clustering order by (created_at desc);

