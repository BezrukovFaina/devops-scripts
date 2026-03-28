# utils.py
import re
import json
import os
import logging
import subprocess

class EnvironmentVariables:
    def __init__(self, config_file_path):
        self.config_file_path = config_file_path

    def load(self):
        try:
            with open(self.config_file_path, 'r') as config_file:
                return json.load(config_file)
        except FileNotFoundError:
            logging.error(f'Config file {self.config_file_path} not found')
            raise

    def get(self, key):
        return os.environ.get(key)

def replace_in_file(file_path, pattern, replacement):
    try:
        with open(file_path, 'r') as file:
            content = file.read()
        new_content = re.sub(pattern, replacement, content)
        with open(file_path, 'w') as file:
            file.write(new_content)
    except Exception as e:
        logging.error(f'Failed to replace in file {file_path}: {str(e)}')
        raise

def run_command(command):
    logging.info(f'Running command: {command}')
    try:
        output = subprocess.check_output(command, shell=True)
        return output.decode('utf-8').strip()
    except Exception as e:
        logging.error(f'Failed to run command: {str(e)}')
        raise