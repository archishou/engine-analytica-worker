from flask import Flask, request

# create the Flask app
app = Flask(__name__)


@app.route('/worker-ready')
def default_get():
    workload = {
        'repoUrl': 'https://github.com/archishou/MidnightChessEngine',
        'baseBranch': 'master',
        'devBranch':  'dev-sample',
        'batchSize': '10',
    }
    return workload


if __name__ == '__main__':
    # run app in debug mode on port 5000
    app.run(debug=True, port=65123)