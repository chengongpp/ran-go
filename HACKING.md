# Hacking Guide

## Service Mesh

Data flows generally are seperated into Control Plane (CPlane) 
and Data Plane (DPlane). Artifact data streams go over the data plane and
routing rules go over the control plane.

## Reconnection

`pkg/ran/conn.go`

Nodes will try to probe aliveness with a short tcp connection at a heartbeat rate of 10 seconds.

Nodes that find out "dead" will adopt a reconnection policy as below:

1. one instant connection attempt.
2. two connections attempt with 10s cool down.
3. one connection attempt with 30s cool down.
4. one connection attempt with 60s cool down.

If the last attempt fails, the node will be marked as "dead" and will no longer try to connect to the mesh.