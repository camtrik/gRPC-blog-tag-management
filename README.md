# Blog Tag Management by gRPC (Updating)

## Overview 
gRPC service used to manage the tag in [gin-blog](https://github.com/camtrik/gin-blog)

### Current Completed
list tags by: 
```bash
grpcurl -plaintext -d '{"name": "optionalTagName"}' localhost:8080 proto.TagService/GetTagList
```