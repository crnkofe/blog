# Sockets in Python

This code is runnable with Python 3

To run this mini demo first make a virtual environment:
```bash
mkvirtualenv --python=/usr/bin/python3 sockets
pip3 install -r requirements.txt
```

Run the server first:
```bash
python socket_server.py
```

Then run a socket client as follows:
```bash
./curl_socket.sh
```

Then submit an url with:
```
./submit_url.sh http://kofe.si
```
