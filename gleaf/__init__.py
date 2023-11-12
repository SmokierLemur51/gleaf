from flask import Flask

from .config import Config
from .extensions import db, bcrypt

def create_app(config_class=Config):
    app = Flask(__name__, static_url_path="/static")
    app.config.from_object(Config)
    app.config["DATABASE"] = "testing.db"
    
    db.init_app(app)
    bcrypt.init_app(app)

    from gleaf.public.routes import public
    app.register_blueprint(public, url_prefix="/")

    from gleaf.portal.routes import portal
    app.register_blueprint(portal, url_prefix="/portal")

    return app
