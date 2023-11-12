import os


class Config:
    DB_FILE = "testing.db"
    # SECRET_KEY = os.environ.get('SECRET_KEY')
    # SQLALCHEMY_DATABASE_URI = os.environ.get('SQLALCHEMY_DATABASE_URI')

    SECRET_KEY = '8faacc874b351d1d8c541b2c'
    SQLALCHEMY_DATABASE_URI = 'sqlite:///testing.db'

    # MAIL_SERVER = ''
    # MAIL_PORT = 587
    # MAIL_USE_TLS = True
    # MAIL_USERNAME = os.environ.get('EMAIL_USER')
    # MAIL_PASSWORD = os.environ.get('EMAIL_PASSWORD')


