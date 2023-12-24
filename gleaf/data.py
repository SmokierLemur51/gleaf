import sqlite3 
import click

from flask import current_app, g

# https://flask.palletsprojects.com/en/3.0.x/tutorial/database/

# This came from the Flask documentation, trying to steer clear of initiatializing the
# db with SQLAlchemy becuase it gets a little messy when you do.



def get_db():
    """
    Opens a new database connection if there is none yet for the
    current application context.
    """
    if 'db' not in g:
        g.db = sqlite3.connect(
            current_app.config['DATABASE'],
            detect_types=sqlite3.PARSE_DECLTYPES
        )
        g.db.row_factory = sqlite3.Row
    return g.db

def close_db(e=None):
    """If this request connected to the database, close the
    connection.
    """
    db = g.pop('db', None)
    if db is not None:
        db.close()


def init_db():
    db = get_db()

    with current_app.open_resource('schema.sql') as f:
        db.executescript(f.read().decode('utf8'))


@click.command('init-db')
def init_db_command():
    """Clear the existing data and create new tables."""
    init_db()
    click.echo('Initialized the database.')


def init_app(app):
    app.teardown_appcontext(close_db)
    app.cli.add_command(init_db_command)


def create_table(db_file, table_name, columns):
    """
        Paramters:
        - db_file: The database file to connect to.
        - table_name: The name of the table to be created.
        - columns:  List of column definitions, where each column definition is a tuple
                    containing (column_name, data_type)
        Example Usage:
        - create_table("countyfair.db", "prize_hog", [("id", "INTEGER PRIMARY KEY AUTOINCREMENT"), ("name", "TEXT"), ("age", "INTEGER")])
    """
    connection = sqlite3.connect(db_file)
    cursor = connection.cursor()
    createTableSQL = f"CREATE TABLE IF NOT EXISTS {table_name} ("
    for column in columns:
        column_name, data_type = column
        createTableSQL += f"{column_name} {data_type}, "
    createTableSQL = createTableSQL.rstrip(", ")
    createTableSQL += ");"
    try:
        cursor.execute(createTableSQL)
        print(f"\t*\tTable {table_name} created successfully.")
        connection.commit()
    except sqlite3.Error as e:
        print(f"Error creating table {table_name}: {e}")
    finally:
        connection.close()
