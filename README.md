CGO-FDK-AAC
===========

Golang CGO binding for the frauenhofer AAC decoder


fdk-aac lib
-----------

The fdk AAC lib is cloned from https://github.com/mstorsjo/fdk-aac.git.

No changes were made and the [fdk-aac LICENSE](fdk-aac/NOTICE) applies.

Some files likde documentation were deleted to keep the repo small.

The version used is v2.0.3.

The libraries will be precompiled so no additional dependencies are required for building.

The library is statically linked so no additional libraries need to be provided/preinstalled.

platforms
---------

Only linux on amd64 supported yet.