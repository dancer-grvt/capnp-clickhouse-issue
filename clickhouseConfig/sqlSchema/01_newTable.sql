CREATE OR REPLACE TABLE default.test_struct_new (
  titleField String,
  innerTuple Tuple(
    normalField String,
    normalField2 Int32,
    specialField UInt256,
    newSpecialField UInt256
  )
)
ENGINE = MergeTree()
ORDER BY titleField;
