# The Ran Transport Protocol

RanTP is a transport protocol modified from HTTP/2 protocol.

1. Length Segment (3 bytes)
2. Statement Segment (2 byte)
4. Reserved Segment (1 byte)
5. Stream ID Segment (3 bytes  )
6. Data Segment ($length bytes)