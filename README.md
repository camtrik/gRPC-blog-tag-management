# Blog Tag Management by gRPC (Updating)

## Overview 
gRPC service used to manage the tag in [gin-blog](https://github.com/camtrik/gin-blog)

### Current Completed
list tags by: 
```bash
grpcurl -plaintext -d '{"name": "optionalTagName"}' localhost:8080 proto.TagService/GetTagList
```

### Documentation 
visit `http://127.0.0.1:8080/swagger-ui`, explore `http://127.0.0.1:8080/swagger/tag.swagger.json` to check the api documentation