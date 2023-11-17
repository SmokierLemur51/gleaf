from flask import Blueprint, render_template, flash, url_for, redirect, request
from gleaf.portal import forms

from gleaf.portal.forms import CreateServiceCategoryForm
from ..extensions import db
from ..models.main import ServiceCategory

portal = Blueprint("portal", __name__, template_folder="templates")

@portal.route("/test", methods=["POST"])
def test():
    # test forms
    if request.method == "POST":

        cat = ServiceCategory(category=request.form.get('category'), description=request.form.get('description'))
        db.session.add(cat)
        db.session.commit()
        return redirect(url_for('portal.index'))
    return redirect(url_for('portal.index'))


@portal.route("/", methods=["GET", "POST"])
def index():
    form = CreateServiceCategoryForm()
    context = {
        "title": "Greenleaf Cleaning",
        "company_name": "Greenleaf Cleaning LLC",
    }
    if form.validate_on_submit():
        print(f"Category: {form.category.data}\tDescription: {form.description.data}")
        cat = ServiceCategory(category=form.category.data,
                              description=form.description.data)
        db.session.add(cat)
        db.session.commit()
        flash("Your service category has been created.", "success")
        return redirect(url_for("portal.index"))
    cats = ServiceCategory.query.order_by(ServiceCategory.category)
    return render_template("index.html",
                           form=form,
                           context=context,
                           legend="New Service Category", cats=cats)


@portal.route("/schedule")
def schedule():
    return render_template("schedule.html")


@portal.route("/finances")
def finances():
    return render_template("finances.html")


@portal.route("/service-management")
def service_management():

    return render_template("service-management.html")


@portal.route("/reports")
def reports():
    return render_template("reports.html")
