CREATE OR REPLACE TABLE default.test_struct_old (
  titleField String,
  innerTuple Tuple(
    normalField String,
    normalField2 Int32,
    specialField UInt256
  )
)
ENGINE = MergeTree()
ORDER BY titleField;
