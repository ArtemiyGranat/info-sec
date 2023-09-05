# Steganography

## Usage
`put-message`:
```
put-message -m message.txt -s stegocontainer.txt -c container.txt
```
or
```
put-message -c container.txt < message.txt > stegocontainer.txt
```
`get-message`:
```
get-message -s stegocontainer.txt -m message.txt
```
or
```
get-message < stegocontainer.txt > message.txt 
```
To get available flags, use the `-h` flags when starting the program, e.g. `put-message -h`. 