
import json

from flask import request, render_template

from app import app

@app.route('/data/')
def