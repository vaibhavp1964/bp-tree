# bp-tree
B+ tree implementation in Golang

This will be a complete file based operation. Thus, while we may have some DS models which can work on data in-memory, all this data will be read from and persisted to the file system before and after each operation. We will try to eliminate in-memory persistence of data for the purposes of demonstration in this repository. Currently the save-to-file operation is not atomic. This will be done later.
