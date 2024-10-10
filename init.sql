CREATE TABLE Employees
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	ID INTEGER PRIMARY KEY,
	Name TEXT,
	Born TIMESTAMP
);

CREATE TABLE Users
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	Login TEXT PRIMARY KEY
);

CREATE TABLE Subscriptions
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	UserLogin TEXT,
	EmployeeID INTEGER,
	FOREIGN KEY (UserLogin) REFERENCES Users (login),
	FOREIGN KEY (EmployeeID) REFERENCES Employees(ID)
);
