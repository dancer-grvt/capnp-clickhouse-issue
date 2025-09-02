CREATE OR REPLACE TABLE default.test_dest_table(
  titleField String,
  innterTuple Tuple(
    normalField String,
    normalField2 Int32,
    specialField UInt256
  )
)
ENGINE = ReplicatedReplacingMergeTree(
  '/clickhouse/tables/{shard}/test_dest_table'
  '{replica}'
)

