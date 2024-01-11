## Steps to Run the Tests

### Step 1: Run the Test
```go
go test
```

### Step 2: Run the Test with Detailed info
```go
go test -v
```

### Step 3: Run the Test and redirect the output to different file
```go
go test -coverprofile coverage.out
```

### Step 4: View the tests report in HTML
```go
go tool cover -html coverage.out -o coverage.html
```

### Step 5: Open the html file 
```bash
open coverage.html
```
**⚠️ If above command may not work in windows, then you can directly open the html file)**
