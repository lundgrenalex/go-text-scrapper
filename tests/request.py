#!/usr/bin/python3.4

import requests;
response = requests.post('http://127.0.0.1:10101/text', json={
    "text": "fdsjdhfs sgfdsf ii sdf Edasda asdad 😍😍😍😍😍😍 Хуй Пизда Макаронина #fsdfsdf #пизденышь"
})
print(response.text)