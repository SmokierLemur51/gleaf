from gleaf import create_app, db

app = create_app()

with app.app_context():
    db.create_all()

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=81, debug=True)




# const (
# 	PORT = ":5000"

# // host     = "localhost"
# // port     = 5432
# // user     = "postgres"
# // password = "1lP(=F=<HHwD]v"
# // dbname   = "gleaftesting"
# )