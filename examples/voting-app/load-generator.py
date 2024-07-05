import requests
import time

# The URL to send the POST requests to
host = "prod.app.localhost"
url = "http://127.0.0.1/"

# Headers to be included in the POST requests
headers = {
    "Origin": f"http://{host}",
    "Host": host,
}

# Data to be sent in the POST requests
data_options = ["option1", "option2"]
data_index = 0


# Function to send a burst of 5 POST requests
def send_burst(data):
    print(f"New burst of {data}")
    for _ in range(5):
        response = None
        try:
            response = requests.post(url, headers=headers, data={"vote": data})
            print(f"Sent '{data}' - Response status code: {response.status_code}")
        except requests.exceptions.RequestException as e:
            print(f"Error sending '{data}' - {e}")


# Send bursts of 5 POST requests every 5 seconds, alternating between 'Cats' and 'Dogs'
while True:
    send_burst(data_options[data_index])
    data_index = (data_index + 1) % 2
    time.sleep(5)
