from flask import Flask, render_template, request, redirect, url_for
import redis
import os

app = Flask(__name__)

redis_server = os.environ["REDIS"]

# Initialize Redis
r = redis.Redis(host=redis_server, port=6379)

# Getting app version
if "APP_VERSION" in os.environ and os.environ["APP_VERSION"]:
    app_version = os.environ["APP_VERSION"]
else:
    app_version = "v1"

print("app_version is: " + app_version)

if "OPTION1" in os.environ and os.environ["OPTION1"]:
    option1 = os.environ["OPTION1"]
else:
    option1 = "Option 1"

if "OPTION2" in os.environ and os.environ["OPTION2"]:
    option2 = os.environ["OPTION2"]
else:
    option2 = "Option 2"

if "OPTION3" in os.environ and os.environ["OPTION3"] and app_version != "v1":
    option3 = os.environ["OPTION3"]
elif app_version != "v1":
    option3 = "Option 3"

if "TITLE" in os.environ and os.environ["TITLE"]:
    title = os.environ["TITLE"]
else:
    title = "Vote For Your Favorite Option"

# Set up initial vote counts
# TODO: implement this on redis proxy
if not r.exists("option1"):
    r.set("option1", 0)
if not r.exists("option2"):
    r.set("option2", 0)

if app_version == "v1":
   if not r.exists("option3"):
       r.set("option3", 0)


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
        vote = request.form["vote"]
        if vote == "option1":
            r.incr("option1")
        elif vote == "option2":
            r.incr("option2")
        elif vote == "option3" and app_version != "v1":
            r.incr("option3")
        return redirect(url_for("index"))

    # Get current vote counts
    option1_votes = int(r.get("option1") or 0)
    option2_votes = int(r.get("option2") or 0)
    if app_version != "v1":
        option3_votes = int(r.get("option3") or 0)

    if app_version != "v1":
        return render_template(
            "index.html",
            option1_votes=option1_votes,
            option2_votes=option2_votes,
            option3_votes=option3_votes,
            title=title,
            option1=option1,
            option2=option2,
            option3=option3,
        )
    else:
        return render_template(
            "index.html",
            option1_votes=option1_votes,
            option2_votes=option2_votes,
            title=title,
            option1=option1,
            option2=option2,
        )



if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=80)
