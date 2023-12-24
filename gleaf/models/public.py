from datetime import datetime
from flask import current_app
from gleaf.extensions import db

# https://www.youtube.com/watch?v=TwNp1UagE9U  // https://github.com/realpython/discover-flask/blob/master/db_create.py
# https://stackoverflow.com/questions/17652937/how-to-build-a-flask-application-around-an-already-existing-database
# https://www.reddit.com/r/flask/comments/2h2vb4/af_how_to_properly_work_with_blueprints/
# https://www.reddit.com/r/flask/comments/xwhvic/how_to_use_flask_with_existing_database/

class Status(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Integer, unique=True, nullable=False)
    description = db.Column(db.Text, nullable=False)


class Address(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    street = db.Column(db.String(120), nullable=False)
    street_2 = db.Column(db.String(120))
    city = db.Column(db.String(120), nullable=False)
    state = db.Column(db.String(25), nullable=False)
    zip = db.Column(db.String(5), nullable=False)


class Contact(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(120), nullable=False)
    phone_number = db.Column(db.String(10), nullable=False, unique=True)
    email = db.Column(db.String(120), nullable=False, unique=True)


class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    create_date = db.Column(db.DateTime, nullable=False, default=datetime.utcnow)
    email = db.Column(db.String(120), unique=True, nullable=False)
    username = db.Column(db.String(60), unique=True, nullable=False)
    hash = db.Column(db.String(60), nullable=False)
    address_id = db.Column(db.Integer, db.ForeignKey('address.id'), nullable=False)
    contact_id = db.Column(db.Integer, db.ForeignKey('contact.id'))
    group_id = db.Column(db.Integer, db.ForeignKey('clean_group.id'))
    bookings = db.relationship('Booking', backref='client', lazy=True)
    clean_group = db.relationship('CleanGroup', backref='member', lazy=True, foreign_keys=[group_id])

    

class ServiceCategory(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    category = db.Column(db.String(60), unique=True, nullable=False)
    description = db.Column(db.Text, nullable=False)


class Service(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(60), unique=True, nullable=False)
    status_id = db.Column(db.Integer, db.ForeignKey('status.id'), nullable=False)
    description = db.Column(db.String(120), nullable=False)
    category_id = db.Column(db.Integer, db.ForeignKey("service_category.id"), nullable=False)
    category = db.relationship("ServiceCategory", backref=db.backref("services", lazy=True))
    selling_price = db.Column(db.Float, nullable=False)
    image_url = db.Column(db.String(120), unique=True, nullable=False)


class CleanGroup(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    group_creator = db.Column(db.Integer, db.ForeignKey('user.id'), nullable=False)
    create_date = db.Column(db.DateTime, nullable=False, default=datetime.utcnow)


class RequestEstimate(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(120), nullable=False)
    phone = db.Column(db.String(10), nullable=False)
    email = db.Column(db.String(120), nullable=False)
    description = db.Column(db.String(1000), nullable=True)
    

class Booking(db.Model):
    id  = db.Column(db.Integer, primary_key=True)
    service_id = db.Column(db.Integer, db.ForeignKey('service.id'), nullable=False)
    selling_price = db.Column(db.Float, nullable=False)
    address_id = db.Column(db.Integer, db.ForeignKey('address.id'), nullable=False)
    contact_id = db.Column(db.Integer, db.ForeignKey('contact.id'), nullable=False)
    user_id = db.Column(db.Integer, db.ForeignKey('user.id'), )
    created_at = db.Column(db.DateTime, nullable=False, default=datetime.utcnow)
    requested_for = db.Column(db.DateTime, nullable=False)
    completed_at = db.Column(db.DateTime)
    status_id = db.Column(db.Integer, db.ForeignKey('status.id'), nullable=False)
    payment_status = db.Column(db.Boolean, nullable=False)
    group_booking = db.Column(db.Boolean, nullable=False)
    group_id = db.Column(db.Integer, db.ForeignKey('clean_group.id'))
    

