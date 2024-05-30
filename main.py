from flask import Flask, request
import ctypes, os, logging

def main():
    log = logging.getLogger('werkzeug')
    app = Flask(__name__)
    kernel32 = ctypes.WinDLL('kernel32')
    
    log.setLevel(logging.ERROR)
    kernel32.SetConsoleTitleW('https://github.com/prrefer/console')
    
    @app.route('/print', methods=['POST'])
    def _print():
        print(request.data.decode())
        return '', 200
    
    @app.route('/warn', methods=['POST'])
    def warn():
        print(f'[[93m*[0m]: {request.data.decode()}')
        return '', 200
    
    @app.route('/error', methods=['POST'])
    def error():
        print(f'[[31m*[0m]: {request.data.decode()}')
        return '', 200
    
    @app.route('/clear', methods=['POST'])
    def clear():
        os.system('cls')
        return '', 200
    
    @app.route('/input', methods=['POST'])
    def _input():
        response = input(request.data.decode())
        return response, 200
    
    @app.route('/title', methods=['POST'])
    def set_title():
        kernel32.SetConsoleTitleW(request.data.decode())
        return '', 200

    app.run(host='localhost', port=8080)

if __name__ == '__main__':
    main() 
