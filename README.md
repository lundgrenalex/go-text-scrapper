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
import requests

response = requests.post('http://127.0.0.1:10101/text', json={
    "text": "fdsjdhfs sgfdsf ii sdf Edasda asdad 😍😍😍😍😍😍 Жил был Дед-медвед #fsdfsdf #которыйговорилпревед"
})

print(response.text)
```

**Response extracted data in json:**
```json
{
    'text': [],
    'tags': [],
    'emoji': []
}
```

