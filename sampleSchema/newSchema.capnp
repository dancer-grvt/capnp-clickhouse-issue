using Go = import "/go.capnp";
@0xb44039789e9fa5cd;
$Go.package("sampleSchema");
$Go.import("sampleSchema/newSchema");

struct TestStruct {
  titleField @0 :Text;
  innerTuple @1 :InnerTuple;
}

struct InnerTuple {
  normalField @0 :Text;
  normalField2 @1 :Int32;
  # UInt256 field
  specialField @2 :Data;
  # Newly added field
  newSpecialField @3 :Data;
}
