# go-text-scrapper
Useful service written on golang for parse and tokenize text to (common text, hashtags and emoji)

## Instruction for use

## Start of Text-Scrapper
* set $GOPATH & $GOROOT
* build app
```bash
go build main.go
```
* run it:
```bash
./start.sh
```
## How use (examples on Python)
```python
#!/usr/bin/python3

import requests

data = {
    "text": 'Lorem ipsum dolor sit. Рыба живет в пруду, нам поймать бы хоть одну. #hashtag #hashtag2 ⌛'
}

response = requests.post('http://127.0.0.1:10101/text', json=data)
data = response.json()
print(data)
```

**Response extracted data in json:**
```python
{
    'emoji': ['⌛'], 
    'text': ['lorem', 'ipsum', 'dolor'], 
    'tags': ['hashtag', 'hashtag2']
}
```

