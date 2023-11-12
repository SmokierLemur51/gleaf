from flask import Blueprint, render_template, url_for, current_app, g, request
from flask.helpers import redirect

from greenleaf.extensions import db
from greenleaf.models import ServiceCategory

public = Blueprint("public", __name__, template_folder="templates")

#

# move this to main models
class Service:
    def __init__(self, id, name, description, cost, imgurl):
        self.id = id
        self.name = name
        self.description = description
        self.cost = cost
        self.imgurl = imgurl
        self.urlstring = f"/services/{id}"
        

SERVICES = [
    Service(1, "Quick Clean", "A cheaper and faster residential cleaning.", 150.00, '/img/deepclean.png'),
    Service(2, "Move-Out Clean", "Moving out deep clean.", 300.00, '/img/moveout.png'),
    Service(3,  "Monthly Deep Clean", "A monthly deep clean to help you get caught up.", 300.00, '/img/monthlydeep.png'),
]


def get_service(id):
    for s in SERVICES:
        if s.id == id:
            return s
    # probably should return err if not


@public.route("/create-service-category", methods=["GET", "POST"])
def create_service_category():
    if request.method == "POST":
        print("POST request")
        cat = ServiceCategory(
            category=request.form["service-category"],
            description=request.form["description"],
        )
        print(cat)
        db.session.add(cat)
        db.session.commit()
        print("Made it to the commit without an error")
        return redirect(url_for('public.services'))
    if request.method == "GET":
        print("GET request")
        return redirect(url_for('public.index'))
    return redirect(url_for('public.about'))
    

@public.route("/")
def index():
    cats = db.session.execute(db.select(ServiceCategory).order_by(ServiceCategory.category)).scalars()
    context = {"title": "Greenleaf Cleaning", "company_name": "Greenleaf Cleaning LLC", "serv-cats": cats,}
    return render_template("index.html", context=context)


@public.route("/about")
def about(): 
    context = {"title": "Greenleaf Cleaning", "company_name": "Greenleaf Cleaning LLC",}
    return render_template("about.html", context=context)

@public.route("/services")
def services():

    context = {
        "title": "Greenleaf Cleaning",
        "services": SERVICES,
        "company_name": "Greenleaf Cleaning",
    }
    return render_template("services.html", context=context)


@public.route("/services/<int:service_id>")
def service(service_id):

    service = get_service(service_id)
    
    context = {
        "title": "Greenleaf Cleaning",
        "service": service,
    }
    return render_template("service.html", context=context)


@public.route("/book-job", methods=["POST"])
def book_cleaning():

    print("Booked job.")
    return redirect(url_for("public.index"))


@public.route("/group-cleaning")
def group_cleaning():
    context = {"title": "Greenleaf Cleaning", "company_name": "Greenleaf Cleaning LLC",}
    return render_template("groups.html", context=context)

