from flask import Flask
from flask_cors import CORS

app = Flask(__name__)
cors = CORS(app)

from views import *

if __name__ == '__main__':
    app.run()